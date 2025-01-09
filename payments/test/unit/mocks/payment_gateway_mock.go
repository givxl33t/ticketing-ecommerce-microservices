// Code generated by MockGen. DO NOT EDIT.
// Source: internal/infrastructure/stripe.go
//
// Generated by this command:
//
//	mockgen -source=internal/infrastructure/stripe.go -destination=test/unit/mocks/payment_gateway_mock.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	domain "ticketing/payments/internal/domain"

	stripe "github.com/stripe/stripe-go/v74"
	gomock "go.uber.org/mock/gomock"
)

// MockPaymentGateway is a mock of PaymentGateway interface.
type MockPaymentGateway struct {
	ctrl     *gomock.Controller
	recorder *MockPaymentGatewayMockRecorder
	isgomock struct{}
}

// MockPaymentGatewayMockRecorder is the mock recorder for MockPaymentGateway.
type MockPaymentGatewayMockRecorder struct {
	mock *MockPaymentGateway
}

// NewMockPaymentGateway creates a new mock instance.
func NewMockPaymentGateway(ctrl *gomock.Controller) *MockPaymentGateway {
	mock := &MockPaymentGateway{ctrl: ctrl}
	mock.recorder = &MockPaymentGatewayMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPaymentGateway) EXPECT() *MockPaymentGatewayMockRecorder {
	return m.recorder
}

// CreatePayment mocks base method.
func (m *MockPaymentGateway) CreatePayment(order *domain.Order) (*stripe.PaymentIntent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePayment", order)
	ret0, _ := ret[0].(*stripe.PaymentIntent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePayment indicates an expected call of CreatePayment.
func (mr *MockPaymentGatewayMockRecorder) CreatePayment(order any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePayment", reflect.TypeOf((*MockPaymentGateway)(nil).CreatePayment), order)
}