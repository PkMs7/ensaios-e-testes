package controllers

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/PkMs7/api-band-legacy-go/internal/models"
	"github.com/PkMs7/api-band-legacy-go/internal/services"
	"github.com/gin-gonic/gin"
)

type ArtistController struct {
	service *services.ArtistService
}

func NewArtistController(service *services.ArtistService) *ArtistController {
	return &ArtistController{service: service}
}

func (c *ArtistController) CreateArtist(ctx *gin.Context) {
	var request struct {
		Name      string `json:"name"`
		BirthDate string `json:"birth_date"`
		Biography string `json:"biography"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	birthDate, err := time.Parse("2006-01-02", request.BirthDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "birth_date must be in YYYY-MM-DD format",
		})
		return
	}

	artist, err := c.service.CreateArtist(ctx, request.Name, birthDate, request.Biography)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusCreated, artist)
}

func (c *ArtistController) GetArtists(ctx *gin.Context) {
	artists, err := c.service.GetArtists(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, artists)
}

func (c *ArtistController) GetArtistsById(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid artist id",
		})
		return
	}

	artist, err := c.service.GetArtistsById(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "artist not found",
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch artist",
		})
		return
	}

	ctx.JSON(http.StatusOK, artist)
}

func (c *ArtistController) UpdateArtistById(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid artist id",
		})
		return
	}

	var request struct {
		Name      string `json:"name"`
		BirthDate string `json:"birth_date"`
		Biography string `json:"biography"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	birthDate, err := time.Parse("2006-01-02", request.BirthDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "birth_date must be in YYYY-MM-DD format",
		})
		return
	}

	artist := &models.Artist{
		ID:        id,
		Name:      request.Name,
		BirthDate: birthDate,
		Biography: request.Biography,
	}

	fmt.Println(artist)

	err = c.service.UpdateArtistById(ctx, artist)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "artist not found",
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to update artist",
		})
		return
	}

	ctx.JSON(http.StatusOK, artist)
}

func (c *ArtistController) DeleteArtistById(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid artist id",
		})
		return
	}

	err = c.service.DeleteArtistById(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "artist not found",
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to delete artist",
		})
		return
	}

	ctx.Status(http.StatusNoContent)
}
