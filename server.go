package main

import (
	"Movies/routes"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	// Routes
	routes.Route(e)
	e.Logger.Fatal(e.Start(":9090"))

}
