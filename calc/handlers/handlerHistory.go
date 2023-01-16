package handlers

import (
	"net/http"

	"main.go/services"

	"github.com/gin-gonic/gin"
)

func History(context *gin.Context) {
	context.JSON(http.StatusOK, services.ShowHistory())
}
