package stores

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "gofr.dev/ZopSmart_GoFr_Project/models"
	gofr "gofr.dev/pkg/gofr"
)

// MockLibrary is a mock of Library interface.
type MockLibrary struct {
	ctrl     *gomock.Controller
	recorder *MockLibraryMockRecorder
}

// MockLibraryMockRecorder is the mock recorder for MockLibrary.
type MockLibraryMockRecorder struct {
	mock *MockLibrary
}

// NewMockLibrary creates a new mock instance.
func NewMockLibrary(ctrl *gomock.Controller) *MockLibrary {
	mock := &MockLibrary{ctrl: ctrl}
	mock.recorder = &MockLibraryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLibrary) EXPECT() *MockLibraryMockRecorder {
	return m.recorder
}

// Create mocks the Create method.
func (m *MockLibrary) Create(ctx *gofr.Context, book models.Book) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, book)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockLibraryMockRecorder) Create(ctx, book interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockLibrary)(nil).Create), ctx, book)
}

// Delete mocks the Delete method.
func (m *MockLibrary) Delete(ctx *gofr.Context, id string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockLibraryMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockLibrary)(nil).Delete), ctx, id)
}

// Get mocks the Get method.
func (m *MockLibrary) Get(ctx *gofr.Context, id string) (models.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, id)
	ret0, _ := ret[0].(models.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockLibraryMockRecorder) Get(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockLibrary)(nil).Get), ctx, id)
}
