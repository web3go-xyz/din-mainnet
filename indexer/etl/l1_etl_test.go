package etl

import (
	"math/big"

	"github.com/ethereum-optimism/optimism/op-service/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/ethereum-optimism/optimism/indexer/config"
	"github.com/ethereum-optimism/optimism/indexer/database"
	"github.com/ethereum-optimism/optimism/indexer/node"

	"testing"
)

// test suite
type l1EtlTs struct {
	db        *database.MockDB
	client    *node.MockEthClient
	start     *big.Int
	contracts config.L1Contracts
}

func Test_L1ETL_Construction(t *testing.T) {
	var tests = []struct {
		name         string
		construction func() *l1EtlTs
		assertion    func(*L1ETL, error)
	}{
		{
			name: "Start from L1 config height",
			construction: func() *l1EtlTs {
				client := new(node.MockEthClient)
				db := database.NewMockDB()

				testStart := big.NewInt(100)

				db.MockBlocks.On("L1LatestBlockHeader").Return(nil, nil)

				client.On("BlockHeaderByNumber", mock.MatchedBy(
					node.BigIntMatcher(100))).Return(
					&types.Header{
						ParentHash: common.HexToHash("0x69"),
					}, nil)

				client.On("GethEthClient").Return(nil)

				return &l1EtlTs{
					db:        db,
					client:    client,
					start:     testStart,
					contracts: config.L1Contracts{},
				}
			},
			assertion: func(etl *L1ETL, err error) {
				require.NoError(t, err)
				require.Equal(t, etl.headerTraversal.LastHeader().ParentHash, common.HexToHash("0x69"))
			},
		},
		{
			name: "Start from recent height stored in DB",
			construction: func() *l1EtlTs {
				client := new(node.MockEthClient)
				db := database.NewMockDB()

				testStart := big.NewInt(100)

				db.MockBlocks.On("L1LatestBlockHeader").Return(
					&database.L1BlockHeader{
						BlockHeader: database.BlockHeader{
							RLPHeader: &database.RLPHeader{
								Number: big.NewInt(69),
							},
						}}, nil)

				client.On("GethEthClient").Return(nil)

				return &l1EtlTs{
					db:        db,
					client:    client,
					start:     testStart,
					contracts: config.L1Contracts{},
				}
			},
			assertion: func(etl *L1ETL, err error) {
				require.NoError(t, err)
				header := etl.headerTraversal.LastHeader()

				require.True(t, header.Number.Cmp(big.NewInt(69)) == 0)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts := test.construction()

			logger := log.NewLogger(log.DefaultCLIConfig())

			etl, err := NewL1ETL(logger, ts.db.DB, ts.client, ts.start, ts.contracts)
			test.assertion(etl, err)
		})
	}
}
