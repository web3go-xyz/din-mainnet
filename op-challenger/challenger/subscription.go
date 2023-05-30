package challenger

import (
	"context"
	"errors"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

var (
	ErrMissingClient = errors.New("missing client")
)

// SubscriptionId is a unique subscription ID.
type SubscriptionId uint64

// Increment returns the next subscription ID.
func (s *SubscriptionId) Increment() SubscriptionId {
	*s++
	return *s
}

// Subscription wraps an [ethereum.Subscription] to provide a restart.
type Subscription struct {
	// The subscription ID
	id SubscriptionId
	// The current subscription
	sub ethereum.Subscription
	// The query used to create the subscription
	query ethereum.FilterQuery
	// The log channel
	logs chan types.Log
	// The quit channel
	quit chan struct{}
	// Filter client used to open the log subscription
	client ethereum.LogFilterer
	// Logger
	log log.Logger
}

// NewSubscription creates a new subscription.
func NewSubscription(query ethereum.FilterQuery, client ethereum.LogFilterer, log log.Logger) *Subscription {
	return &Subscription{
		id:     SubscriptionId(0),
		sub:    nil,
		query:  query,
		logs:   make(chan types.Log),
		quit:   make(chan struct{}),
		client: client,
		log:    log,
	}
}

// ID returns the subscription ID.
func (s *Subscription) ID() SubscriptionId {
	return s.id
}

// Started returns true if the subscription has started.
func (s *Subscription) Started() bool {
	return s.sub != nil
}

// Subscribe constructs the subscription.
func (s *Subscription) Subscribe() error {
	s.log.Info("Subscribing to", "query", s.query.Topics, "id", s.id)
	if s.client == nil {
		s.log.Error("missing client")
		return ErrMissingClient
	}
	sub, err := s.client.SubscribeFilterLogs(context.Background(), s.query, s.logs)
	if err != nil {
		s.log.Error("failed to subscribe to logs", "err", err)
		return err
	}
	s.sub = sub
	return nil
}

// Quit closes the subscription.
func (s *Subscription) Quit() {
	s.log.Info("Quitting subscription", "id", s.id)
	s.sub.Unsubscribe()
	s.quit <- struct{}{}
	s.log.Info("Quit subscription", "id", s.id)
}
