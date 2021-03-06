// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import auth "github.com/erickoliv/finances-api/auth"
import context "context"
import mock "github.com/stretchr/testify/mock"

// Authenticator is an autogenerated mock type for the Authenticator type
type Authenticator struct {
	mock.Mock
}

// Login provides a mock function with given fields: ctx, username, password
func (_m *Authenticator) Login(ctx context.Context, username string, password string) (*auth.User, error) {
	ret := _m.Called(ctx, username, password)

	var r0 *auth.User
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *auth.User); ok {
		r0 = rf(ctx, username, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*auth.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, username, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: ctx, user
func (_m *Authenticator) Register(ctx context.Context, user *auth.User) error {
	ret := _m.Called(ctx, user)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *auth.User) error); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
