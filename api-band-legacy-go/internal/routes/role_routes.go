package routes

import (
	"github.com/PkMs7/api-band-legacy-go/internal/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoleRoutes(router *gin.Engine, controller *controllers.RoleController) {
	group := router.Group("roles/")
	{
		group.POST("", controller.CreateRole)
		group.GET("", controller.GetRoles)
		group.GET("/:id", controller.GetRoleById)
		group.PUT("/:id", controller.UpdateRoleById)
		group.DELETE("/:id", controller.DeleteRoleById)
	}
}
