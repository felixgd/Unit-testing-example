package handler

import (
	"net/http"

	"github.com/felixgd/unittest-example/service"
	"github.com/gin-gonic/gin"
)

type handler struct {
	service service.ServiceInterface
}

func NewHandler(service service.ServiceInterface) HandlerInterface {
	return &handler{service}
}

func (h *handler) ReturnParam(c *gin.Context) {
	param := c.Param("name")

	resp, err := h.service.ReturnText(param)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"response": resp,
	})

	return
}
