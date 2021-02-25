package http

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/wagaru/Recodar-backend/internal/usecase"
)

type httpDelivery struct {
	usecase usecase.Usecase
	router  *gin.Engine
}

func NewHttpDelivery(usecase usecase.Usecase) *httpDelivery {
	return &httpDelivery{
		usecase: usecase,
		router:  gin.Default(),
	}
}

func (delivery *httpDelivery) buildRoute() {
	delivery.router.GET("/users", delivery.routeUsers)
}

func (delivery *httpDelivery) Run(port uint16) {
	delivery.buildRoute()
	delivery.router.Run(fmt.Sprintf(":%v", port))
}
