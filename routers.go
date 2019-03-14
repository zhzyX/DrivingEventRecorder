package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"os"
)

func initEcho() *echo.Echo {
	e := echo.New()
	e.HideBanner = true

	// logger middleware
	if C.Log {
		fp, err := os.OpenFile(C.LogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("Failed to open file: %v\n", C.LogPath)
			os.Exit(1)
		}
		e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: "${time_rfc3339} method=${method}, uri=${uri}, status=${status}\n",
			Output: fp,
		}))
	} else {
		e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: "${time_rfc3339} method=${method}, uri=${uri}, status=${status}\n",
		}))
	}
	// CORS middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	// static file router
	e.File("/", C.PublicPath)

	// api router
	api := e.Group("/api")
	{
		api.GET("/ping", ping)
		api.GET("/event_definition", eventDefinition)
	}

	fmt.Println("Router initialized.")
	return e
}
