// Code generated by mockery v2.46.0. DO NOT EDIT.

package mocks

import (
	context "context"
	entity "user/internal/entity"

	mock "github.com/stretchr/testify/mock"

	repository "user/internal/repository"

	time "time"

	uuid "github.com/google/uuid"
)

// UserUseCase is an autogenerated mock type for the UserUseCase type
type UserUseCase struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: ctx, user
func (_m *UserUseCase) CreateUser(ctx context.Context, user entity.User) error {
	ret := _m.Called(ctx, user)

	if len(ret) == 0 {
		panic("no return value specified for CreateUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.User) error); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteUser provides a mock function with given fields: ctx, id
func (_m *UserUseCase) DeleteUser(ctx context.Context, id uuid.UUID) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetNewUsers provides a mock function with given fields: ctx, from, to
func (_m *UserUseCase) GetNewUsers(ctx context.Context, from time.Time, to time.Time) ([]entity.User, error) {
	ret := _m.Called(ctx, from, to)

	if len(ret) == 0 {
		panic("no return value specified for GetNewUsers")
	}

	var r0 []entity.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, time.Time, time.Time) ([]entity.User, error)); ok {
		return rf(ctx, from, to)
	}
	if rf, ok := ret.Get(0).(func(context.Context, time.Time, time.Time) []entity.User); ok {
		r0 = rf(ctx, from, to)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, time.Time, time.Time) error); ok {
		r1 = rf(ctx, from, to)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByID provides a mock function with given fields: ctx, id
func (_m *UserUseCase) GetUserByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByID")
	}

	var r0 *entity.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (*entity.User, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *entity.User); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUsers provides a mock function with given fields: ctx, filter
func (_m *UserUseCase) GetUsers(ctx context.Context, filter repository.UserFilter) ([]entity.User, error) {
	ret := _m.Called(ctx, filter)

	if len(ret) == 0 {
		panic("no return value specified for GetUsers")
	}

	var r0 []entity.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, repository.UserFilter) ([]entity.User, error)); ok {
		return rf(ctx, filter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, repository.UserFilter) []entity.User); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, repository.UserFilter) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUsersBatch provides a mock function with given fields: ctx, limit, offset
func (_m *UserUseCase) GetUsersBatch(ctx context.Context, limit int, offset int) ([]entity.User, error) {
	ret := _m.Called(ctx, limit, offset)

	if len(ret) == 0 {
		panic("no return value specified for GetUsersBatch")
	}

	var r0 []entity.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, int) ([]entity.User, error)); ok {
		return rf(ctx, limit, offset)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, int) []entity.User); ok {
		r0 = rf(ctx, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, int) error); ok {
		r1 = rf(ctx, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUser provides a mock function with given fields: ctx, user
func (_m *UserUseCase) UpdateUser(ctx context.Context, user *entity.User) error {
	ret := _m.Called(ctx, user)

	if len(ret) == 0 {
		panic("no return value specified for UpdateUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.User) error); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewUserUseCase creates a new instance of UserUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserUseCase(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserUseCase {
	mock := &UserUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}