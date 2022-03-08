// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	entities "github.com/OmarAouini/golang-api-starter/entities"
	mock "github.com/stretchr/testify/mock"
)

// ProjectStore is an autogenerated mock type for the ProjectStore type
type ProjectStore struct {
	mock.Mock
}

// Get provides a mock function with given fields: id
func (_m *ProjectStore) Get(id int) (*entities.Project, error) {
	ret := _m.Called(id)

	var r0 *entities.Project
	if rf, ok := ret.Get(0).(func(int) *entities.Project); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Project)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByName provides a mock function with given fields: name
func (_m *ProjectStore) GetByName(name string) (*entities.Project, error) {
	ret := _m.Called(name)

	var r0 *entities.Project
	if rf, ok := ret.Get(0).(func(string) *entities.Project); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Project)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}