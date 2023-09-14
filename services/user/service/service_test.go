package service

import (
	"gokit-basic/services/user/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockedUserRepository struct {
	mock.Mock
}

func (m *MockedUserRepository) CreateUser(v *domain.UserDomain) (*domain.UserDomain, error) {
	args := m.Called(v)

	return args.Get(0).(*domain.UserDomain), args.Error(1)
}

func (m *MockedUserRepository) GetListUser() []*domain.UserDomain {

	args := m.Called()

	return args.Get(0).([]*domain.UserDomain)
}

func TestCreateUser(t *testing.T) {

	userRepoMock := new(MockedUserRepository)

	now := time.Now()

	v := &domain.UserDomain{
		ID:        domain.RandomUUID(),
		Name:      domain.NameGenerator(),
		Age:       "22",
		Phone:     "08911",
		CreatedAt: now,
		UpdatedAt: now,
	}

	ser := NewUserService(userRepoMock)

	userRepoMock.On("CreateUser", v).Return(v, nil).Once()

	res, err := ser.CreateUser(v)
	userRepoMock.AssertExpectations(t)

	assert.NotNil(t, res, "CreateUser")
	assert.NoError(t, err, "Create user not error")
}

func TestGetAll(t *testing.T) {

	userRepoMock := new(MockedUserRepository)

	var listUser []*domain.UserDomain

	listUser = append(listUser, &domain.UserDomain{
		ID:    "213",
		Name:  "andrew",
		Age:   "21",
		Phone: "99",
	})

	userRepoMock.On("GetListUser").Return(listUser).Once()

	serv := NewUserService(userRepoMock)

	res, err := serv.GetListUser()
	userRepoMock.AssertExpectations(t)

	assert.NotEmpty(t, res)
	assert.NoError(t, err)
}

func TestGetAllFailed(t *testing.T) {

	userRepoMock := new(MockedUserRepository)

	var listUser []*domain.UserDomain

	userRepoMock.On("GetListUser").Return(listUser).Once()

	serv := NewUserService(userRepoMock)

	res, err := serv.GetListUser()
	userRepoMock.AssertExpectations(t)

	assert.Empty(t, res)
	assert.Error(t, err)
}
