package repositories

import (
	"/datamodels"
	"errors"
	"sync"
)

type Query func(datamodels.Movie) bool

type MovieRepository interface {
	Exec(query Query, action Query, limit int, mode int) (ok bool)

	Select(query Query) (movie datamodels.Movie, found bool)
	SelectMany(query Query, limit int) (results []datamodels.Movie)

	InsertOrUpdate(movie datamodels.Movie) (updateMovie datamodels.Movie, err error)
	Delete(query Query, limit int) (deleted bool)
}
