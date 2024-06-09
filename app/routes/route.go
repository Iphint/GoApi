package routes

import (
	"goapi/app/controller"
	"goapi/app/middleware"

	"github.com/labstack/echo/v4"
)

func Router(e *echo.Echo) {
	// Routes auth
	e.POST("/register", controller.Register)
	e.POST("/login", controller.Login)
	e.GET("/products", controller.ShowProductsHandler)
	e.GET("/users/:id", controller.GetUserById)

	// JWT protected routes
	productGroup := e.Group("/products")
	transactionGroup := e.Group("/transaction")
	productGroup.Use(middleware.JWTMiddleware)
	transactionGroup.Use(middleware.JWTMiddleware)

	// Routes products
	productGroup.POST("", controller.CreateProductHandler)
	productGroup.GET("/:id", controller.ShowProductHandler)
	
	productGroup.PUT("/:id", controller.UpdateProducthandler)
	productGroup.DELETE("/:id", controller.DeleteProducthandler)

	// Routes transactions
	transactionGroup.POST("", controller.CreateTransaction)
	transactionGroup.GET("", controller.ShowTransactions)
	transactionGroup.DELETE("/:id", controller.DeleteTransaction)
}
