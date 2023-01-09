// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/blocknative/dreamboat/pkg/relay (interfaces: Datastore,State)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	structs "github.com/blocknative/dreamboat/pkg/structs"
	types "github.com/flashbots/go-boost-utils/types"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockDatastore is a mock of Datastore interface
type MockDatastore struct {
	ctrl     *gomock.Controller
	recorder *MockDatastoreMockRecorder
}

// MockDatastoreMockRecorder is the mock recorder for MockDatastore
type MockDatastoreMockRecorder struct {
	mock *MockDatastore
}

// NewMockDatastore creates a new mock instance
func NewMockDatastore(ctrl *gomock.Controller) *MockDatastore {
	mock := &MockDatastore{ctrl: ctrl}
	mock.recorder = &MockDatastoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDatastore) EXPECT() *MockDatastoreMockRecorder {
	return m.recorder
}

// CacheBlock mocks base method
func (m *MockDatastore) CacheBlock(arg0 context.Context, arg1 *structs.CompleteBlockstruct) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CacheBlock", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CacheBlock indicates an expected call of CacheBlock
func (mr *MockDatastoreMockRecorder) CacheBlock(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CacheBlock", reflect.TypeOf((*MockDatastore)(nil).CacheBlock), arg0, arg1)
}

// CheckSlotDelivered mocks base method
func (m *MockDatastore) CheckSlotDelivered(arg0 context.Context, arg1 uint64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckSlotDelivered", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckSlotDelivered indicates an expected call of CheckSlotDelivered
func (mr *MockDatastoreMockRecorder) CheckSlotDelivered(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckSlotDelivered", reflect.TypeOf((*MockDatastore)(nil).CheckSlotDelivered), arg0, arg1)
}

// GetDelivered mocks base method
func (m *MockDatastore) GetDelivered(arg0 context.Context, arg1 structs.PayloadQuery) (structs.BidTraceWithTimestamp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDelivered", arg0, arg1)
	ret0, _ := ret[0].(structs.BidTraceWithTimestamp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDelivered indicates an expected call of GetDelivered
func (mr *MockDatastoreMockRecorder) GetDelivered(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDelivered", reflect.TypeOf((*MockDatastore)(nil).GetDelivered), arg0, arg1)
}

// GetDeliveredBatch mocks base method
func (m *MockDatastore) GetDeliveredBatch(arg0 context.Context, arg1 []structs.PayloadQuery) ([]structs.BidTraceWithTimestamp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeliveredBatch", arg0, arg1)
	ret0, _ := ret[0].([]structs.BidTraceWithTimestamp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeliveredBatch indicates an expected call of GetDeliveredBatch
func (mr *MockDatastoreMockRecorder) GetDeliveredBatch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeliveredBatch", reflect.TypeOf((*MockDatastore)(nil).GetDeliveredBatch), arg0, arg1)
}

// GetHeadersByBlockHash mocks base method
func (m *MockDatastore) GetHeadersByBlockHash(arg0 context.Context, arg1 types.Hash) ([]structs.HeaderAndTrace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHeadersByBlockHash", arg0, arg1)
	ret0, _ := ret[0].([]structs.HeaderAndTrace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHeadersByBlockHash indicates an expected call of GetHeadersByBlockHash
func (mr *MockDatastoreMockRecorder) GetHeadersByBlockHash(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHeadersByBlockHash", reflect.TypeOf((*MockDatastore)(nil).GetHeadersByBlockHash), arg0, arg1)
}

// GetHeadersByBlockNum mocks base method
func (m *MockDatastore) GetHeadersByBlockNum(arg0 context.Context, arg1 uint64) ([]structs.HeaderAndTrace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHeadersByBlockNum", arg0, arg1)
	ret0, _ := ret[0].([]structs.HeaderAndTrace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHeadersByBlockNum indicates an expected call of GetHeadersByBlockNum
func (mr *MockDatastoreMockRecorder) GetHeadersByBlockNum(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHeadersByBlockNum", reflect.TypeOf((*MockDatastore)(nil).GetHeadersByBlockNum), arg0, arg1)
}

// GetHeadersBySlot mocks base method
func (m *MockDatastore) GetHeadersBySlot(arg0 context.Context, arg1 uint64) ([]structs.HeaderAndTrace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHeadersBySlot", arg0, arg1)
	ret0, _ := ret[0].([]structs.HeaderAndTrace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHeadersBySlot indicates an expected call of GetHeadersBySlot
func (mr *MockDatastoreMockRecorder) GetHeadersBySlot(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHeadersBySlot", reflect.TypeOf((*MockDatastore)(nil).GetHeadersBySlot), arg0, arg1)
}

// GetLatestHeaders mocks base method
func (m *MockDatastore) GetLatestHeaders(arg0 context.Context, arg1, arg2 uint64) ([]structs.HeaderAndTrace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLatestHeaders", arg0, arg1, arg2)
	ret0, _ := ret[0].([]structs.HeaderAndTrace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLatestHeaders indicates an expected call of GetLatestHeaders
func (mr *MockDatastoreMockRecorder) GetLatestHeaders(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLatestHeaders", reflect.TypeOf((*MockDatastore)(nil).GetLatestHeaders), arg0, arg1, arg2)
}

// GetMaxProfitHeader mocks base method
func (m *MockDatastore) GetMaxProfitHeader(arg0 context.Context, arg1 uint64) (structs.HeaderAndTrace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMaxProfitHeader", arg0, arg1)
	ret0, _ := ret[0].(structs.HeaderAndTrace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMaxProfitHeader indicates an expected call of GetMaxProfitHeader
func (mr *MockDatastoreMockRecorder) GetMaxProfitHeader(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMaxProfitHeader", reflect.TypeOf((*MockDatastore)(nil).GetMaxProfitHeader), arg0, arg1)
}

// GetPayload mocks base method
func (m *MockDatastore) GetPayload(arg0 context.Context, arg1 structs.PayloadKey) (*structs.BlockBidAndTrace, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPayload", arg0, arg1)
	ret0, _ := ret[0].(*structs.BlockBidAndTrace)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetPayload indicates an expected call of GetPayload
func (mr *MockDatastoreMockRecorder) GetPayload(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPayload", reflect.TypeOf((*MockDatastore)(nil).GetPayload), arg0, arg1)
}

// PutDelivered mocks base method
func (m *MockDatastore) PutDelivered(arg0 context.Context, arg1 structs.Slot, arg2 structs.DeliveredTrace, arg3 time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutDelivered", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutDelivered indicates an expected call of PutDelivered
func (mr *MockDatastoreMockRecorder) PutDelivered(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutDelivered", reflect.TypeOf((*MockDatastore)(nil).PutDelivered), arg0, arg1, arg2, arg3)
}

// PutHeader mocks base method
func (m *MockDatastore) PutHeader(arg0 context.Context, arg1 structs.HeaderData, arg2 time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutHeader", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutHeader indicates an expected call of PutHeader
func (mr *MockDatastoreMockRecorder) PutHeader(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutHeader", reflect.TypeOf((*MockDatastore)(nil).PutHeader), arg0, arg1, arg2)
}

// PutPayload mocks base method
func (m *MockDatastore) PutPayload(arg0 context.Context, arg1 structs.PayloadKey, arg2 *structs.BlockBidAndTrace, arg3 time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutPayload", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutPayload indicates an expected call of PutPayload
func (mr *MockDatastoreMockRecorder) PutPayload(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutPayload", reflect.TypeOf((*MockDatastore)(nil).PutPayload), arg0, arg1, arg2, arg3)
}

// MockState is a mock of State interface
type MockState struct {
	ctrl     *gomock.Controller
	recorder *MockStateMockRecorder
}

// MockStateMockRecorder is the mock recorder for MockState
type MockStateMockRecorder struct {
	mock *MockState
}

// NewMockState creates a new mock instance
func NewMockState(ctrl *gomock.Controller) *MockState {
	mock := &MockState{ctrl: ctrl}
	mock.recorder = &MockStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockState) EXPECT() *MockStateMockRecorder {
	return m.recorder
}

// Beacon mocks base method
func (m *MockState) Beacon() *structs.BeaconState {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Beacon")
	ret0, _ := ret[0].(*structs.BeaconState)
	return ret0
}

// Beacon indicates an expected call of Beacon
func (mr *MockStateMockRecorder) Beacon() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Beacon", reflect.TypeOf((*MockState)(nil).Beacon))
}
