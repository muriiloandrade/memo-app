package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/muriiloandrade/memo-app/model"
	"github.com/muriiloandrade/memo-app/model/errors"
	"github.com/muriiloandrade/memo-app/model/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestMe(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {
		userId, _ := uuid.NewRandom()

		mockUserResp := &model.User{
			UserId: userId,
			Email:  "test@test.com",
			Name:   "Test Testingson",
		}

		mockUserService := new(mocks.MockUserService)
		mockUserService.On("Get", mock.AnythingOfType("*gin.Context"), userId).Return(mockUserResp, nil)

		// Response recorder
		rr := httptest.NewRecorder()

		router := gin.Default()
		router.Use(func(ctx *gin.Context) {
			ctx.Set("user", &model.User{
				UserId: userId,
			})
		})

		NewHandler(&Config{
			R:           router,
			UserService: mockUserService,
		})

		req, err := http.NewRequest(http.MethodGet, "/me", nil)
		assert.NoError(t, err)

		router.ServeHTTP(rr, req)

		respBody, err := json.Marshal(gin.H{
			"user": mockUserResp,
		})
		assert.NoError(t, err)

		assert.Equal(t, 200, rr.Code)
		assert.Equal(t, respBody, rr.Body.Bytes())
		// assert that UserService.Get was called
		mockUserService.AssertExpectations(t)
	})

	t.Run("NoContextUser", func(t *testing.T) {
		mockUserService := new(mocks.MockUserService)
		mockUserService.On("Get", mock.Anything, mock.Anything).Return(nil, nil)

		// Response recorder
		rr := httptest.NewRecorder()

		router := gin.Default()
		NewHandler(&Config{
			R:           router,
			UserService: mockUserService,
		})

		req, err := http.NewRequest(http.MethodGet, "/me", nil)
		assert.NoError(t, err)

		router.ServeHTTP(rr, req)

		assert.Equal(t, 500, rr.Code)
		mockUserService.AssertNotCalled(t, "Get", mock.Anything)
	})

	t.Run("Not found", func(t *testing.T) {
		userId, _ := uuid.NewRandom()
		mockUserService := new(mocks.MockUserService)
		mockUserService.On("Get", mock.Anything, userId).Return(nil, fmt.Errorf("Some error down call chain"))

		// Response recorder
		rr := httptest.NewRecorder()

		router := gin.Default()
		router.Use(func(ctx *gin.Context) {
			ctx.Set("user", &model.User{
				UserId: userId,
			})
		})

		NewHandler(&Config{
			R:           router,
			UserService: mockUserService,
		})

		req, err := http.NewRequest(http.MethodGet, "/me", nil)
		assert.NoError(t, err)

		router.ServeHTTP(rr, req)

		respError := errors.NotFoundError("user", userId.String())

		respBody, err := json.Marshal(gin.H{
			"error": respError,
		})
		assert.NoError(t, err)

		assert.Equal(t, respError.Status(), rr.Code)
		assert.Equal(t, respBody, rr.Body.Bytes())
		// assert that UserService.Get was called
		mockUserService.AssertExpectations(t)
	})
}
