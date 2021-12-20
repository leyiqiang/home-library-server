package controllers

import (
	"github.com/leyiqiang/home-library-server/repositories"
)

type Controller struct {
	Repo repositories.Repository
}
