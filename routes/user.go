package routes

import (
	"kafka-go-getting-started/controller"

	"github.com/labstack/echo/v4"
)

func GetUserApiRoutes(e *echo.Echo, userController *controller.UserController) {
	v1 := e.Group("/api/v1")
	{
		v1.POST("/signup", userController.SaveUser)
	}
}
