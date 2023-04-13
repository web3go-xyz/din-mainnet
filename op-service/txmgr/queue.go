package txmgr

import (
	"context"
	"math"
	"sync"
	"sync/atomic"

	"github.com/ethereum/go-ethereum/core/types"
	"golang.org/x/sync/semaphore"
)

type TxReceipt[T any] struct {
	// ID can be used to identify unique tx receipts within the recept channel
	ID T
	// Receipt result from the transaction send
	Receipt *types.Receipt
	// Err contains any error that occurred during the tx send
	Err error
}

// TxFactory should return the next transaction to send (and associated identifier).
// If no transaction is available, an error should be returned (such as io.EOF).
type TxFactory[T any] func(ctx context.Context) (TxCandidate, T, error)

type Queue[T any] struct {
	txMgr          TxManager
	pendingChanged func(uint64)
	pending        atomic.Uint64
	semaphore      *semaphore.Weighted
	wg             sync.WaitGroup
}

// NewQueue creates a new transaction sending Queue, with the following parameters:
//   - maxPending: max number of pending txs at once (0 == no limit)
//   - pendingChanged: called whenever a tx send starts or finishes. The
//     number of currently pending txs is passed as a parameter.
func NewQueue[T any](txMgr TxManager, maxPending uint64, pendingChanged func(uint64)) *Queue[T] {
	if maxPending > math.MaxInt64 {
		// ensure we don't overflow as semaphore only accepts int64; in reality this will never be an issue
		maxPending = math.MaxInt64
	}
	var s *semaphore.Weighted
	if maxPending > 0 {
		// only create a semaphore for limited-size queues
		s = semaphore.NewWeighted(int64(maxPending))
	}
	return &Queue[T]{
		txMgr:          txMgr,
		pendingChanged: pendingChanged,
		semaphore:      s,
	}
}

// Wait waits for all pending txs to complete (or fail).
func (q *Queue[T]) Wait() {
	q.wg.Wait()
}

// Send will wait until the number of pending txs is below the max pending,
// and then send the next tx. The TxFactory should return an error if the
// next tx does not exist, which will be returned from this method.
//
// The actual tx sending is non-blocking, with the receipt returned on the
// provided receipt channel.
func (q *Queue[T]) Send(ctx context.Context, factory TxFactory[T], receiptCh chan TxReceipt[T]) error {
	if q.semaphore != nil {
		if err := q.semaphore.Acquire(ctx, 1); err != nil {
			return err
		}
	}
	return q.trySend(ctx, factory, receiptCh)
}

// TrySend sends the next tx, but only if the number of pending txs is below the
// max pending, otherwise the TxFactory is not called (and no error is returned).
// The TxFactory should return an error if the next tx does not exist, which is
// returned from this method.
//
// Returns false if there is no room in the queue to send. Otherwise, the
// transaction is queued and this method returns true.
//
// The actual tx sending is non-blocking, with the receipt returned on the
// provided receipt channel.
func (q *Queue[T]) TrySend(ctx context.Context, factory TxFactory[T], receiptCh chan TxReceipt[T]) (bool, error) {
	if q.semaphore != nil && !q.semaphore.TryAcquire(1) {
		return false, nil
	}
	err := q.trySend(ctx, factory, receiptCh)
	return err == nil, err
}

func (q *Queue[T]) trySend(ctx context.Context, factory TxFactory[T], receiptCh chan TxReceipt[T]) error {
	candidate, id, err := factory(ctx)
	release := func() {
		if q.semaphore != nil {
			q.semaphore.Release(1)
		}
	}
	if err != nil {
		release()
		return err
	}

	q.pendingChanged(q.pending.Add(1))
	q.wg.Add(1)
	go func() {
		defer func() {
			release()
			q.pendingChanged(q.pending.Add(^uint64(0))) // -1
			q.wg.Done()
		}()
		receipt, err := q.txMgr.Send(ctx, candidate)
		receiptCh <- TxReceipt[T]{
			ID:      id,
			Receipt: receipt,
			Err:     err,
		}
	}()
	return nil
}
