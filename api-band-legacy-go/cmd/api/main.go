package main

import (
	"log"

	"github.com/PkMs7/api-band-legacy-go/internal/config"
	"github.com/PkMs7/api-band-legacy-go/internal/controllers"
	"github.com/PkMs7/api-band-legacy-go/internal/repositories"
	"github.com/PkMs7/api-band-legacy-go/internal/routes"
	"github.com/PkMs7/api-band-legacy-go/internal/services"
	"github.com/gin-gonic/gin"
)

func main() {

	_ = config.LoadEnv()
	dbConfig := config.DefaultPostgresConfig()
	db, _ := config.NewPostgresConnection(dbConfig)

	repo := repositories.NewRoleRepository(db)
	service := services.NewRoleService(repo)
	controller := controllers.NewRoleController(service)

	router := gin.Default()
	routes.RegisterRoleRoutes(router, controller)

	log.Println("🚀 API running on http://localhost:8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}

}
