package query

import (
	"github.com/ThanhTien96/airbnb-api/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetMovies(db *gorm.DB) ([]*models.MovieBase, error, int64) {
	var movies []*models.MovieBase
	result := db.Table(models.MovieTableName).Find(&movies)
	if result.Error != nil {
		return nil, result.Error, 0
	}
	var total int64
	db.Table(models.MovieTableName).Count(&total)
	return movies, nil, total
}

func GetMovie(db *gorm.DB, movId int) (*models.Movie, error) {
	var movies *models.Movie
	result := db.Table(
		models.MovieTableName,
	).Preload(
		"Rating.Reviewer",
	).Preload(
		"MovieGenre.Genre",
	).Preload(
		"MovieCast.Actor",
	).Preload(
		"MovieDirection.Director",
	).Find(&movies, movId)
	if result.Error != nil {
		return nil, result.Error
	}
	return movies, nil
}

func CreateMovie(db *gorm.DB, createMovieRequest models.CreateMovieRequest) (*models.MovieBase, error) {
	movie := models.MovieBase{
		CreateMovieRequest: createMovieRequest,
	}
	result := db.Create(&movie)
	if result.Error != nil {
		return nil, result.Error
	}
	return &movie, nil
}

func UpdateMovie(db *gorm.DB, createMovieRequest models.CreateMovieRequest, id int) (*models.MovieBase, error) {
	movie := models.MovieBase{
		MovID:              id,
		CreateMovieRequest: createMovieRequest,
	}
	result := db.Omit("CreatedAt").Updates(&movie)
	if result.Error != nil {
		return nil, result.Error
	}
	return &movie, nil
}

func DeleteMovie(db *gorm.DB, id int) (*models.MovieBase, error) {
	var movie *models.MovieBase
	result := db.Clauses(clause.Returning{}).Delete(&movie, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return movie, nil
}

func DeleteMovies(db *gorm.DB, ids []int) ([]*models.MovieBase, error) {
	var movie []*models.MovieBase
	for _, id := range ids {
		movie = append(movie, &models.MovieBase{MovID: id})
	}
	result := db.Clauses(clause.Returning{}).Delete(&movie)
	if result.Error != nil {
		return nil, result.Error
	}
	return movie, nil
}
