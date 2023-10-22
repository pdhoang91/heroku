package mocks

import (
	model "heroku/internal/delivery/http/model"

	mock "github.com/stretchr/testify/mock"
)

type UserUseCase struct {
	mock.Mock
}

func (_m *UserUseCase) GetUserInfo(id int) (*model.UserInfo, error) {
	ret := _m.Called(id)

	var r0 *model.UserInfo
	if rf, ok := ret.Get(0).(func(int) *model.UserInfo); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.UserInfo)
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

func (_m *UserUseCase) GetAllUserInfo() ([]*model.UserInfo, error) {
	ret := _m.Called()

	var r0 []*model.UserInfo
	if rf, ok := ret.Get(0).(func() []*model.UserInfo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.UserInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
