package routes

import (
	"go_shop/controllers"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	products := e.Group("/products")
	products.GET("", controllers.ListProducts)
	products.POST("", controllers.CreateProduct)
	products.GET("/:id", controllers.GetProduct)
	products.PUT("/:id", controllers.UpdateProduct)
	products.DELETE("/:id", controllers.RemoveProduct)

	categories := e.Group("/categories")
	categories.GET("", controllers.ListCategories)
	categories.POST("", controllers.CreateProduct)
	categories.GET("/:id", controllers.GetProduct)
	categories.PUT("/:id", controllers.UpdateProduct)
	categories.DELETE("/:id", controllers.RemoveProduct)
	categories.GET("/:id/products", controllers.GetProductsByCategoryId)

	carts := e.Group("/carts")
	carts.GET("", controllers.ListCarts)
	carts.POST("", controllers.CreateCart)
	carts.GET("/:id", controllers.GetCart)
	carts.PUT("/:id", controllers.UpdateCart)
	carts.DELETE("/:id", controllers.RemoveCart)
	carts.POST("/:id/products/:itemId", controllers.AddProductToCart)
	carts.PUT("/:id/products/:itemId", controllers.UpdateCartItem)
	carts.DELETE("/:id/products/:itemId", controllers.RemoveCartItem)

	e.POST("/checkout", controllers.HandleCheckout)
}
