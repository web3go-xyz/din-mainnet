#!/usr/bin/env bash

# Usage: check-changed.sh <diff-pattern>.
#
# This script compares the files changed in the <diff-pattern> to the git diff,
# and writes TRUE or FALSE to stdout if the diff matches/does not match. It is
# used by CircleCI jobs to determine if they need to run.

set -e

echoerr() { echo "$@" 1>&2; }

# Check if this is a CircleCI PR.
if [[ -n $CIRCLE_PULL_REQUEST ]]; then
	PACKAGE=$1
	# Craft the URL to the GitHub API. The access token is optional for the monorepo since it's an open-source repo.
	GITHUB_API_URL=$(echo "https://api.github.com/repos/${CIRCLE_PULL_REQUEST}?access_token=$GITHUB_ACCESS_TOKEN" | sed "s/\/pull\//\/pulls\//")
	# Grab the PR's base ref using the GitHub API.
	REF=$(curl -s "$GITHUB_API_URL" | jq -r ".base.ref")

	echoerr "Base Ref:     $REF"
	echoerr "Base Ref SHA: $(git show-branch --sha1-name "$REF")"
 	echoerr "Curr Ref:     $(git rev-parse --short HEAD)"

 	# Compare HEAD to the PR's base ref, stripping out the change percentages that come with git diff --dirstat.
 	# Pass in the diff pattern to grep, and echo TRUE if there's a match. False otherwise.
 	(git diff --dirstat=files,0 "$REF...HEAD" | sed 's/^[ 0-9.]\+% //g' | grep -q -E "$PACKAGE" && echo "TRUE") || echo "FALSE"
else
	# Non-PR builds always require a rebuild.
	echoerr "Not a PR build, requiring a total rebuild."
	echo "TRUE"
fi
