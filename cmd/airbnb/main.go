package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ThanhTien96/airbnb-api/internal/config"
	"github.com/labstack/echo/v4"
)

func main() {
	config, err := config.LoadConfigFromFile("../etc/config.toml")
	if err != nil {
		log.Fatalf("load config errors: %w", err)
	}

	fmt.Printf("config: %v\n", config)
	e := echo.New()

	apiV1 := e.Group("/v1")
	apiV1.GET("/healthcheck", hello)
	e.Logger.Fatal(e.Start(":8000"))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "hello world")
}
