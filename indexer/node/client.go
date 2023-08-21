package node

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

const (
	// defaultDialTimeout is default duration the processor will wait on
	// startup to make a connection to the backend
	defaultDialTimeout = 5 * time.Second

	// defaultRequestTimeout is the default duration the processor will
	// wait for a request to be fulfilled
	defaultRequestTimeout = 10 * time.Second
)

type EthClient interface {
	FinalizedBlockHeight() (*big.Int, error)

	BlockHeadersByRange(*big.Int, *big.Int) ([]types.Header, error)
	BlockHeaderByHash(common.Hash) (*types.Header, error)

	StorageHash(common.Address, *big.Int) (common.Hash, error)

	GethRpcClient() *rpc.Client
	GethEthClient() *ethclient.Client
}

type client struct {
	rpcClient *rpc.Client
}

func DialEthClient(rpcUrl string) (EthClient, error) {
	ctxwt, cancel := context.WithTimeout(context.Background(), defaultDialTimeout)
	defer cancel()

	rpcClient, err := rpc.DialContext(ctxwt, rpcUrl)
	if err != nil {
		return nil, err
	}

	client := &client{rpcClient: rpcClient}
	return client, nil
}

func NewEthClient(rpcClient *rpc.Client) EthClient {
	return &client{rpcClient}
}

func (c *client) GethRpcClient() *rpc.Client {
	return c.rpcClient
}

func (c *client) GethEthClient() *ethclient.Client {
	return ethclient.NewClient(c.GethRpcClient())
}

// FinalizedBlockHeight retrieves the latest block height in a finalized state
func (c *client) FinalizedBlockHeight() (*big.Int, error) {
	ctxwt, cancel := context.WithTimeout(context.Background(), defaultRequestTimeout)
	defer cancel()

	// Local devnet is having issues with the "finalized" block tag. Switch to "latest"
	// to iterate faster locally but this needs to be updated
	header := new(types.Header)
	err := c.rpcClient.CallContext(ctxwt, header, "eth_getBlockByNumber", "latest", false)
	if err != nil {
		return nil, err
	}

	return header.Number, nil
}

// BlockHeaderByHash retrieves the block header attributed to the supplied hash
func (c *client) BlockHeaderByHash(hash common.Hash) (*types.Header, error) {
	ctxwt, cancel := context.WithTimeout(context.Background(), defaultRequestTimeout)
	defer cancel()

	header, err := ethclient.NewClient(c.rpcClient).HeaderByHash(ctxwt, hash)
	if err != nil {
		return nil, err
	}

	// sanity check on the data returned
	if header.Hash() != hash {
		return nil, errors.New("header mismatch")
	}

	return header, nil
}

// BlockHeadersByRange will retrieve block headers within the specified range -- includsive. No restrictions
// are placed on the range such as blocks in the "latest", "safe" or "finalized" states. If the specified
// range is too large, `endHeight > latest`, the resulting list is truncated to the available headers
func (c *client) BlockHeadersByRange(startHeight, endHeight *big.Int) ([]types.Header, error) {
	count := new(big.Int).Sub(endHeight, startHeight).Uint64() + 1
	batchElems := make([]rpc.BatchElem, count)
	for i := uint64(0); i < count; i++ {
		height := new(big.Int).Add(startHeight, new(big.Int).SetUint64(i))
		batchElems[i] = rpc.BatchElem{
			Method: "eth_getBlockByNumber",
			Args:   []interface{}{toBlockNumArg(height), false},
			Result: new(types.Header),
			Error:  nil,
		}
	}

	ctxwt, cancel := context.WithTimeout(context.Background(), defaultRequestTimeout)
	defer cancel()
	err := c.rpcClient.BatchCallContext(ctxwt, batchElems)
	if err != nil {
		return nil, err
	}

	// Parse the headers.
	//  - Ensure integrity that they build on top of each other
	//  - Truncate out headers that do not exist (endHeight > "latest")
	size := 0
	headers := make([]types.Header, count)
	for i, batchElem := range batchElems {
		if batchElem.Error != nil {
			return nil, batchElem.Error
		} else if batchElem.Result == nil {
			break
		}

		header, ok := batchElem.Result.(*types.Header)
		if !ok {
			return nil, fmt.Errorf("unable to transform rpc response %v into types.Header", batchElem.Result)
		}
		if i > 0 && header.ParentHash != headers[i-1].Hash() {
			return nil, fmt.Errorf("queried header %s does not follow parent %s", header.Hash(), headers[i-1].Hash())
		}

		headers[i] = *header
		size = size + 1
	}

	headers = headers[:size]
	return headers, nil
}

// StorageHash returns the sha3 of the storage root for the specified account
func (c *client) StorageHash(address common.Address, blockNumber *big.Int) (common.Hash, error) {
	ctxwt, cancel := context.WithTimeout(context.Background(), defaultRequestTimeout)
	defer cancel()

	proof := struct{ StorageHash common.Hash }{}
	err := c.rpcClient.CallContext(ctxwt, &proof, "eth_getProof", address, nil, toBlockNumArg(blockNumber))
	if err != nil {
		return common.Hash{}, err
	}

	return proof.StorageHash, nil
}

func toBlockNumArg(number *big.Int) string {
	if number == nil {
		return "latest"
	} else if number.Sign() >= 0 {
		return hexutil.EncodeBig(number)
	}

	// It's negative.
	if number.IsInt64() {
		tag, _ := rpc.BlockNumber(number.Int64()).MarshalText()
		return string(tag)
	}

	// It's negative and large, which is invalid.
	return fmt.Sprintf("<invalid %d>", number)
}
