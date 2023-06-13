package main

import (
	"api-gateway/auth"
	"api-gateway/cart"
	"api-gateway/database"
	"api-gateway/manga"
	"api-gateway/order"
	"api-gateway/payment"
	"api-gateway/user"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	customerDb := database.StartConnection()

	// USER
	var (
		userRepository = user.NewRepository(customerDb)
		userService    = user.NewService(userRepository)
		authService    = auth.NewService()
		userHandler    = user.NewHandler(userService, authService)
	)
	router.POST("/api/user/register", userHandler.RegisterUserHandler)
	router.POST("/api/user/login", userHandler.LoginHandler)

	// CART
	var (
		cartHandler = cart.NewHandler()
	)
	router.GET("/api/cart/:userUUID", userHandler.AuthenticateHandler, cartHandler.GetCartByUUIDHandler)
	router.POST("/api/cart", userHandler.AuthenticateHandler, cartHandler.AddItemByProductUUIDHandler)
	router.POST("/api/cart/update", userHandler.AuthenticateHandler, cartHandler.UpdateQuantityByCartUUIDHandler)
	router.DELETE("/api/cart/item/:cartUUID", userHandler.AuthenticateHandler, cartHandler.DeleteCartByUUIDHandler)

	// PRODUCT
	var (
		mangaHandler = manga.NewHandler()
	)
	router.GET("/api/manga", mangaHandler.GetAllMangasHandler)
	router.GET("/api/manga/title/:title", mangaHandler.GetMangasByTitleHandler)
	router.GET("/api/manga/author/:author", mangaHandler.GetMangasByAuthorHandler)

	// ORDER
	var (
		orderHandler = order.NewHandler()
	)
	router.POST("/api/order/cart/:userUUID", userHandler.AuthenticateHandler, orderHandler.CreateOrderHandler)

	// PAYMENT
	var (
		paymentHandler = payment.NewHandler()
	)
	router.POST("/api/order/pay", userHandler.AuthenticateHandler, paymentHandler.CreatePaymentHandler)

	router.GET("/api/test", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "api gateway ok"}) })
	router.GET("/api/cart/test", cartHandler.Test)
	router.GET("/api/manga/test", mangaHandler.Test)
	router.GET("/api/payment/test", paymentHandler.Test)

	router.Run(":8080")
}
