package controllers

import (
	"errors"

	"../../datamodels"
	"../../services"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type MovieController struct {
	mvc.C

	Service services.MovieService
}

func (c *MovieController) Get() (results []datamodels.Movie) {
	return c.Service.GetAll()
}

func (c *MovieController) GetBy(id int64) (movie datamodels.Movie, found bool) {
	return c.Service.GetByID(id)
}
