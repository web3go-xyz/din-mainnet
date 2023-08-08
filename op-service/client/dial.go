package client

import (
	"context"
	"fmt"
	"net"
	"net/url"
	"time"

	"github.com/ethereum-optimism/optimism/op-node/client"
	"github.com/ethereum-optimism/optimism/op-node/sources"
	"github.com/ethereum-optimism/optimism/op-service/backoff"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"
)

// DefaultDialTimeout is a default timeout for dialing a client.
const DefaultDialTimeout = 30 * time.Second
const defaultRetryCount = 30
const defaultRetryTime = 1 * time.Second

// DialEthClientWithTimeout attempts to dial the L1 provider using the provided
// URL. If the dial doesn't complete within defaultDialTimeout seconds, this
// method will return an error.
func DialEthClientWithTimeout(timeout time.Duration, log log.Logger, url string) (*ethclient.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	c, err := dialRPCClientWithBackoff(ctx, log, url)
	if err != nil {
		return nil, err
	}

	return ethclient.NewClient(c), nil
}

// DialRollupClientWithTimeout attempts to dial the RPC provider using the provided URL.
// If the dial doesn't complete within timeout seconds, this method will return an error.
func DialRollupClientWithTimeout(timeout time.Duration, log log.Logger, url string) (*sources.RollupClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	rpcCl, err := dialRPCClientWithBackoff(ctx, log, url)
	if err != nil {
		return nil, err
	}

	return sources.NewRollupClient(client.NewBaseRPCClient(rpcCl)), nil
}

// Dials a JSON-RPC endpoint repeatedly, with a backoff, until a client connection is established. Auth is optional.
func dialRPCClientWithBackoff(ctx context.Context, log log.Logger, addr string) (*rpc.Client, error) {
	bOff := backoff.Fixed(defaultRetryTime)
	var ret *rpc.Client
	err := backoff.DoCtx(ctx, defaultRetryCount, bOff, func() error {
		client, err := rpc.DialOptions(ctx, addr)
		if err != nil {
			return fmt.Errorf("failed to dial address (%s): %w", addr, err)
		}
		// log.Warn("failed to dial address, but may connect later", "addr", addr, "err", err)
		ret = client
		return nil
	})
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func IsURLAvailable(address string) bool {
	u, err := url.Parse(address)
	if err != nil {
		return false
	}
	conn, err := net.DialTimeout("tcp", u.Host, 1*time.Second)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}
