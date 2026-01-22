package routes

import (
	"github.com/PkMs7/api-band-legacy-go/internal/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterArtistRoutes(router *gin.Engine, controller *controllers.ArtistController) {
	group := router.Group("artists/")
	{
		group.POST("", controller.CreateArtist)
		group.GET("", controller.GetArtists)
		group.GET("/:id", controller.GetArtistsById)
		group.PUT("/:id", controller.UpdateArtistById)
		group.DELETE("/:id", controller.DeleteArtistById)
	}
}
