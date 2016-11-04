package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/fasthttp"
	"github.com/labstack/echo/middleware"
	"github.com/rafaeljesus/kyp-secret-key-gen/handlers"
	"log"
	"os"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.Use(middleware.Secure())
	e.Use(middleware.Gzip())

	v1 := e.Group("/v1")
	v1.GET("/healthz", handlers.HealthzIndex)
	v1.POST("/secret-key", handlers.SecretKeyCreate)

	log.Print("Starting Kyp Secret Key Gen Service...")

	e.Run(fasthttp.New(":" + os.Getenv("KYP_SECRET_KEY_GEN_PORT")))
}
