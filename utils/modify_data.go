package utils

import (
	"log"
	"math/rand"
	"time"

	"github.com/ThanhTien96/airbnb-api/models"
	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/gorm"
)

var (
	MOVIE_LIMIT           = 100
	GENRE_LIMIT           = 8
	REVIEWER_LIMIT        = 100
	RATING_LIMIT          = 1000
	ACTOR_LIMIT           = 500
	MOVIE_CAST_LIMIT      = 10000
	MOVIE_GENRE_LIMIT     = 200
	MOVIE_DIRECTION_LIMIT = 90
	DIRECTOR_LIMIT        = 80
)

func CreateMovies(db *gorm.DB, limit int) {
	log.Println("Creating Movies ...")
	for i := 0; i < limit; i++ {
		movie := models.MovieBase{
			CreateMovieRequest: models.CreateMovieRequest{
				MovTitle:      gofakeit.Sentence(1 + rand.Intn(4)),
				MovYear:       gofakeit.Year(),
				MovTime:       gofakeit.IntRange(60, 200),
				MovLang:       gofakeit.Language(),
				MovDtRel:      gofakeit.DateRange(time.Date(2020, 1, 0, 0, 0, 0, 0, time.UTC), time.Now()).Unix(),
				MovRelCountry: gofakeit.Country(),
			},
		}
		_ = db.Create(&movie)
	}
}

func CreateGenres(db *gorm.DB) {
	log.Println("Creating Genres ...")
	var movieGenres = []string{"Action", "Comedy", "Drama", "Fantasy", "Horror", "Mystery", "Romance", "Thriller"}
	for i := 0; i < len(movieGenres); i++ {
		genres := models.GenreBase{
			CreateGenreRequest: models.CreateGenreRequest{
				GenTitle: movieGenres[i],
			},
		}
		_ = db.Create(&genres)
	}
}

func CreateReviewer(db *gorm.DB, limit int) {
	log.Println("Creating Reviewers ...")
	for i := 0; i < limit; i++ {
		reviewers := models.Reviewer{
			RevName: gofakeit.Name(),
		}
		_ = db.Create(&reviewers)
	}
}

func CreateRating(db *gorm.DB, limit int) {
	log.Println("Creating Ratings ...")
	for i := 0; i < limit; i++ {
		ratings := models.Rating{
			MovID:       1 + rand.Intn(MOVIE_LIMIT),
			RevID:       1 + rand.Intn(REVIEWER_LIMIT),
			RevStars:    1 + rand.Intn(5),
			NumORatings: 1 + rand.Intn(10),
		}
		_ = db.Create(&ratings)
	}
}

func CreateMovieGenre(db *gorm.DB, limit int) {
	log.Println("Creating Movie Genres ...")
	for i := 0; i < limit; i++ {
		movieGenres := models.MovieGenre{
			MovID: 1 + rand.Intn(MOVIE_LIMIT),
			GenID: 1 + rand.Intn(GENRE_LIMIT),
		}
		_ = db.Create(&movieGenres)
	}
}

func CreateActors(db *gorm.DB, limit int) {
	log.Println("Creating Actors ...")
	for i := 0; i < limit; i++ {
		actors := models.Actor{
			ActFName:  gofakeit.FirstName(),
			ActLName:  gofakeit.LastName(),
			ActGender: gofakeit.Gender(),
		}
		_ = db.Create(&actors)
	}
}

func CreateMovieCasts(db *gorm.DB, limit int) {
	log.Println("Creating Movie Casts ...")
	for i := 0; i < limit; i++ {
		movieCasts := models.MovieCast{
			ActID: 1 + rand.Intn(ACTOR_LIMIT),
			MovID: 1 + rand.Intn(MOVIE_LIMIT),
			Role:  gofakeit.Name(),
		}
		_ = db.Create(&movieCasts)
	}
}

func CreateDirectors(db *gorm.DB, limit int) {
	log.Println("Creating Directors ...")
	for i := 0; i < limit; i++ {
		directors := models.Director{
			DirFName: gofakeit.FirstName(),
			DirLName: gofakeit.LastName(),
		}
		_ = db.Create(&directors)
	}
}

func CreateMovieDirection(db *gorm.DB, limit int) {
	log.Println("Creating Movie Direction ...")
	for i := 0; i < limit; i++ {
		movieDirections := models.MovieDirection{
			DirID: 1 + rand.Intn(DIRECTOR_LIMIT),
			MovID: 1 + rand.Intn(MOVIE_LIMIT),
		}
		_ = db.Create(&movieDirections)
	}
}

func CleanData(db *gorm.DB) {
	var movies []models.Movie
	log.Println("Deleting Movies")
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&movies)
	log.Println("Reset sequence of movies")
	db.Exec("ALTER SEQUENCE movies_mov_id_seq RESTART WITH 1;")

	var actors []models.Actor
	log.Println("Deleting Actors")
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&actors)
	log.Println("Reset sequence of actors")
	db.Exec("ALTER SEQUENCE actors_act_id_seq RESTART WITH 1;")

	var directors []models.Director
	log.Println("Deleting Directors")
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&directors)
	log.Println("Reset sequence of directors")
	db.Exec("ALTER SEQUENCE directors_dir_id_seq RESTART WITH 1;")

	var reviewers []models.Reviewer
	log.Println("Deleting Reviewers")
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&reviewers)
	log.Println("Reset sequence of reviewers")
	db.Exec("ALTER SEQUENCE reviewers_rev_id_seq RESTART WITH 1;")

}
