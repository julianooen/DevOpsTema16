package handlers

import (
	"main.go/services"
	"main.go/validator"

	"net/http"

	"github.com/gin-gonic/gin"
)

func CalcHandler(context *gin.Context) {
	operation := context.Param("operation")

	num1, err := validator.ValidadeNum(context.Param("num1"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Error": "invalid number 1"})
		return
	}

	num2, err := validator.ValidadeNum(context.Param("num2"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Error": "invalid number 2"})
		return
	}

	response, errMsg := services.OperationFunc(operation, num1, num2)
	if errMsg != "" {
		context.JSON(http.StatusBadRequest, gin.H{"Error": errMsg})
		return
	}
	context.JSON(http.StatusOK, response)
}
