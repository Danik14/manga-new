package order

import (
	"fmt"
	"net/http"
	"os"
	"payment-microservice/helper"

	"github.com/gin-gonic/gin"
)

type handler struct {
	orderService Service
}

func NewHandler(orderService Service) *handler {
	return &handler{orderService}
}

func (this *handler) CreateOrderHandler(c *gin.Context) {
	uuid := c.Params.ByName("userUUID")
	var cartInput []CartInput
	err := c.ShouldBindJSON(&cartInput)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Create order failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	orderAdded, err := this.orderService.AddOrderByUUID(uuid)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Create order failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	orderDetailsAdded, err := this.orderService.AddOrderDetailsByOrderUUID(cartInput, orderAdded.ID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Create order details failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	client := &http.Client{}

	httpRequest, err := http.NewRequest("DELETE", "http://"+os.Getenv("CART_MICROSERVICE_URL")+fmt.Sprintf("/api/cart/%s", uuid), nil)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Delete cart failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	httpRequest.Header.Set("Content-Type", "application/json")

	_, err = client.Do(httpRequest)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Delete cart failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	c.JSON(http.StatusOK, orderDetailsAdded)
	return
}
