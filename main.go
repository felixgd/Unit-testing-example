package main

import (
	"github.com/felixgd/unittest-example/app"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	app := app.GetApp()
	app.MapRoutes(engine)
	app.StartServer(engine, false)
}
