package http

import (
	"fmt"

	"github.com/wagaru/Recodar/server/internal/config"
	"github.com/wagaru/Recodar/server/internal/delivery/http/router"
	"github.com/wagaru/Recodar/server/internal/usecase"
)

type httpDelivery struct {
	usecase usecase.Usecase
	router  *router.Router
	config  *config.Config
}

func NewHttpDelivery(usecase usecase.Usecase, config *config.Config) *httpDelivery {
	return &httpDelivery{
		usecase: usecase,
		router:  router.NewRouter(),
		config:  config,
	}
}

func (delivery *httpDelivery) buildRoute() {
	login := delivery.router.Group("/login")
	{
		login.GET("/google", delivery.googleLogin)
		login.GET("/google/callback", delivery.googleLoginCallback)
	}
	v1 := delivery.router.Group("/api/v1")
	{
		// auth := delivery.router.Group("/oauth2")
		// {
		// 	auth.GET("/google", delivery.oauthGoogle)
		// 	//oauth.GET("/facebook", delivery.oauthFacebook)
		// }

		v1.GET("/users", delivery.routeUsers)
	}
}

func (delivery *httpDelivery) Run(port uint16) {
	delivery.buildRoute()
	delivery.router.Run(fmt.Sprintf(":%v", port))
}
