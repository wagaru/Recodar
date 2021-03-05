package http

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (delivery *httpDelivery) routeUsers(c *gin.Context) {
	log.Print("abc")
	c.String(http.StatusOK, "hello")
}
