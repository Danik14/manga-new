package main

import (
	"cart-microservice/cart"
	"cart-microservice/database"
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	router.POST("/api/cart/test", func(c *gin.Context) {
		fmt.Println("+++++++1234++++++")
		var testInput = struct {
			Message string `json:"message"`
		}{}

		err := c.ShouldBindJSON(&testInput)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "bad"})
			return
		}

		c.JSON(http.StatusOK, testInput)
		return
	}) // Test and be done by api gateway

	db := database.StartConnection()

	// CART
	var (
		cartRepository = cart.NewRepository(db)
		cartService    = cart.NewService(cartRepository)
		cartHandler    = cart.NewHandler(cartService)
	)
	router.GET("/api/cart/:userUUID", cartHandler.GetCartByUUIDHandler)
	router.POST("/api/cart", cartHandler.AddItemByMangaUUIDHandler)
	router.POST("/api/cart/update", cartHandler.UpdateQuantityByCartUUIDHandler)
	router.DELETE("/api/cart/item/:cartUUID", cartHandler.DeleteCartByUUIDHandler)

	router.DELETE("/api/cart/:userUUID", cartHandler.DeleteUserCartByUUIDHandler)

	router.Run(":8080")
}
