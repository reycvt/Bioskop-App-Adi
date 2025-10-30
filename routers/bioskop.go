package routers

import (
	"bioskop-app-adi/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterBioskopRoutes(router *gin.Engine) {
	bioskop := router.Group("/bioskop")

	bioskop.POST("/", controllers.CreateBioskop)
	bioskop.GET("/", controllers.GetAllBioskop)
}
