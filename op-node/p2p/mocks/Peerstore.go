// Code generated by mockery v2.22.1. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	peer "github.com/libp2p/go-libp2p/core/peer"

	store "github.com/ethereum-optimism/optimism/op-node/p2p/store"
)

// Peerstore is an autogenerated mock type for the Peerstore type
type Peerstore struct {
	mock.Mock
}

// PeerInfo provides a mock function with given fields: _a0
func (_m *Peerstore) PeerInfo(_a0 peer.ID) peer.AddrInfo {
	ret := _m.Called(_a0)

	var r0 peer.AddrInfo
	if rf, ok := ret.Get(0).(func(peer.ID) peer.AddrInfo); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(peer.AddrInfo)
	}

	return r0
}

// Peers provides a mock function with given fields:
func (_m *Peerstore) Peers() peer.IDSlice {
	ret := _m.Called()

	var r0 peer.IDSlice
	if rf, ok := ret.Get(0).(func() peer.IDSlice); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(peer.IDSlice)
		}
	}

	return r0
}

// SetScore provides a mock function with given fields: _a0, _a1, _a2
func (_m *Peerstore) SetScore(_a0 peer.ID, _a1 store.ScoreType, _a2 float64) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(peer.ID, store.ScoreType, float64) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewPeerstore interface {
	mock.TestingT
	Cleanup(func())
}

// NewPeerstore creates a new instance of Peerstore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPeerstore(t mockConstructorTestingTNewPeerstore) *Peerstore {
	mock := &Peerstore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
