// Copyright 2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of this
// software and associated documentation files (the "Software"), to deal in the Software
// without restriction, including without limitation the rights to use, copy, modify,
// merge, publish, distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED,
// INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A
// PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
// SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/containers/image/types (interfaces: ImageReference,ImageSource)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	io "io"
	reflect "reflect"

	reference "github.com/containers/image/docker/reference"
	types "github.com/containers/image/types"
	gomock "github.com/golang/mock/gomock"
	go_digest "github.com/opencontainers/go-digest"
)

// MockImageReference is a mock of ImageReference interface
type MockImageReference struct {
	ctrl     *gomock.Controller
	recorder *MockImageReferenceMockRecorder
}

// MockImageReferenceMockRecorder is the mock recorder for MockImageReference
type MockImageReferenceMockRecorder struct {
	mock *MockImageReference
}

// NewMockImageReference creates a new mock instance
func NewMockImageReference(ctrl *gomock.Controller) *MockImageReference {
	mock := &MockImageReference{ctrl: ctrl}
	mock.recorder = &MockImageReferenceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockImageReference) EXPECT() *MockImageReferenceMockRecorder {
	return m.recorder
}

// DeleteImage mocks base method
func (m *MockImageReference) DeleteImage(arg0 context.Context, arg1 *types.SystemContext) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteImage", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteImage indicates an expected call of DeleteImage
func (mr *MockImageReferenceMockRecorder) DeleteImage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteImage", reflect.TypeOf((*MockImageReference)(nil).DeleteImage), arg0, arg1)
}

// DockerReference mocks base method
func (m *MockImageReference) DockerReference() reference.Named {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DockerReference")
	ret0, _ := ret[0].(reference.Named)
	return ret0
}

// DockerReference indicates an expected call of DockerReference
func (mr *MockImageReferenceMockRecorder) DockerReference() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DockerReference", reflect.TypeOf((*MockImageReference)(nil).DockerReference))
}

// NewImage mocks base method
func (m *MockImageReference) NewImage(arg0 context.Context, arg1 *types.SystemContext) (types.ImageCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewImage", arg0, arg1)
	ret0, _ := ret[0].(types.ImageCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewImage indicates an expected call of NewImage
func (mr *MockImageReferenceMockRecorder) NewImage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewImage", reflect.TypeOf((*MockImageReference)(nil).NewImage), arg0, arg1)
}

// NewImageDestination mocks base method
func (m *MockImageReference) NewImageDestination(arg0 context.Context, arg1 *types.SystemContext) (types.ImageDestination, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewImageDestination", arg0, arg1)
	ret0, _ := ret[0].(types.ImageDestination)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewImageDestination indicates an expected call of NewImageDestination
func (mr *MockImageReferenceMockRecorder) NewImageDestination(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewImageDestination", reflect.TypeOf((*MockImageReference)(nil).NewImageDestination), arg0, arg1)
}

// NewImageSource mocks base method
func (m *MockImageReference) NewImageSource(arg0 context.Context, arg1 *types.SystemContext) (types.ImageSource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewImageSource", arg0, arg1)
	ret0, _ := ret[0].(types.ImageSource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewImageSource indicates an expected call of NewImageSource
func (mr *MockImageReferenceMockRecorder) NewImageSource(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewImageSource", reflect.TypeOf((*MockImageReference)(nil).NewImageSource), arg0, arg1)
}

// PolicyConfigurationIdentity mocks base method
func (m *MockImageReference) PolicyConfigurationIdentity() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PolicyConfigurationIdentity")
	ret0, _ := ret[0].(string)
	return ret0
}

// PolicyConfigurationIdentity indicates an expected call of PolicyConfigurationIdentity
func (mr *MockImageReferenceMockRecorder) PolicyConfigurationIdentity() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PolicyConfigurationIdentity", reflect.TypeOf((*MockImageReference)(nil).PolicyConfigurationIdentity))
}

// PolicyConfigurationNamespaces mocks base method
func (m *MockImageReference) PolicyConfigurationNamespaces() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PolicyConfigurationNamespaces")
	ret0, _ := ret[0].([]string)
	return ret0
}

// PolicyConfigurationNamespaces indicates an expected call of PolicyConfigurationNamespaces
func (mr *MockImageReferenceMockRecorder) PolicyConfigurationNamespaces() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PolicyConfigurationNamespaces", reflect.TypeOf((*MockImageReference)(nil).PolicyConfigurationNamespaces))
}

// StringWithinTransport mocks base method
func (m *MockImageReference) StringWithinTransport() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StringWithinTransport")
	ret0, _ := ret[0].(string)
	return ret0
}

// StringWithinTransport indicates an expected call of StringWithinTransport
func (mr *MockImageReferenceMockRecorder) StringWithinTransport() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StringWithinTransport", reflect.TypeOf((*MockImageReference)(nil).StringWithinTransport))
}

// Transport mocks base method
func (m *MockImageReference) Transport() types.ImageTransport {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Transport")
	ret0, _ := ret[0].(types.ImageTransport)
	return ret0
}

// Transport indicates an expected call of Transport
func (mr *MockImageReferenceMockRecorder) Transport() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Transport", reflect.TypeOf((*MockImageReference)(nil).Transport))
}

// MockImageSource is a mock of ImageSource interface
type MockImageSource struct {
	ctrl     *gomock.Controller
	recorder *MockImageSourceMockRecorder
}

// MockImageSourceMockRecorder is the mock recorder for MockImageSource
type MockImageSourceMockRecorder struct {
	mock *MockImageSource
}

// NewMockImageSource creates a new mock instance
func NewMockImageSource(ctrl *gomock.Controller) *MockImageSource {
	mock := &MockImageSource{ctrl: ctrl}
	mock.recorder = &MockImageSourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockImageSource) EXPECT() *MockImageSourceMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockImageSource) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockImageSourceMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockImageSource)(nil).Close))
}

// GetBlob mocks base method
func (m *MockImageSource) GetBlob(arg0 context.Context, arg1 types.BlobInfo, arg2 types.BlobInfoCache) (io.ReadCloser, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlob", arg0, arg1, arg2)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetBlob indicates an expected call of GetBlob
func (mr *MockImageSourceMockRecorder) GetBlob(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlob", reflect.TypeOf((*MockImageSource)(nil).GetBlob), arg0, arg1, arg2)
}

// GetManifest mocks base method
func (m *MockImageSource) GetManifest(arg0 context.Context, arg1 *go_digest.Digest) ([]byte, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetManifest", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetManifest indicates an expected call of GetManifest
func (mr *MockImageSourceMockRecorder) GetManifest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetManifest", reflect.TypeOf((*MockImageSource)(nil).GetManifest), arg0, arg1)
}

// GetSignatures mocks base method
func (m *MockImageSource) GetSignatures(arg0 context.Context, arg1 *go_digest.Digest) ([][]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSignatures", arg0, arg1)
	ret0, _ := ret[0].([][]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSignatures indicates an expected call of GetSignatures
func (mr *MockImageSourceMockRecorder) GetSignatures(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSignatures", reflect.TypeOf((*MockImageSource)(nil).GetSignatures), arg0, arg1)
}

// HasThreadSafeGetBlob mocks base method
func (m *MockImageSource) HasThreadSafeGetBlob() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasThreadSafeGetBlob")
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasThreadSafeGetBlob indicates an expected call of HasThreadSafeGetBlob
func (mr *MockImageSourceMockRecorder) HasThreadSafeGetBlob() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasThreadSafeGetBlob", reflect.TypeOf((*MockImageSource)(nil).HasThreadSafeGetBlob))
}

// LayerInfosForCopy mocks base method
func (m *MockImageSource) LayerInfosForCopy(arg0 context.Context) ([]types.BlobInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LayerInfosForCopy", arg0)
	ret0, _ := ret[0].([]types.BlobInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LayerInfosForCopy indicates an expected call of LayerInfosForCopy
func (mr *MockImageSourceMockRecorder) LayerInfosForCopy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LayerInfosForCopy", reflect.TypeOf((*MockImageSource)(nil).LayerInfosForCopy), arg0)
}

// Reference mocks base method
func (m *MockImageSource) Reference() types.ImageReference {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reference")
	ret0, _ := ret[0].(types.ImageReference)
	return ret0
}

// Reference indicates an expected call of Reference
func (mr *MockImageSourceMockRecorder) Reference() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reference", reflect.TypeOf((*MockImageSource)(nil).Reference))
}
