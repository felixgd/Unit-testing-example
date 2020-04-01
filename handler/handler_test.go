package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	mocks "github.com/felixgd/unitTest-example/_mocks"
	. "github.com/felixgd/unitTest-example/handler"
)

func TestHandler(t *testing.T) {
	mockService := new(mocks.ServiceInterface)

	gin.SetMode(gin.TestMode)
	res := httptest.NewRecorder()

	ctx, router := gin.CreateTestContext(res)

	test := func(method, path, param string, expected int) func(t *testing.T) {
		return func(t *testing.T) {
			response := "Hello! " + param
			mockService.On("ReturnText", param).Return(response, nil).Once()

			handler := NewHandler(mockService)

			router.GET("/test/:name", func(context *gin.Context) {
				handler.ReturnParam(ctx)
			})

			req, err := http.NewRequest(method, "/test/felix", nil)

			if err != nil {
				t.Errorf("an Error occured: %v", err)
			}

			router.ServeHTTP(res, req)

			assert.Equal(t, expected, res.Code)
			mockService.AssertExpectations(t)
		}
	}

	t.Run("Handler OK", test("GET", "/felix", "felix", http.StatusOK))

}
