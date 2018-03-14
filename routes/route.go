package routes

import (
	"Movies/controller"

	"github.com/labstack/echo"
)

// Route multi route
func Route(e *echo.Echo) (err error) {
	e.GET("/get/", controller.GetData)
	e.GET("/get/name/", controller.GetDataTitle)
	e.GET("/get/range/", controller.GetDatainRange)
	e.PUT("/update", controller.UpdateData)
	return err
}
