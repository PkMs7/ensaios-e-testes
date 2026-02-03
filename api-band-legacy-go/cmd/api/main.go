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

	config.LoadEnv()
	dbConfig := config.DefaultPostgresConfig()
	db, _ := config.NewPostgresConnection(dbConfig)

	repositoryRole := repositories.NewRoleRepository(db)
	serviceRole := services.NewRoleService(repositoryRole)
	controllerRole := controllers.NewRoleController(serviceRole)

	repositoryArtist := repositories.NewArtistRepository(db)
	serviceArtist := services.NewArtistService(repositoryArtist)
	controllerArtist := controllers.NewArtistController(serviceArtist)

	repositoryBand := repositories.NewBandRepository(db)
	serviceBand := services.NewBandService(repositoryBand)
	controllerBand := controllers.NewBandController(serviceBand)

	repositoryAlbum := repositories.NewAlbumRepository(db)
	serviceAlbum := services.NewAlbumService(repositoryAlbum)
	controllerAlbum := controllers.NewAlbumController(serviceAlbum)

	repositoryTrack := repositories.NewTrackRepository(db)
	serviceTrack := services.NewTrackService(repositoryTrack)
	controllerTrack := controllers.NewTrackController(serviceTrack)

	router := gin.Default()
	routes.RegisterRoleRoutes(router, controllerRole)
	routes.RegisterArtistRoutes(router, controllerArtist)
	routes.RegisterBandRoutes(router, controllerBand)
	routes.RegisterAlbumRoutes(router, controllerAlbum)
	routes.RegisterTrackRoutes(router, controllerTrack)

	log.Println("🚀 API running on http://localhost:8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}

}
