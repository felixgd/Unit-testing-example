package app

import (
	"github.com/gin-gonic/gin"
)

func (app *Application) MapRoutes(r *gin.Engine) {
	r.GET("/:name", app.Controller.ReturnParam)
}
