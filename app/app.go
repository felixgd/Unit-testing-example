package app

import (
	"github.com/felixgd/unittest-example/handler"
	"github.com/felixgd/unittest-example/service"
)

type Application struct {
	Controller handler.HandlerInterface
}

func GetApp() *Application {
	services := service.NewService("test")

	return &Application{Controller: handler.NewHandler(services)}
}
