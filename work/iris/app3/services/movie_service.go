package services

import (
	"../datamodels"
	"../repositories"
)

type MovieService interface {
	GetAll() []datamodels.Movie
	GetByID(id int64) (datamodels.Movie, bool)
	DeleteByID(id int64) bool
	updatePosterAndGenreByID(id int64, poster string, genre string) (datamodels.Movie, error)
}

func NewMovieService(repo repositories.MovieRepository) MovieService {
	return &movieService{
		repo: repo,
	}
}

type movieService struct{
	repo repositories.MovieRepository
}

func (s *movieService) GetAll() []datamodels.Movie {
	return s.repo.SelectMany(fucn(_ datamodels.Movie) bool {
		return true
	}, -1)
}

func (s *movieService) updatePosterAndGenreByID(id int64, poster string, genre string) (datamodels.Movie, error) {
	return s.repo.InsertOrUpdate(datamodels.Movie{
		ID: id,
		Poster: poster,
		Genre: genre,
	})
}

func (s *movieService) DeleteByID(id int64) bool {
	return s.repo.Delete(func(m datamodels.Movie) bool {
		return m.ID == id
	},1)
}