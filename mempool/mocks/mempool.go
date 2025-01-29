// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	abcicli "github.com/cometbft/cometbft/abci/client"
	mempool "github.com/cometbft/cometbft/mempool"

	mock "github.com/stretchr/testify/mock"

	p2p "github.com/cometbft/cometbft/p2p"

	types "github.com/cometbft/cometbft/types"

	v1 "github.com/cometbft/cometbft/api/cometbft/abci/v1"
)

// Mempool is an autogenerated mock type for the Mempool type
type Mempool struct {
	mock.Mock
}

// CheckTx provides a mock function with given fields: tx, sender
func (_m *Mempool) CheckTx(tx types.Tx, sender p2p.ID) (*abcicli.ReqRes, error) {
	ret := _m.Called(tx, sender)

	if len(ret) == 0 {
		panic("no return value specified for CheckTx")
	}

	var r0 *abcicli.ReqRes
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Tx, p2p.ID) (*abcicli.ReqRes, error)); ok {
		return rf(tx, sender)
	}
	if rf, ok := ret.Get(0).(func(types.Tx, p2p.ID) *abcicli.ReqRes); ok {
		r0 = rf(tx, sender)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*abcicli.ReqRes)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Tx, p2p.ID) error); ok {
		r1 = rf(tx, sender)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EnableTxsAvailable provides a mock function with no fields
func (_m *Mempool) EnableTxsAvailable() {
	_m.Called()
}

// Flush provides a mock function with no fields
func (_m *Mempool) Flush() {
	_m.Called()
}

// FlushAppConn provides a mock function with no fields
func (_m *Mempool) FlushAppConn() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for FlushAppConn")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetTxByKey provides a mock function with given fields: key
func (_m *Mempool) GetTxByKey(key types.TxKey) (types.Tx, bool) {
	ret := _m.Called(key)

	if len(ret) == 0 {
		panic("no return value specified for GetTxByKey")
	}

	var r0 types.Tx
	var r1 bool
	if rf, ok := ret.Get(0).(func(types.TxKey) (types.Tx, bool)); ok {
		return rf(key)
	}
	if rf, ok := ret.Get(0).(func(types.TxKey) types.Tx); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.Tx)
		}
	}

	if rf, ok := ret.Get(1).(func(types.TxKey) bool); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// Lock provides a mock function with no fields
func (_m *Mempool) Lock() {
	_m.Called()
}

// PreUpdate provides a mock function with no fields
func (_m *Mempool) PreUpdate() {
	_m.Called()
}

// ReapMaxBytesMaxGas provides a mock function with given fields: maxBytes, maxGas
func (_m *Mempool) ReapMaxBytesMaxGas(maxBytes int64, maxGas int64) types.Txs {
	ret := _m.Called(maxBytes, maxGas)

	if len(ret) == 0 {
		panic("no return value specified for ReapMaxBytesMaxGas")
	}

	var r0 types.Txs
	if rf, ok := ret.Get(0).(func(int64, int64) types.Txs); ok {
		r0 = rf(maxBytes, maxGas)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.Txs)
		}
	}

	return r0
}

// ReapMaxTxs provides a mock function with given fields: max
func (_m *Mempool) ReapMaxTxs(max int) types.Txs {
	ret := _m.Called(max)

	if len(ret) == 0 {
		panic("no return value specified for ReapMaxTxs")
	}

	var r0 types.Txs
	if rf, ok := ret.Get(0).(func(int) types.Txs); ok {
		r0 = rf(max)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.Txs)
		}
	}

	return r0
}

// RemoveTxByKey provides a mock function with given fields: txKey
func (_m *Mempool) RemoveTxByKey(txKey types.TxKey) error {
	ret := _m.Called(txKey)

	if len(ret) == 0 {
		panic("no return value specified for RemoveTxByKey")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(types.TxKey) error); ok {
		r0 = rf(txKey)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Size provides a mock function with no fields
func (_m *Mempool) Size() int {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Size")
	}

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// SizeBytes provides a mock function with no fields
func (_m *Mempool) SizeBytes() int64 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for SizeBytes")
	}

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// TxsAvailable provides a mock function with no fields
func (_m *Mempool) TxsAvailable() <-chan struct{} {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for TxsAvailable")
	}

	var r0 <-chan struct{}
	if rf, ok := ret.Get(0).(func() <-chan struct{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan struct{})
		}
	}

	return r0
}

// Unlock provides a mock function with no fields
func (_m *Mempool) Unlock() {
	_m.Called()
}

// Update provides a mock function with given fields: blockHeight, blockTxs, deliverTxResponses, newPreFn, newPostFn
func (_m *Mempool) Update(blockHeight int64, blockTxs types.Txs, deliverTxResponses []*v1.ExecTxResult, newPreFn mempool.PreCheckFunc, newPostFn mempool.PostCheckFunc) error {
	ret := _m.Called(blockHeight, blockTxs, deliverTxResponses, newPreFn, newPostFn)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int64, types.Txs, []*v1.ExecTxResult, mempool.PreCheckFunc, mempool.PostCheckFunc) error); ok {
		r0 = rf(blockHeight, blockTxs, deliverTxResponses, newPreFn, newPostFn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WasRecentlyEvicted provides a mock function with given fields: key
func (_m *Mempool) WasRecentlyEvicted(key types.TxKey) bool {
	ret := _m.Called(key)

	if len(ret) == 0 {
		panic("no return value specified for WasRecentlyEvicted")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(types.TxKey) bool); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// NewMempool creates a new instance of Mempool. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMempool(t interface {
	mock.TestingT
	Cleanup(func())
}) *Mempool {
	mock := &Mempool{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
