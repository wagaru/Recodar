package usecase

import (
	"github.com/wagaru/Recodar/server/internal/config"
	"github.com/wagaru/Recodar/server/internal/delivery/http/router"
	"github.com/wagaru/Recodar/server/internal/repository"
)

type Usecase interface {
	// Login
	HandleUserLogin(session router.Session, info []byte, source string) error
	GetGoogleOauthURL(session router.Session) string
	GetGoogleOauthAccessToken(state, code string, session router.Session) (string, error)
}

type usecase struct {
	repo   repository.Repository
	config *config.Config
}

func NewUsecase(repo repository.Repository, config *config.Config) Usecase {
	return &usecase{
		repo:   repo,
		config: config,
	}
}
