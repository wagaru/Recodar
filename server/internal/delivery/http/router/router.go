package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
}

type Session interface {
	sessions.Session
}

func NewRouter() *Router {
	router := &Router{gin.Default()}
	newSession(router)
	return router
}

func newSession(router *Router) {
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))
}
