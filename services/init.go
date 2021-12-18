package services

import "github.com/leyiqiang/home-library-server/interfaces"

type service struct {
	repo interfaces.Repository
}

func NewService(r interfaces.Repository) interfaces.Service {
	return &service{
		repo: r,
	}
}
