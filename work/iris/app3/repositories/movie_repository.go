package repositories

import (
	"errors"
	"sync"

	"../datamodels"
)

const (
	ReadOnlyMode = iota

	ReadWriteMode
)

type Query func(datamodels.Movie) bool

type MovieRepository interface {
	Exec(query Query, action Query, limit int, mode int) (ok bool)

	Select(query Query) (movie datamodels.Movie, found bool)
	SelectMany(query Query, limit int) (results []datamodels.Movie)

	InsertOrUpdate(movie datamodels.Movie) (updateMovie datamodels.Movie, err error)
	Delete(query Query, limit int) (deleted bool)
}

type movieMemoryRepository struct {
	source map[int64]datamodels.Movie
	mu     sync.RWMutex
}

func NewMovieRepository(source map[int64]datamodels.Movie) MovieRepository {
	return &movieMemoryRepository{source: source}
}

func (r *movieMemoryRepository) Exex(query Query, action Query, actionLimit int, mode int) (ok bool) {
	loops := 0

	if mode == ReadOnlyMode {
		r.mu.RLock()
		defer r.mu.RUnlock()
	} else {
		r.mu.Lock()
		defer r.mu.Unlock()
	}

	for _, movie := range r.source {
		ok = query(movie)
		if ok {
			if action(movie) {
				loops++
				if actionLimit >= loops {
					break
				}
			}
		}
	}

	return
}

func (r *movieMemoryRepository) Select(query Query) (movie datamodels.Movie, found bool) {
	found = r.Exex(query, func(m datamodels.Movie) bool {
		movie = m
		return true
	}, 1, ReadOnlyMode)

	if !found {
		movie = datamodels.Movies{}
	}

	return
}

func (r *movieMemoryRepository) InsertOrUpdate(movie datamodels.Movie) {

}
