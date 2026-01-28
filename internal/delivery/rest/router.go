package rest

import (
	"github.com/labstack/echo/v4"
)

func LoadRoutes(e *echo.Echo, handler *handler) {
	authMiddleware := GetAuthMiddleware(handler.restoUsecase)

	menuGroup := e.Group("/menu")
	menuGroup.GET("", handler.GetMenuList)

	orderGroup := e.Group("/order")
	orderGroup.POST("", handler.Order,
		authMiddleware.CheckAuth,
	)
	orderGroup.GET("/:orderID", handler.GetOrderInfo,
		authMiddleware.CheckAuth,
	)

	userGroup := e.Group("/user")
	userGroup.POST("/register", handler.RegisterUser)
	userGroup.POST("/login", handler.Login)
}
