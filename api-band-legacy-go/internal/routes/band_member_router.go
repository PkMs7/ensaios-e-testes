package routes

import (
	"github.com/PkMs7/api-band-legacy-go/internal/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterBandMemberRoutes(router *gin.Engine, controller *controllers.BandMemberController) {
	group := router.Group("band-members/")
	{
		group.POST("", controller.CrateBandMember)
		group.GET("", controller.GetBandMembers)
		group.GET("/:id", controller.GetBandMemberById)
		group.PUT("/:id", controller.UpdateBandMemberById)
		group.DELETE("/:id", controller.DeleteBandMemberById)
	}
}
