package usecase

import "github.com/wagaru/Recodar-backend/internal/repository"

type Usecase interface {
}

type usecase struct {
	repo repository.Repository
}

func NewUsecase(repo repository.Repository) Usecase {
	return &usecase{
		repo: repo,
	}
}
