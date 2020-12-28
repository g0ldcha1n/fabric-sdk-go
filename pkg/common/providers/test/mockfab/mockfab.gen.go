// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/huijinchain/fabric-sdk-go/pkg/common/providers/fab (interfaces: EndpointConfig,ProposalProcessor,Providers)

// Package mockfab is a generated GoMock package.
package mockfab

import (
	context "context"
	tls "crypto/tls"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	fab "github.com/huijinchain/fabric-sdk-go/pkg/common/providers/fab"
	tls0 "github.com/huijinchain/fabric-sdk-go/pkg/core/config/comm/tls"
	metrics "github.com/huijinchain/fabric-sdk-go/pkg/fabsdk/metrics"
)

// MockEndpointConfig is a mock of EndpointConfig interface
type MockEndpointConfig struct {
	ctrl     *gomock.Controller
	recorder *MockEndpointConfigMockRecorder
}

// MockEndpointConfigMockRecorder is the mock recorder for MockEndpointConfig
type MockEndpointConfigMockRecorder struct {
	mock *MockEndpointConfig
}

// NewMockEndpointConfig creates a new mock instance
func NewMockEndpointConfig(ctrl *gomock.Controller) *MockEndpointConfig {
	mock := &MockEndpointConfig{ctrl: ctrl}
	mock.recorder = &MockEndpointConfigMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEndpointConfig) EXPECT() *MockEndpointConfigMockRecorder {
	return m.recorder
}

// ChannelConfig mocks base method
func (m *MockEndpointConfig) ChannelConfig(arg0 string) *fab.ChannelEndpointConfig {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChannelConfig", arg0)
	ret0, _ := ret[0].(*fab.ChannelEndpointConfig)
	return ret0
}

// ChannelConfig indicates an expected call of ChannelConfig
func (mr *MockEndpointConfigMockRecorder) ChannelConfig(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChannelConfig", reflect.TypeOf((*MockEndpointConfig)(nil).ChannelConfig), arg0)
}

// ChannelOrderers mocks base method
func (m *MockEndpointConfig) ChannelOrderers(arg0 string) []fab.OrdererConfig {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChannelOrderers", arg0)
	ret0, _ := ret[0].([]fab.OrdererConfig)
	return ret0
}

// ChannelOrderers indicates an expected call of ChannelOrderers
func (mr *MockEndpointConfigMockRecorder) ChannelOrderers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChannelOrderers", reflect.TypeOf((*MockEndpointConfig)(nil).ChannelOrderers), arg0)
}

// ChannelPeers mocks base method
func (m *MockEndpointConfig) ChannelPeers(arg0 string) []fab.ChannelPeer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChannelPeers", arg0)
	ret0, _ := ret[0].([]fab.ChannelPeer)
	return ret0
}

// ChannelPeers indicates an expected call of ChannelPeers
func (mr *MockEndpointConfigMockRecorder) ChannelPeers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChannelPeers", reflect.TypeOf((*MockEndpointConfig)(nil).ChannelPeers), arg0)
}

// CryptoConfigPath mocks base method
func (m *MockEndpointConfig) CryptoConfigPath() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CryptoConfigPath")
	ret0, _ := ret[0].(string)
	return ret0
}

// CryptoConfigPath indicates an expected call of CryptoConfigPath
func (mr *MockEndpointConfigMockRecorder) CryptoConfigPath() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CryptoConfigPath", reflect.TypeOf((*MockEndpointConfig)(nil).CryptoConfigPath))
}

// NetworkConfig mocks base method
func (m *MockEndpointConfig) NetworkConfig() *fab.NetworkConfig {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NetworkConfig")
	ret0, _ := ret[0].(*fab.NetworkConfig)
	return ret0
}

// NetworkConfig indicates an expected call of NetworkConfig
func (mr *MockEndpointConfigMockRecorder) NetworkConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetworkConfig", reflect.TypeOf((*MockEndpointConfig)(nil).NetworkConfig))
}

// NetworkPeers mocks base method
func (m *MockEndpointConfig) NetworkPeers() []fab.NetworkPeer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NetworkPeers")
	ret0, _ := ret[0].([]fab.NetworkPeer)
	return ret0
}

// NetworkPeers indicates an expected call of NetworkPeers
func (mr *MockEndpointConfigMockRecorder) NetworkPeers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetworkPeers", reflect.TypeOf((*MockEndpointConfig)(nil).NetworkPeers))
}

// OrdererConfig mocks base method
func (m *MockEndpointConfig) OrdererConfig(arg0 string) (*fab.OrdererConfig, bool, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OrdererConfig", arg0)
	ret0, _ := ret[0].(*fab.OrdererConfig)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(bool)
	return ret0, ret1, ret2
}

// OrdererConfig indicates an expected call of OrdererConfig
func (mr *MockEndpointConfigMockRecorder) OrdererConfig(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrdererConfig", reflect.TypeOf((*MockEndpointConfig)(nil).OrdererConfig), arg0)
}

// OrderersConfig mocks base method
func (m *MockEndpointConfig) OrderersConfig() []fab.OrdererConfig {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OrderersConfig")
	ret0, _ := ret[0].([]fab.OrdererConfig)
	return ret0
}

// OrderersConfig indicates an expected call of OrderersConfig
func (mr *MockEndpointConfigMockRecorder) OrderersConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrderersConfig", reflect.TypeOf((*MockEndpointConfig)(nil).OrderersConfig))
}

// PeerConfig mocks base method
func (m *MockEndpointConfig) PeerConfig(arg0 string) (*fab.PeerConfig, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PeerConfig", arg0)
	ret0, _ := ret[0].(*fab.PeerConfig)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// PeerConfig indicates an expected call of PeerConfig
func (mr *MockEndpointConfigMockRecorder) PeerConfig(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PeerConfig", reflect.TypeOf((*MockEndpointConfig)(nil).PeerConfig), arg0)
}

// PeersConfig mocks base method
func (m *MockEndpointConfig) PeersConfig(arg0 string) ([]fab.PeerConfig, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PeersConfig", arg0)
	ret0, _ := ret[0].([]fab.PeerConfig)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// PeersConfig indicates an expected call of PeersConfig
func (mr *MockEndpointConfigMockRecorder) PeersConfig(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PeersConfig", reflect.TypeOf((*MockEndpointConfig)(nil).PeersConfig), arg0)
}

// TLSCACertPool mocks base method
func (m *MockEndpointConfig) TLSCACertPool() tls0.CertPool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TLSCACertPool")
	ret0, _ := ret[0].(tls0.CertPool)
	return ret0
}

// TLSCACertPool indicates an expected call of TLSCACertPool
func (mr *MockEndpointConfigMockRecorder) TLSCACertPool() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TLSCACertPool", reflect.TypeOf((*MockEndpointConfig)(nil).TLSCACertPool))
}

// TLSClientCerts mocks base method
func (m *MockEndpointConfig) TLSClientCerts() []tls.Certificate {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TLSClientCerts")
	ret0, _ := ret[0].([]tls.Certificate)
	return ret0
}

// TLSClientCerts indicates an expected call of TLSClientCerts
func (mr *MockEndpointConfigMockRecorder) TLSClientCerts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TLSClientCerts", reflect.TypeOf((*MockEndpointConfig)(nil).TLSClientCerts))
}

// Timeout mocks base method
func (m *MockEndpointConfig) Timeout(arg0 fab.TimeoutType) time.Duration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Timeout", arg0)
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// Timeout indicates an expected call of Timeout
func (mr *MockEndpointConfigMockRecorder) Timeout(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Timeout", reflect.TypeOf((*MockEndpointConfig)(nil).Timeout), arg0)
}

// MockProposalProcessor is a mock of ProposalProcessor interface
type MockProposalProcessor struct {
	ctrl     *gomock.Controller
	recorder *MockProposalProcessorMockRecorder
}

// MockProposalProcessorMockRecorder is the mock recorder for MockProposalProcessor
type MockProposalProcessorMockRecorder struct {
	mock *MockProposalProcessor
}

// NewMockProposalProcessor creates a new mock instance
func NewMockProposalProcessor(ctrl *gomock.Controller) *MockProposalProcessor {
	mock := &MockProposalProcessor{ctrl: ctrl}
	mock.recorder = &MockProposalProcessorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProposalProcessor) EXPECT() *MockProposalProcessorMockRecorder {
	return m.recorder
}

// ProcessTransactionProposal mocks base method
func (m *MockProposalProcessor) ProcessTransactionProposal(arg0 context.Context, arg1 fab.ProcessProposalRequest) (*fab.TransactionProposalResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessTransactionProposal", arg0, arg1)
	ret0, _ := ret[0].(*fab.TransactionProposalResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProcessTransactionProposal indicates an expected call of ProcessTransactionProposal
func (mr *MockProposalProcessorMockRecorder) ProcessTransactionProposal(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessTransactionProposal", reflect.TypeOf((*MockProposalProcessor)(nil).ProcessTransactionProposal), arg0, arg1)
}

// MockProviders is a mock of Providers interface
type MockProviders struct {
	ctrl     *gomock.Controller
	recorder *MockProvidersMockRecorder
}

// MockProvidersMockRecorder is the mock recorder for MockProviders
type MockProvidersMockRecorder struct {
	mock *MockProviders
}

// NewMockProviders creates a new mock instance
func NewMockProviders(ctrl *gomock.Controller) *MockProviders {
	mock := &MockProviders{ctrl: ctrl}
	mock.recorder = &MockProvidersMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProviders) EXPECT() *MockProvidersMockRecorder {
	return m.recorder
}

// ChannelProvider mocks base method
func (m *MockProviders) ChannelProvider() fab.ChannelProvider {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChannelProvider")
	ret0, _ := ret[0].(fab.ChannelProvider)
	return ret0
}

// ChannelProvider indicates an expected call of ChannelProvider
func (mr *MockProvidersMockRecorder) ChannelProvider() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChannelProvider", reflect.TypeOf((*MockProviders)(nil).ChannelProvider))
}

// EndpointConfig mocks base method
func (m *MockProviders) EndpointConfig() fab.EndpointConfig {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EndpointConfig")
	ret0, _ := ret[0].(fab.EndpointConfig)
	return ret0
}

// EndpointConfig indicates an expected call of EndpointConfig
func (mr *MockProvidersMockRecorder) EndpointConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EndpointConfig", reflect.TypeOf((*MockProviders)(nil).EndpointConfig))
}

// GetMetrics mocks base method
func (m *MockProviders) GetMetrics() *metrics.ClientMetrics {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMetrics")
	ret0, _ := ret[0].(*metrics.ClientMetrics)
	return ret0
}

// GetMetrics indicates an expected call of GetMetrics
func (mr *MockProvidersMockRecorder) GetMetrics() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMetrics", reflect.TypeOf((*MockProviders)(nil).GetMetrics))
}

// InfraProvider mocks base method
func (m *MockProviders) InfraProvider() fab.InfraProvider {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InfraProvider")
	ret0, _ := ret[0].(fab.InfraProvider)
	return ret0
}

// InfraProvider indicates an expected call of InfraProvider
func (mr *MockProvidersMockRecorder) InfraProvider() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InfraProvider", reflect.TypeOf((*MockProviders)(nil).InfraProvider))
}

// LocalDiscoveryProvider mocks base method
func (m *MockProviders) LocalDiscoveryProvider() fab.LocalDiscoveryProvider {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LocalDiscoveryProvider")
	ret0, _ := ret[0].(fab.LocalDiscoveryProvider)
	return ret0
}

// LocalDiscoveryProvider indicates an expected call of LocalDiscoveryProvider
func (mr *MockProvidersMockRecorder) LocalDiscoveryProvider() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LocalDiscoveryProvider", reflect.TypeOf((*MockProviders)(nil).LocalDiscoveryProvider))
}
