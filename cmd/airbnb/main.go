package main

import (
	"flag"
	"log"
	"net/http"

	api "github.com/ThanhTien96/airbnb-api/internal/api"
	"github.com/ThanhTien96/airbnb-api/utils"

	// "github.com/ThanhTien96/airbnb-api/internal/config"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// config, err := config.LoadConfigFromFile("./config.toml")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	dsn := "host=localhost user=postgres password=123456 dbname=movie port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	isCleanData := flag.Bool("c", false, "Clean data")
	isGenerateData := flag.Bool("gd", false, "Generate data")
	isMigrate := flag.Bool("m", false, "Migrate model from code to database")
	flag.Parse()

	if *isCleanData {
		utils.CleanData(db)
	}
	if *isMigrate {
		err := utils.MigrateData(db)
		if err != nil {
			log.Fatal("Error while migrating data from model to database ", err)
		}
		// os.Exit(0)
		log.Println("Successfully migrated data")
	}
	if *isGenerateData {
		utils.CreateMovies(db, utils.MOVIE_LIMIT)
		utils.CreateGenres(db)
		utils.CreateReviewer(db, utils.REVIEWER_LIMIT)
		utils.CreateDirectors(db, utils.DIRECTOR_LIMIT)
		utils.CreateActors(db, utils.ACTOR_LIMIT)
		utils.CreateMovieCasts(db, utils.MOVIE_CAST_LIMIT)
		utils.CreateMovieGenre(db, utils.MOVIE_GENRE_LIMIT)
		utils.CreateRating(db, utils.RATING_LIMIT)
		utils.CreateMovieDirection(db, utils.MOVIE_DIRECTION_LIMIT)
	}

	e := echo.New()

	apiV1 := e.Group("/v1")
	// Movie APIs
	apiV1.GET("/movies", api.ApiGetMovies(db))
	apiV1.GET("/movies/:mov_id", api.ApiGetMovie(db))
	apiV1.POST("/movies", api.ApiCreateMovie(db))
	apiV1.PUT("/movies/:mov_id", api.ApiUpdateMovie(db))
	apiV1.DELETE("/movies/:mov_id", api.ApiDeleteMovie(db))
	apiV1.DELETE("/movies", api.ApiDeleteMovies(db))
	// Swagger endpoint

	// Swagger endpoint
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":8000"))

}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "hello world")
}
