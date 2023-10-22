package mocks

import (
	"heroku/internal/entities"

	"github.com/stretchr/testify/mock"
)

type UserRepository struct {
	mock.Mock
}

func (uc *UserRepository) FindUserByID(userID int) (*entities.User, error) {
	ret := uc.Called(userID)

	var r0 *entities.User
	if rf, ok := ret.Get(0).(func(int) *entities.User); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (uc *UserRepository) GetUserAccounts(userID int) ([]*entities.Account, error) {
	ret := uc.Called(userID)

	var r0 []*entities.Account
	if rf, ok := ret.Get(0).(func(int) []*entities.Account); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (uc *UserRepository) GetAllUser() ([]*entities.User, error) {
	ret := uc.Called()

	var r0 []*entities.User
	if rf, ok := ret.Get(0).(func() []*entities.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.User)
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
