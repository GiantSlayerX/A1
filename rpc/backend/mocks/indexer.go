// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	abcitypes "github.com/cometbft/cometbft/abci/types"
	cometbfttypes "github.com/cometbft/cometbft/types"

	common "github.com/ethereum/go-ethereum/common"

	mock "github.com/stretchr/testify/mock"

	types "github.com/EscanBE/evermint/v12/types"
)

// EVMTxIndexer is an autogenerated mock type for the EVMTxIndexer type
type EVMTxIndexer struct {
	mock.Mock
}

// GetByBlockAndIndex provides a mock function with given fields: _a0, _a1
func (_m *EVMTxIndexer) GetByBlockAndIndex(_a0 int64, _a1 int32) (*types.TxResult, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetByBlockAndIndex")
	}

	var r0 *types.TxResult
	var r1 error
	if rf, ok := ret.Get(0).(func(int64, int32) (*types.TxResult, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(int64, int32) *types.TxResult); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.TxResult)
		}
	}

	if rf, ok := ret.Get(1).(func(int64, int32) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByTxHash provides a mock function with given fields: _a0
func (_m *EVMTxIndexer) GetByTxHash(_a0 common.Hash) (*types.TxResult, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for GetByTxHash")
	}

	var r0 *types.TxResult
	var r1 error
	if rf, ok := ret.Get(0).(func(common.Hash) (*types.TxResult, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(common.Hash) *types.TxResult); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.TxResult)
		}
	}

	if rf, ok := ret.Get(1).(func(common.Hash) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IndexBlock provides a mock function with given fields: _a0, _a1
func (_m *EVMTxIndexer) IndexBlock(_a0 *cometbfttypes.Block, _a1 []*abcitypes.ResponseDeliverTx) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for IndexBlock")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*cometbfttypes.Block, []*abcitypes.ResponseDeliverTx) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LastIndexedBlock provides a mock function with given fields:
func (_m *EVMTxIndexer) LastIndexedBlock() (int64, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for LastIndexedBlock")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func() (int64, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewEVMTxIndexer interface {
	mock.TestingT
	Cleanup(func())
}

// NewEVMTxIndexer creates a new instance of EVMTxIndexer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEVMTxIndexer(t mockConstructorTestingTNewEVMTxIndexer) *EVMTxIndexer {
	mock := &EVMTxIndexer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
