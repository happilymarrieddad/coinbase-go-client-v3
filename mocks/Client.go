// Code generated by MockGen. DO NOT EDIT.
// Source: coinbase-go-client-v3 (interfaces: Client)

// Package mock_client is a generated GoMock package.
package mock_client

import (
	coinbasegoclientv3 "coinbase-go-client-v3"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// GetAccount mocks base method.
func (m *MockClient) GetAccount(arg0 context.Context, arg1 string) (*coinbasegoclientv3.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", arg0, arg1)
	ret0, _ := ret[0].(*coinbasegoclientv3.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccount indicates an expected call of GetAccount.
func (mr *MockClientMockRecorder) GetAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockClient)(nil).GetAccount), arg0, arg1)
}

// GetMarketTrades mocks base method.
func (m *MockClient) GetMarketTrades(arg0 context.Context, arg1 string, arg2 int) ([]*coinbasegoclientv3.MarketTrade, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMarketTrades", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*coinbasegoclientv3.MarketTrade)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMarketTrades indicates an expected call of GetMarketTrades.
func (mr *MockClientMockRecorder) GetMarketTrades(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMarketTrades", reflect.TypeOf((*MockClient)(nil).GetMarketTrades), arg0, arg1, arg2)
}

// GetProduct mocks base method.
func (m *MockClient) GetProduct(arg0 context.Context, arg1 string) (*coinbasegoclientv3.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProduct", arg0, arg1)
	ret0, _ := ret[0].(*coinbasegoclientv3.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProduct indicates an expected call of GetProduct.
func (mr *MockClientMockRecorder) GetProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProduct", reflect.TypeOf((*MockClient)(nil).GetProduct), arg0, arg1)
}

// GetProductCandles mocks base method.
func (m *MockClient) GetProductCandles(arg0 context.Context, arg1, arg2, arg3 string, arg4 coinbasegoclientv3.Granularity) ([]*coinbasegoclientv3.ProductCandle, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductCandles", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].([]*coinbasegoclientv3.ProductCandle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductCandles indicates an expected call of GetProductCandles.
func (mr *MockClientMockRecorder) GetProductCandles(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductCandles", reflect.TypeOf((*MockClient)(nil).GetProductCandles), arg0, arg1, arg2, arg3, arg4)
}

// GetTransactionHistory mocks base method.
func (m *MockClient) GetTransactionHistory(arg0 context.Context, arg1, arg2, arg3 string, arg4 coinbasegoclientv3.ProductType) (*coinbasegoclientv3.TransactionSummary, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionHistory", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*coinbasegoclientv3.TransactionSummary)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactionHistory indicates an expected call of GetTransactionHistory.
func (mr *MockClientMockRecorder) GetTransactionHistory(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionHistory", reflect.TypeOf((*MockClient)(nil).GetTransactionHistory), arg0, arg1, arg2, arg3, arg4)
}

// IsActive mocks base method.
func (m *MockClient) IsActive(arg0 context.Context) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsActive", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsActive indicates an expected call of IsActive.
func (mr *MockClientMockRecorder) IsActive(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsActive", reflect.TypeOf((*MockClient)(nil).IsActive), arg0)
}

// ListAccounts mocks base method.
func (m *MockClient) ListAccounts(arg0 context.Context) ([]*coinbasegoclientv3.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAccounts", arg0)
	ret0, _ := ret[0].([]*coinbasegoclientv3.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAccounts indicates an expected call of ListAccounts.
func (mr *MockClientMockRecorder) ListAccounts(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccounts", reflect.TypeOf((*MockClient)(nil).ListAccounts), arg0)
}

// ListOrders mocks base method.
func (m *MockClient) ListOrders(arg0 context.Context, arg1, arg2, arg3 string, arg4 coinbasegoclientv3.OrderType, arg5 coinbasegoclientv3.ProductType, arg6 *coinbasegoclientv3.ListOrdersOpts) (*coinbasegoclientv3.ListOrderData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListOrders", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(*coinbasegoclientv3.ListOrderData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOrders indicates an expected call of ListOrders.
func (mr *MockClientMockRecorder) ListOrders(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOrders", reflect.TypeOf((*MockClient)(nil).ListOrders), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// ListProducts mocks base method.
func (m *MockClient) ListProducts(arg0 context.Context) ([]*coinbasegoclientv3.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProducts", arg0)
	ret0, _ := ret[0].([]*coinbasegoclientv3.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProducts indicates an expected call of ListProducts.
func (mr *MockClientMockRecorder) ListProducts(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProducts", reflect.TypeOf((*MockClient)(nil).ListProducts), arg0)
}
