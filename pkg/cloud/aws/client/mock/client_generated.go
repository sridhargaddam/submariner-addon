// Code generated by MockGen. DO NOT EDIT.
// Source: ./client.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	ec2 "github.com/aws/aws-sdk-go/service/ec2"
	ec2iface "github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	gomock "github.com/golang/mock/gomock"
)

// MockInterface is a mock of Interface interface.
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface.
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance.
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// AuthorizeSecurityGroupIngress mocks base method.
func (m *MockInterface) AuthorizeSecurityGroupIngress(input *ec2.AuthorizeSecurityGroupIngressInput) (*ec2.AuthorizeSecurityGroupIngressOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthorizeSecurityGroupIngress", input)
	ret0, _ := ret[0].(*ec2.AuthorizeSecurityGroupIngressOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthorizeSecurityGroupIngress indicates an expected call of AuthorizeSecurityGroupIngress.
func (mr *MockInterfaceMockRecorder) AuthorizeSecurityGroupIngress(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthorizeSecurityGroupIngress", reflect.TypeOf((*MockInterface)(nil).AuthorizeSecurityGroupIngress), input)
}

// CreateSecurityGroup mocks base method.
func (m *MockInterface) CreateSecurityGroup(input *ec2.CreateSecurityGroupInput) (*ec2.CreateSecurityGroupOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSecurityGroup", input)
	ret0, _ := ret[0].(*ec2.CreateSecurityGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSecurityGroup indicates an expected call of CreateSecurityGroup.
func (mr *MockInterfaceMockRecorder) CreateSecurityGroup(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSecurityGroup", reflect.TypeOf((*MockInterface)(nil).CreateSecurityGroup), input)
}

// CreateTags mocks base method.
func (m *MockInterface) CreateTags(input *ec2.CreateTagsInput) (*ec2.CreateTagsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTags", input)
	ret0, _ := ret[0].(*ec2.CreateTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTags indicates an expected call of CreateTags.
func (mr *MockInterfaceMockRecorder) CreateTags(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTags", reflect.TypeOf((*MockInterface)(nil).CreateTags), input)
}

// DeleteTags mocks base method.
func (m *MockInterface) DeleteTags(input *ec2.DeleteTagsInput) (*ec2.DeleteTagsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTags", input)
	ret0, _ := ret[0].(*ec2.DeleteTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteTags indicates an expected call of DeleteTags.
func (mr *MockInterfaceMockRecorder) DeleteTags(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTags", reflect.TypeOf((*MockInterface)(nil).DeleteTags), input)
}

// DescribeInstanceTypeOfferings mocks base method.
func (m *MockInterface) DescribeInstanceTypeOfferings(input *ec2.DescribeInstanceTypeOfferingsInput) (*ec2.DescribeInstanceTypeOfferingsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeInstanceTypeOfferings", input)
	ret0, _ := ret[0].(*ec2.DescribeInstanceTypeOfferingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeInstanceTypeOfferings indicates an expected call of DescribeInstanceTypeOfferings.
func (mr *MockInterfaceMockRecorder) DescribeInstanceTypeOfferings(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeInstanceTypeOfferings", reflect.TypeOf((*MockInterface)(nil).DescribeInstanceTypeOfferings), input)
}

// DescribeInstances mocks base method.
func (m *MockInterface) DescribeInstances(input *ec2.DescribeInstancesInput) (*ec2.DescribeInstancesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeInstances", input)
	ret0, _ := ret[0].(*ec2.DescribeInstancesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeInstances indicates an expected call of DescribeInstances.
func (mr *MockInterfaceMockRecorder) DescribeInstances(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeInstances", reflect.TypeOf((*MockInterface)(nil).DescribeInstances), input)
}

// DescribeSecurityGroups mocks base method.
func (m *MockInterface) DescribeSecurityGroups(input *ec2.DescribeSecurityGroupsInput) (*ec2.DescribeSecurityGroupsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeSecurityGroups", input)
	ret0, _ := ret[0].(*ec2.DescribeSecurityGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeSecurityGroups indicates an expected call of DescribeSecurityGroups.
func (mr *MockInterfaceMockRecorder) DescribeSecurityGroups(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeSecurityGroups", reflect.TypeOf((*MockInterface)(nil).DescribeSecurityGroups), input)
}

// DescribeSubnets mocks base method.
func (m *MockInterface) DescribeSubnets(input *ec2.DescribeSubnetsInput) (*ec2.DescribeSubnetsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeSubnets", input)
	ret0, _ := ret[0].(*ec2.DescribeSubnetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeSubnets indicates an expected call of DescribeSubnets.
func (mr *MockInterfaceMockRecorder) DescribeSubnets(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeSubnets", reflect.TypeOf((*MockInterface)(nil).DescribeSubnets), input)
}

// DescribeVpcs mocks base method.
func (m *MockInterface) DescribeVpcs(input *ec2.DescribeVpcsInput) (*ec2.DescribeVpcsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeVpcs", input)
	ret0, _ := ret[0].(*ec2.DescribeVpcsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeVpcs indicates an expected call of DescribeVpcs.
func (mr *MockInterfaceMockRecorder) DescribeVpcs(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeVpcs", reflect.TypeOf((*MockInterface)(nil).DescribeVpcs), input)
}

// FullEC2API mocks base method.
func (m *MockInterface) FullEC2API() ec2iface.EC2API {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FullEC2API")
	ret0, _ := ret[0].(ec2iface.EC2API)
	return ret0
}

// FullEC2API indicates an expected call of FullEC2API.
func (mr *MockInterfaceMockRecorder) FullEC2API() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FullEC2API", reflect.TypeOf((*MockInterface)(nil).FullEC2API))
}

// RevokeSecurityGroupIngress mocks base method.
func (m *MockInterface) RevokeSecurityGroupIngress(input *ec2.RevokeSecurityGroupIngressInput) (*ec2.RevokeSecurityGroupIngressOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RevokeSecurityGroupIngress", input)
	ret0, _ := ret[0].(*ec2.RevokeSecurityGroupIngressOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RevokeSecurityGroupIngress indicates an expected call of RevokeSecurityGroupIngress.
func (mr *MockInterfaceMockRecorder) RevokeSecurityGroupIngress(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokeSecurityGroupIngress", reflect.TypeOf((*MockInterface)(nil).RevokeSecurityGroupIngress), input)
}
