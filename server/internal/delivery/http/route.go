package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (delivery *httpDelivery) routeUsers(c *gin.Context) {
	c.String(http.StatusOK, "hello")
}
