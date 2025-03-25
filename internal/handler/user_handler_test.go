package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"users/internal/models"
	"users/internal/service"
)

type mockUserService struct {
	mock.Mock
}

func (m *mockUserService) CreateUser(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}
func (m *mockUserService) GetById(id int) (*models.User, error) {
	args := m.Called(id)
	user, ok := args.Get(0).(*models.User)
	if !ok {
		return nil, args.Error(1)
	}
	return user, args.Error(1)
}
func (m *mockUserService) UpdateUser(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *mockUserService) DeleteUser(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestHandler_CreateUser(t *testing.T) {
	var _ service.UserService = (*mockUserService)(nil)
	var response models.User
	mockService := new(mockUserService)
	handler := NewHandler(mockService)

	body := `{"name": "Alice", "email":"newemail@gmail.com"}`
	req, _ := http.NewRequest(http.MethodPost, "/users", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req
	expectedUser := &models.User{Name: "Alice", Email: "newemail@gmail.com"}
	mockService.On("CreateUser", expectedUser).Return(nil)
	handler.CreateUser(ctx)
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, expectedUser.Name, response.Name)
	assert.Equal(t, expectedUser.Email, response.Email)
	mockService.AssertExpectations(t)
}

func TestHandler_GetById(t *testing.T) {
	var _ service.UserService = (*mockUserService)(nil)
	var response models.User
	mockService := new(mockUserService)
	handler := NewHandler(mockService)

	body := `{"id": 1,name": "Alice", "email":"newemail@gmail.com"}`
	req, _ := http.NewRequest(http.MethodGet, "/users/:id", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req
	ctx.Params = gin.Params{{Key: "id", Value: "1"}}
	expectedUser := &models.User{ID: 1, Name: "Alice", Email: "newemail@gmail.com"}
	mockService.On("GetById", 1).Return(expectedUser, nil)
	handler.GetById(ctx)
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, expectedUser.Name, response.Name)
	assert.Equal(t, expectedUser.Email, response.Email)
	mockService.AssertExpectations(t)
}
func TestHandler_UpdateUser(t *testing.T) {
	var _ service.UserService = (*mockUserService)(nil)
	var response models.User
	mockService := new(mockUserService)
	handler := NewHandler(mockService)

	body := `{"id": 1, "name": "Alice", "email":"newemail@gmail.com"}`
	req, _ := http.NewRequest(http.MethodPut, "/users/:id", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req
	ctx.Params = gin.Params{{Key: "id", Value: "1"}}
	expectedUser := &models.User{ID: 1, Name: "Alice", Email: "newemail@gmail.com"}
	fmt.Printf(">>> actual user: %+v\n", expectedUser)

	mockService.On("UpdateUser", expectedUser).Return(nil)
	handler.UpdateUser(ctx)
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, expectedUser.Name, response.Name)
	assert.Equal(t, expectedUser.Email, response.Email)
	mockService.AssertExpectations(t)
}

func TestHandler_DeleteUser(t *testing.T) {
	var _ service.UserService = (*mockUserService)(nil)
	//var response models.User
	mockService := new(mockUserService)
	handler := NewHandler(mockService)

	body := `{"id": 1, "name": "Alice", "email":"newemail@gmail.com"}`
	req, _ := http.NewRequest(http.MethodPost, "/users/:id", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req
	ctx.Params = gin.Params{{Key: "id", Value: "1"}}
	//expectedUser := &models.User{ID: 1, Name: "Alice", Email: "newemail@gmail.com"}
	mockService.On("DeleteUser", 1).Return(nil)
	handler.DeleteUser(ctx)
	assert.Empty(t, w.Body.String())
	assert.Equal(t, http.StatusOK, w.Code)
	//assert.Equal(t, expectedUser.Name, response.Name)
	//assert.Equal(t, expectedUser.Email, response.Email)
	mockService.AssertExpectations(t)
}
