package database

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/mock"
)

type MockBlocksView struct {
	mock.Mock
}

func (m *MockBlocksView) L1BlockHeader(common.Hash) (*L1BlockHeader, error) {
	args := m.Called()

	header, ok := args.Get(0).(*L1BlockHeader)
	if !ok {
		header = nil
	}
	return header, args.Error(1)
}

func (m *MockBlocksView) L1BlockHeaderWithFilter(BlockHeader) (*L1BlockHeader, error) {
	args := m.Called()
	return args.Get(0).(*L1BlockHeader), args.Error(1)
}

func (m *MockBlocksView) L1LatestBlockHeader() (*L1BlockHeader, error) {
	args := m.Called()

	header, ok := args.Get(0).(*L1BlockHeader)
	if !ok {
		header = nil
	}

	return header, args.Error(1)
}

func (m *MockBlocksView) L2BlockHeader(common.Hash) (*L2BlockHeader, error) {
	args := m.Called()
	return args.Get(0).(*L2BlockHeader), args.Error(1)
}

func (m *MockBlocksView) L2BlockHeaderWithFilter(BlockHeader) (*L2BlockHeader, error) {
	args := m.Called()
	return args.Get(0).(*L2BlockHeader), args.Error(1)
}

func (m *MockBlocksView) L2LatestBlockHeader() (*L2BlockHeader, error) {
	args := m.Called()
	return args.Get(0).(*L2BlockHeader), args.Error(1)
}

func (m *MockBlocksView) LatestCheckpointedOutput() (*OutputProposal, error) {
	args := m.Called()
	return args.Get(0).(*OutputProposal), args.Error(1)
}

func (m *MockBlocksView) OutputProposal(index *big.Int) (*OutputProposal, error) {
	args := m.Called()
	return args.Get(0).(*OutputProposal), args.Error(1)
}

func (m *MockBlocksView) LatestEpoch() (*Epoch, error) {
	args := m.Called()
	return args.Get(0).(*Epoch), args.Error(1)
}

var _ BlocksDB = (*MockBlocksDB)(nil)

type MockBlocksDB struct {
	MockBlocksView
}

func (m *MockBlocksDB) StoreL1BlockHeaders(headers []L1BlockHeader) error {
	args := m.Called(headers)
	return args.Error(1)
}

func (m *MockBlocksDB) StoreL2BlockHeaders(headers []L2BlockHeader) error {
	args := m.Called(headers)
	return args.Error(1)
}

func (m *MockBlocksDB) StoreLegacyStateBatches(headers []LegacyStateBatch) error {
	args := m.Called(headers)
	return args.Error(1)
}
func (m *MockBlocksDB) StoreOutputProposals(headers []OutputProposal) error {
	args := m.Called(headers)
	return args.Error(1)
}

type MockDB struct {
	MockBlockView *MockBlocksView
	MockBlocks    *MockBlocksDB
	DB            *DB
}

func NewMockDB() *MockDB {
	// This is currently just mocking the BlocksDB interface
	// but can be expanded to mock other inner DB interfaces
	// as well
	mockBlocks := new(MockBlocksDB)
	db := &DB{Blocks: mockBlocks}

	return &MockDB{MockBlocks: mockBlocks, DB: db,
		MockBlockView: new(MockBlocksView)}
}
