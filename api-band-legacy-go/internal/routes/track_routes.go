package routes

import (
	"github.com/PkMs7/api-band-legacy-go/internal/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterTrackRoutes(router *gin.Engine, controller *controllers.TrackController) {
	group := router.Group("tracks")
	{
		group.POST("", controller.CreateTrack)
		group.GET("", controller.GetTracks)
		group.GET("/:id", controller.GetTrackById)
		group.PUT("/:id", controller.UpdateTrackById)
		group.DELETE("/:id", controller.DeleteTrackById)
	}
}
