package rpc

import (
	"context"
	"fmt"

	oplog "github.com/ethereum-optimism/optimism/op-service/log"
	"github.com/ethereum/go-ethereum/log"
)

type RpcMetrics interface {
	// RecordRPCServerRequest returns a function that records the duration of serving the given RPC method
	RecordRPCServerRequest(method string) func()
}

type CommonAdminAPI struct {
	M   RpcMetrics
	log log.Logger
}

func NewCommonAdminAPI(m RpcMetrics, log log.Logger) *CommonAdminAPI {
	return &CommonAdminAPI{
		M:   m,
		log: log,
	}
}

func (n *CommonAdminAPI) SetLogLevel(ctx context.Context, lvlStr string) error {
	recordDur := n.M.RecordRPCServerRequest("admin_setLogLevel")
	defer recordDur()

	h := n.log.GetHandler()

	lvl, err := log.LvlFromString(lvlStr)
	if err != nil {
		return err
	}

	// We set the log level, and do not wrap the handler with an additional filter handler,
	// as the underlying handler would otherwise also still filter with the previous log level.
	lvlSetter, ok := h.(oplog.LvlSetter)
	if !ok {
		return fmt.Errorf("log handler type %T cannot change log level", h)
	}
	lvlSetter.SetLogLevel(lvl)
	return nil
}
