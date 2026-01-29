package routes

import (
	"github.com/PkMs7/api-band-legacy-go/internal/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterAlbumRoutes(router *gin.Engine, controller *controllers.AlbumController) {
	group := router.Group("albums/")
	{
		group.POST("", controller.CreateAlbum)
		group.GET("", controller.GetAlbums)
		group.GET("/:id", controller.GetAlbumsById)
		group.PUT("/:id", controller.UpdateAlbumById)
		group.DELETE("/:id", controller.DeleteAlbumById)
	}
}
