package controllers

import (
	"net/http"

	"github.com/Felipedelima123/excelor/dtos"
	"github.com/Felipedelima123/excelor/services"

	"github.com/gin-gonic/gin"
)

func GenerateExcel(c *gin.Context) {
	var payload dtos.GenerateExcelDto

	c.BindJSON(&payload)

	services.GenerateExcel(payload)

	c.JSON(http.StatusOK, payload)
}
