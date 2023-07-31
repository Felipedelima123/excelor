package routes

import (
	"github.com/Felipedelima123/excelor/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()

	r.POST("/generate-excel", controllers.GenerateExcel)

	r.Run()
}
