package service

import (
	"testing"
	"users/internal/models"
	"users/internal/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockUserRepo struct {
	mock.Mock
}

func (m *mockUserRepo) CreateUser(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *mockUserRepo) UpdateUser(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *mockUserRepo) DeleteUser(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *mockUserRepo) GetById(id int) (*models.User, error) {
	args := m.Called(id)
	user, ok := args.Get(0).(*models.User)
	if !ok {
		return nil, args.Error(1)
	}
	return user, args.Error(1)
}

func TestUserService_CreateUser(t *testing.T) {
	var _ repository.UserRepository = (*mockUserRepo)(nil)
	mockRepo := new(mockUserRepo)
	svc := NewUserService(mockRepo)

	user := &models.User{Name: "Alice", Email: "newemail@gmail.com"}

	mockRepo.On("CreateUser", user).Return(nil)

	err := svc.CreateUser(user)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUserService_GetById(t *testing.T) {
	var _ repository.UserRepository = (*mockUserRepo)(nil)
	mockRepo := new(mockUserRepo)
	svc := NewUserService(mockRepo)

	expectedUser := &models.User{ID: 1, Name: "Bob", Email: "asdsadasdas@gmail.com"}

	mockRepo.On("GetById", 1).Return(expectedUser, nil)

	user, err := svc.GetById(1)

	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)
	mockRepo.AssertExpectations(t)
}

func TestUserService_UpdateUser(t *testing.T) {
	var _ repository.UserRepository = (*mockUserRepo)(nil)
	mockRepo := new(mockUserRepo)
	svc := NewUserService(mockRepo)

	user := &models.User{Name: "newName", Email: "newemail2345@gmail.com"}

	mockRepo.On("UpdateUser", user).Return(nil)

	err := svc.UpdateUser(user)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUserService_DeleteUser(t *testing.T) {
	var _ repository.UserRepository = (*mockUserRepo)(nil)
	mockRepo := new(mockUserRepo)
	svc := NewUserService(mockRepo)

	mockRepo.On("DeleteUser", 1).Return(nil)

	err := svc.DeleteUser(1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
