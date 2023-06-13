package main

import (
	"net/http"
	"payment-microservice/database"
	"payment-microservice/order"
	"payment-microservice/payment"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	router.POST("/api/payment/test", func(c *gin.Context) {
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
	})

	db := database.StartConnection()

	var (
		orderRepository = order.NewRepository(db)
		orderService    = order.NewService(orderRepository)
		orderHandler    = order.NewHandler(orderService)
	)
	router.POST("/api/order/cart/:userUUID", orderHandler.CreateOrderHandler)

	var (
		paymentRepository = payment.NewRepository(db)
		paymentService    = payment.NewService(paymentRepository, orderRepository)
		paymentHandler    = payment.NewHandler(paymentService)
	)
	router.POST("/api/order/pay", paymentHandler.CreatePaymentHandler)

	router.Run(":8080")
}
