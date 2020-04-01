package app

import (
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func (app *Application) StartServer(engine *gin.Engine, testingMode bool) {
	if testingMode {
		servertest := httptest.NewServer(engine)
		log.Println(servertest)
		return
	}
	srv := &http.Server{
		Addr:    ":8080",
		Handler: engine,
	}

	// service connections
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}
