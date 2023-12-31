package payment

import (
	"net/http"
	"payment-microservice/helper"

	"github.com/gin-gonic/gin"
)

type handler struct {
	paymentService Service
}

func NewHandler(paymentService Service) *handler {
	return &handler{paymentService}
}

func (this *handler) CreatePaymentHandler(c *gin.Context) {
	var orderInput OrderInput
	err := c.ShouldBindJSON(&orderInput)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Create payment failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	paymentAdded, err := this.paymentService.AddPayment(orderInput)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Create payment failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// Output
	c.JSON(http.StatusOK, paymentAdded)
	return
}
