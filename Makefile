COMPOSEFLAGS=-d
ITESTS_L2_HOST=http://localhost:9545
BEDROCK_TAGS_REMOTE?=origin

build: build-go build-ts
.PHONY: build

build-go: submodules op-node op-proposer op-batcher
.PHONY: build-go

build-ts: submodules
	if [ -n "$$NVM_DIR" ]; then \
		. $$NVM_DIR/nvm.sh && nvm use; \
	fi
	yarn install
	yarn build
.PHONY: build-ts

submodules:
	# CI will checkout submodules on its own (and fails on these commands)
	if [ -z "$$GITHUB_ENV" ]; then \
		git submodule init; \
		git submodule update; \
	fi
.PHONY: submodules

op-bindings:
	make -C ./op-bindings
.PHONY: op-bindings

op-node:
	make -C ./op-node op-node
.PHONY: op-node

op-batcher:
	make -C ./op-batcher op-batcher
.PHONY: op-batcher

op-proposer:
	make -C ./op-proposer op-proposer
.PHONY: op-proposer

mod-tidy:
	# Below GOPRIVATE line allows mod-tidy to be run immediately after
	# releasing new versions. This bypasses the Go modules proxy, which
	# can take a while to index new versions.
	# 
	# See https://proxy.golang.org/ for more info.
	export GOPRIVATE="github.com/ethereum-optimism" && \
	cd ./op-service && go mod tidy && cd .. && \
	cd ./op-node && go mod tidy && cd .. && \
	cd ./op-proposer && go mod tidy && cd ..  && \
	cd ./op-batcher && go mod tidy && cd ..  && \
	cd ./op-bindings && go mod tidy && cd ..  && \
	cd ./op-e2e && go mod tidy && cd ..
.PHONY: mod-tidy

clean:
	rm -rf ./bin
.PHONY: clean

nuke: clean devnet-clean
	git clean -Xdf
.PHONY: nuke

devnet-up:
	@bash ./ops-bedrock/devnet-up.sh
.PHONY: devnet-up

devnet-down:
	@(cd ./ops-bedrock && GENESIS_TIMESTAMP=$(shell date +%s) docker-compose stop)
.PHONY: devnet-down

devnet-clean:
	rm -rf ./packages/contracts-bedrock/deployments/devnetL1
	rm -rf ./.devnet
	cd ./ops-bedrock && docker-compose down
	docker image ls 'ops-bedrock*' --format='{{.Repository}}' | xargs -r docker rmi
	docker volume ls --filter name=ops-bedrock --format='{{.Name}}' | xargs -r docker volume rm
.PHONY: devnet-clean

devnet-logs:
	@(cd ./ops-bedrock && docker-compose logs -f)
	.PHONY: devnet-logs

test-unit:
	make -C ./op-node test
	make -C ./op-proposer test
	make -C ./op-batcher test
	make -C ./op-e2e test
	yarn test
.PHONY: test-unit

test-integration:
	bash ./ops-bedrock/test-integration.sh \
		./packages/contracts-bedrock/deployments/devnetL1
.PHONY: test-integration

# Remove the baseline-commit to generate a base reading & show all issues
semgrep:
	$(eval DEV_REF := $(shell git rev-parse develop))
	SEMGREP_REPO_NAME=ethereum-optimism/optimism semgrep ci --baseline-commit=$(DEV_REF)
.PHONY: semgrep

clean-node-modules:
	rm -rf node_modules
	rm -rf packages/**/node_modules


tag-bedrock-go-modules:
	@if [ -z "$(VERSION)" ]; then \
		echo "You must specify a version."; \
		exit 0; \
	fi

	@FIRST_CHAR=$(shell printf '%s' "$(VERSION)" | cut -c1); \
	if [ "$$FIRST_CHAR" != "v" ]; then \
		echo "Tag must start with v."; \
		exit 0; \
	fi

	git tag "op-proposer/$(VERSION)"
	git tag "op-node/$(VERSION)"
	git tag "op-e2e/$(VERSION)"
	git tag "op-bindings/$(VERSION)"
	git tag "op-batcher/$(VERSION)"
	git tag "op-service/$(VERSION)"
	git push $(BEDROCK_TAGS_REMOTE) "op-proposer/$(VERSION)"
	git push $(BEDROCK_TAGS_REMOTE) "op-node/$(VERSION)"
	git push $(BEDROCK_TAGS_REMOTE) "op-e2e/$(VERSION)"
	git push $(BEDROCK_TAGS_REMOTE) "op-bindings/$(VERSION)"
	git push $(BEDROCK_TAGS_REMOTE) "op-batcher/$(VERSION)"
	git push $(BEDROCK_TAGS_REMOTE) "op-service/$(VERSION)"

