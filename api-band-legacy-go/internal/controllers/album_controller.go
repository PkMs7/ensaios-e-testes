package controllers

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/PkMs7/api-band-legacy-go/internal/models"
	"github.com/PkMs7/api-band-legacy-go/internal/services"
	"github.com/gin-gonic/gin"
)

type AlbumController struct {
	service *services.AlbumService
}

func NewAlbumController(service *services.AlbumService) *AlbumController {
	return &AlbumController{service: service}
}

func (c *AlbumController) CreateAlbum(ctx *gin.Context) {
	var request struct {
		BandId      int64  `json:"band_id"`
		Title       string `json:"title"`
		ReleaseDate string `json:"release_date"`
		AlbumType   string `json:"album_type"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	releaseDate, err := time.Parse("2006-01-02", request.ReleaseDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "release_date must be in YYYY-MM-DD format",
		})
		return
	}

	album, err := c.service.CreateAlbum(ctx, request.BandId, request.Title, releaseDate, request.AlbumType)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusCreated, album)

}

func (c *AlbumController) GetAlbums(ctx *gin.Context) {
	albums, err := c.service.GetAlbums(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, albums)
}

func (c *AlbumController) GetAlbumsById(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid album id",
		})
	}

	album, err := c.service.GetAlbumsById(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "album not found",
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch album",
		})
		return
	}

	ctx.JSON(http.StatusOK, album)
}

func (c *AlbumController) UpdateAlbumById(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid album id",
		})
		return
	}

	var request struct {
		BandId      int64  `json:"band_id"`
		Title       string `json:"title"`
		ReleaseDate string `json:"release_date"`
		AlbumType   string `json:"album_type"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	releaseDate, err := time.Parse("2006-01-02", request.ReleaseDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "release_date must be in YYYY-MM-DD format",
		})
		return
	}

	album := &models.Album{
		ID:          id,
		BandId:      request.BandId,
		Title:       request.Title,
		ReleaseDate: releaseDate,
		AlbumType:   request.AlbumType,
	}

	err = c.service.UpdateAlbumById(ctx, album)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "album not found",
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to update album",
		})
		return
	}

	ctx.JSON(http.StatusOK, album)

}

func (c *AlbumController) DeleteAlbumById(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid album id",
		})
		return
	}

	err = c.service.DeleteAlbumById(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "album not found",
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to delete album",
		})
		return
	}

	ctx.Status(http.StatusNoContent)

}
