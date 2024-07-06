// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// PasswordUtilityInterface is an autogenerated mock type for the PasswordUtilityInterface type
type PasswordUtilityInterface struct {
	mock.Mock
}

// CheckPassword provides a mock function with given fields: inputPw, currentPw
func (_m *PasswordUtilityInterface) CheckPassword(inputPw []byte, currentPw []byte) error {
	ret := _m.Called(inputPw, currentPw)

	if len(ret) == 0 {
		panic("no return value specified for CheckPassword")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte, []byte) error); ok {
		r0 = rf(inputPw, currentPw)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GeneratePassword provides a mock function with given fields: currentPw
func (_m *PasswordUtilityInterface) GeneratePassword(currentPw string) ([]byte, error) {
	ret := _m.Called(currentPw)

	if len(ret) == 0 {
		panic("no return value specified for GeneratePassword")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]byte, error)); ok {
		return rf(currentPw)
	}
	if rf, ok := ret.Get(0).(func(string) []byte); ok {
		r0 = rf(currentPw)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(currentPw)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewPasswordUtilityInterface creates a new instance of PasswordUtilityInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPasswordUtilityInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *PasswordUtilityInterface {
	mock := &PasswordUtilityInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}