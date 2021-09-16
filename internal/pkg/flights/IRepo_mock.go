// Code generated by MockGen. DO NOT EDIT.
// Source: IRepo.go

// Package flights is a generated GoMock package.
package flights

import (
	models "RSOI_CW/internal/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockIRepo is a mock of IRepo interface.
type MockIRepo struct {
	ctrl     *gomock.Controller
	recorder *MockIRepoMockRecorder
}

// MockIRepoMockRecorder is the mock recorder for MockIRepo.
type MockIRepoMockRecorder struct {
	mock *MockIRepo
}

// NewMockIRepo creates a new mock instance.
func NewMockIRepo(ctrl *gomock.Controller) *MockIRepo {
	mock := &MockIRepo{ctrl: ctrl}
	mock.recorder = &MockIRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRepo) EXPECT() *MockIRepoMockRecorder {
	return m.recorder
}

// CreateFlight mocks base method.
func (m *MockIRepo) CreateFlight(flight models.Flight) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFlight", flight)
	ret0, _ := ret[0].(int)
	return ret0
}

// CreateFlight indicates an expected call of CreateFlight.
func (mr *MockIRepoMockRecorder) CreateFlight(flight interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFlight", reflect.TypeOf((*MockIRepo)(nil).CreateFlight), flight)
}

// ReadFlight mocks base method.
func (m *MockIRepo) ReadFlight(FlightUUID uuid.UUID) (models.Flight, int) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadFlight", FlightUUID)
	ret0, _ := ret[0].(models.Flight)
	ret1, _ := ret[1].(int)
	return ret0, ret1
}

// ReadFlight indicates an expected call of ReadFlight.
func (mr *MockIRepoMockRecorder) ReadFlight(FlightUUID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadFlight", reflect.TypeOf((*MockIRepo)(nil).ReadFlight), FlightUUID)
}

// ReadFlights mocks base method.
func (m *MockIRepo) ReadFlights() ([]models.Flight, int) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadFlights")
	ret0, _ := ret[0].([]models.Flight)
	ret1, _ := ret[1].(int)
	return ret0, ret1
}

// ReadFlights indicates an expected call of ReadFlights.
func (mr *MockIRepoMockRecorder) ReadFlights() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadFlights", reflect.TypeOf((*MockIRepo)(nil).ReadFlights))
}

// UpdateFlight mocks base method.
func (m *MockIRepo) UpdateFlight(FlightUUID uuid.UUID, flight models.Flight) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFlight", FlightUUID, flight)
	ret0, _ := ret[0].(int)
	return ret0
}

// UpdateFlight indicates an expected call of UpdateFlight.
func (mr *MockIRepoMockRecorder) UpdateFlight(FlightUUID, flight interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFlight", reflect.TypeOf((*MockIRepo)(nil).UpdateFlight), FlightUUID, flight)
}
