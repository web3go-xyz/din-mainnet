package p2p_test

import (
	"testing"

	p2p "github.com/ethereum-optimism/optimism/op-node/p2p"
	p2pMocks "github.com/ethereum-optimism/optimism/op-node/p2p/mocks"
	"github.com/ethereum-optimism/optimism/op-node/testlog"
	log "github.com/ethereum/go-ethereum/log"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	peer "github.com/libp2p/go-libp2p/core/peer"
	suite "github.com/stretchr/testify/suite"
)

// PeerScorerTestSuite tests peer parameterization.
type PeerScorerTestSuite struct {
	suite.Suite

	// mockConnGater *p2pMocks.ConnectionGater
	mockGater      *p2pMocks.PeerGater
	mockStore      *p2pMocks.Peerstore
	mockMetricer   *p2pMocks.GossipMetricer
	mockBandScorer *p2pMocks.BandScorer
	logger         log.Logger
}

// SetupTest sets up the test suite.
func (testSuite *PeerScorerTestSuite) SetupTest() {
	testSuite.mockGater = &p2pMocks.PeerGater{}
	testSuite.mockStore = &p2pMocks.Peerstore{}
	testSuite.mockMetricer = &p2pMocks.GossipMetricer{}
	testSuite.mockBandScorer = &p2pMocks.BandScorer{}
	testSuite.logger = testlog.Logger(testSuite.T(), log.LvlError)
}

// TestPeerScorer runs the PeerScorerTestSuite.
func TestPeerScorer(t *testing.T) {
	suite.Run(t, new(PeerScorerTestSuite))
}

// TestPeerScorerOnConnect ensures we can call the OnConnect method on the peer scorer.
func (testSuite *PeerScorerTestSuite) TestPeerScorerOnConnect() {
	scorer := p2p.NewScorer(
		testSuite.mockGater,
		testSuite.mockStore,
		testSuite.mockMetricer,
		testSuite.mockBandScorer,
		testSuite.logger,
	)
	scorer.OnConnect()
}

// TestPeerScorerOnDisconnect ensures we can call the OnDisconnect method on the peer scorer.
func (testSuite *PeerScorerTestSuite) TestPeerScorerOnDisconnect() {
	scorer := p2p.NewScorer(
		testSuite.mockGater,
		testSuite.mockStore,
		testSuite.mockMetricer,
		testSuite.mockBandScorer,
		testSuite.logger,
	)
	scorer.OnDisconnect()
}

// TestSnapshotHook tests running the snapshot hook on the peer scorer.
func (testSuite *PeerScorerTestSuite) TestSnapshotHook() {
	scorer := p2p.NewScorer(
		testSuite.mockGater,
		testSuite.mockStore,
		testSuite.mockMetricer,
		testSuite.mockBandScorer,
		testSuite.logger,
	)
	inspectFn := scorer.SnapshotHook()

	// Mock the band scorer calls
	testSuite.mockBandScorer.On("Reset").Return(nil)
	testSuite.mockBandScorer.On("Bucket", float64(-100)).Return("graylist")

	// Mock the peer gater call
	testSuite.mockGater.On("Update", peer.ID("peer1"), float64(-100)).Return(nil)

	// The metricer should then be called with the peer score band map
	testSuite.mockMetricer.On("SetPeerScores", map[string]float64{
		"graylist": 1,
	}).Return(nil)

	// Apply the snapshot
	snapshotMap := map[peer.ID]*pubsub.PeerScoreSnapshot{
		peer.ID("peer1"): {
			Score: -100,
		},
	}
	inspectFn(snapshotMap)
}

// TestSnapshotHookBlockPeer tests running the snapshot hook on the peer scorer with a peer score below the threshold.
// This implies that the peer should be blocked.
func (testSuite *PeerScorerTestSuite) TestSnapshotHookBlockPeer() {
	scorer := p2p.NewScorer(
		testSuite.mockGater,
		testSuite.mockStore,
		testSuite.mockMetricer,
		testSuite.mockBandScorer,
		testSuite.logger,
	)
	inspectFn := scorer.SnapshotHook()

	// Mock the band scorer calls
	testSuite.mockBandScorer.On("Reset").Return(nil)
	testSuite.mockBandScorer.On("Bucket", float64(-101)).Return("graylist")

	// Mock the peer gater call
	testSuite.mockGater.On("Update", peer.ID("peer1"), float64(-101)).Return(nil)

	// The metricer should then be called with the peer score band map
	testSuite.mockMetricer.On("SetPeerScores", map[string]float64{
		"graylist": 1,
	}).Return(nil)

	// Apply the snapshot
	snapshotMap := map[peer.ID]*pubsub.PeerScoreSnapshot{
		peer.ID("peer1"): {
			Score: -101,
		},
	}
	inspectFn(snapshotMap)
}
