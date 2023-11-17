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

	url, err := services.GenerateExcel(payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error while generating excel",
		})
	}

	c.JSON(http.StatusOK, url)
}
