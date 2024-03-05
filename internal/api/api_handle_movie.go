package api

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/ThanhTien96/airbnb-api/internal/common"
	"github.com/ThanhTien96/airbnb-api/internal/query"
	"github.com/ThanhTien96/airbnb-api/models"
	"github.com/ThanhTien96/airbnb-api/utils"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// Get Movies example
//
//	@Summary		Get Movies
//	@Description	Get Movies
//	@Tags			Movie
//	@ID				get-movies
//	@Accept			json
//	@Produce		json
//	@Param			limit		query	int					false	"limit"
//	@Param			page		query	int					false	"page"
//	@Param			start_date	query	int64				false	"start_date"
//	@Param			end_date	query	int64				false	"end_date"
//	@Param			sorts		query	[]string			false	"sorts"
//	@Param			search		query	string			false	"search"
//	@Param			filters		query	string	false	"filters"
//	@Router			/movies [get]
func ApiGetMovies(db *gorm.DB) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		var filterParams common.FilterParams

		err := c.Bind(&filterParams)
		if err != nil {
			return c.JSON(http.StatusOK, JsonError(http.StatusInternalServerError, err.Error()))
		}
		db, paging, err := utils.ApplyFilterQuery(&filterParams, db)
		if err != nil {
			return c.JSON(http.StatusOK, JsonError(http.StatusInternalServerError, err.Error()))
		}
		movies, err, total := query.GetMovies(db)
		paging.Total = total
		if err != nil {
			return c.JSON(http.StatusOK, JsonError(http.StatusInternalServerError, err.Error()))
		}
		return c.JSON(http.StatusOK, PagingSuccessResponse("success", movies, paging))
	})
}

// Get Movie example
//
//	@Summary		Get Movie
//	@Description	Get Movie
//	@Tags			Movie
//	@ID				get-movie
//	@Accept			json
//	@Produce		json
//	@Param			mov_id	path	int	true	"mov_id"
//	@Router			/movies/{mov_id} [get]
func ApiGetMovie(db *gorm.DB) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		movId := c.Param("mov_id")
		id, err := strconv.Atoi(movId)
		if err != nil {
			return c.JSON(http.StatusOK, JsonError(http.StatusInternalServerError, err.Error()))
		}
		movies, err := query.GetMovie(db, id)
		if err != nil {
			return c.JSON(http.StatusOK, JsonError(http.StatusInternalServerError, err.Error()))
		}
		return c.JSON(http.StatusOK, DataSuccessResponse("success", movies))
	})
}

// Create Movie example
//
//	@Summary		Create Movie
//	@Description	Create Movie
//	@Tags			Movie
//	@ID				create-movie
//	@Accept			json
//	@Produce		json
//	@Param			MovieCreateRequest	body	models.CreateMovieRequest	true	"MovieCreateRequest"
//	@Router			/movies [post]
func ApiCreateMovie(db *gorm.DB) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		var movieCreateRequest models.CreateMovieRequest
		err := c.Bind(&movieCreateRequest)
		if err != nil {
			return c.JSON(http.StatusOK, JsonError(http.StatusInternalServerError, err.Error()))
		}
		fmt.Println("movie", movieCreateRequest)
		movies, err := query.CreateMovie(db, movieCreateRequest)
		if err != nil {
			return c.JSON(http.StatusOK, JsonError(http.StatusInternalServerError, err.Error()))
		}
		return c.JSON(http.StatusOK, DataSuccessResponse("success", movies))
	})
}

// Update Movie example
//
//	@Summary		Update Movie
//	@Description	Update Movie
//	@Tags			Movie
//	@ID				update-movie
//	@Accept			json
//	@Produce		json
//	@Param			mov_id				path	int							true	"mov_id"
//	@Param			MovieCreateRequest	body	models.CreateMovieRequest	true	"MovieCreateRequest"
//	@Router			/movies/{mov_id} [put]
func ApiUpdateMovie(db *gorm.DB) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		movId := c.Param("mov_id")
		id, err := strconv.Atoi(movId)
		if err != nil {
			return c.JSON(http.StatusOK, JsonError(http.StatusInternalServerError, err.Error()))
		}
		var movieCreateRequest models.CreateMovieRequest
		err = c.Bind(&movieCreateRequest)
		if err != nil {
			return c.JSON(http.StatusOK, JsonError(http.StatusInternalServerError, err.Error()))
		}
		movies, err := query.UpdateMovie(db, movieCreateRequest, id)
		if err != nil {
			return c.JSON(http.StatusOK, JsonError(http.StatusInternalServerError, err.Error()))
		}
		return c.JSON(http.StatusOK, DataSuccessResponse("success", movies))
	})
}

// Delete Movie example
//
//	@Summary		Delete Movie
//	@Description	Delete Movie
//	@Tags			Movie
//	@ID				delete-movie
//	@Accept			json
//	@Produce		json
//	@Param			mov_id	path	int	true	"mov_id"
//	@Router			/movies/{mov_id} [delete]
func ApiDeleteMovie(db *gorm.DB) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		movId := c.Param("mov_id")
		id, err := strconv.Atoi(movId)
		if err != nil {
			return c.JSON(http.StatusOK, JsonError(http.StatusInternalServerError, err.Error()))
		}
		movies, err := query.DeleteMovie(db, id)
		if err != nil {
			return c.JSON(http.StatusOK, JsonError(http.StatusInternalServerError, err.Error()))
		}
		return c.JSON(http.StatusOK, DataSuccessResponse("success", movies))
	})
}

// Delete Movies example
//
//	@Summary		Delete Movies
//	@Description	Delete Movies
//	@Tags			Movie
//	@ID				delete-movies
//	@Accept			json
//	@Produce		json
//	@Param			mov_ids	query	[]int	true	"mov_ids"
//	@Router			/movies [delete]
func ApiDeleteMovies(db *gorm.DB) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		movIds := c.QueryParam("mov_ids")
		splitId := strings.Split(movIds, ",")
		ids := make([]int, 0, len(splitId))
		for _, raw := range splitId {
			v, err := strconv.Atoi(raw)
			if err != nil {
				log.Print(err)
				continue
			}
			ids = append(ids, v)
		}
		movies, err := query.DeleteMovies(db, ids)
		if err != nil {
			return c.JSON(http.StatusOK, JsonError(http.StatusInternalServerError, err.Error()))
		}
		return c.JSON(http.StatusOK, DataSuccessResponse("success", movies))
	})
}
