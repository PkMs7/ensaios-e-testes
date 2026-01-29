package routes

import (
	"github.com/PkMs7/api-band-legacy-go/internal/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterBandRoutes(router *gin.Engine, controller *controllers.BandController) {
	group := router.Group("bands/")
	{
		group.POST("", controller.CreateBand)
		group.GET("", controller.GetBands)
		group.GET("/:id", controller.GetBandById)
		group.PUT("/:id", controller.UpdateBandById)
		group.DELETE("/:id", controller.DeleteBandById)
	}
}
