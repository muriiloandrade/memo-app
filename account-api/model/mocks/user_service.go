package mocks

import (
	"context"

	"github.com/google/uuid"
	"github.com/muriiloandrade/memo-app/model"
	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) Get(ctx context.Context, userId uuid.UUID) (*model.User, error) {
	ret := m.Called(ctx, userId)

	var ret0 *model.User
	if ret.Get(0) != nil {
		ret0 = ret.Get(0).(*model.User)
	}

	var ret1 error
	if ret.Get(1) != nil {
		ret1 = ret.Get(1).(error)
	}

	return ret0, ret1
}
