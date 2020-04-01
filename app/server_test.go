package app_test

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	. "github.com/felixgd/unitTest-example/app"
)

func TestServer(t *testing.T) {
	app := Application{}

	gin.SetMode(gin.TestMode)
	res := httptest.NewRecorder()
	_, router := gin.CreateTestContext(res)

	go app.StartServer(router, false)
}
