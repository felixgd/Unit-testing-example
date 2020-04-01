package handler

import (
	"github.com/gin-gonic/gin"
)

type HandlerInterface interface {
	ReturnParam(c *gin.Context)
}
