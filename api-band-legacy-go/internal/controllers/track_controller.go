package controllers

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/PkMs7/api-band-legacy-go/internal/models"
	"github.com/PkMs7/api-band-legacy-go/internal/services"
	"github.com/gin-gonic/gin"
)

type TrackController struct {
	service *services.TrackService
}

func NewTrackController(service *services.TrackService) *TrackController {
	return &TrackController{service: service}
}

func (c *TrackController) CreateTrack(ctx *gin.Context) {
	var request struct {
		AlbumId         int64  `json:"album_id"`
		Title           string `json:"title"`
		TrackNumber     int64  `json:"track_number"`
		DurationSeconds int64  `json:"duration_seconds"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	track, err := c.service.CreateTrack(ctx, request.AlbumId, request.Title, request.TrackNumber, request.DurationSeconds)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusCreated, track)
}

func (c *TrackController) GetTracks(ctx *gin.Context) {
	tracks, err := c.service.GetTracks(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tracks)
}

func (c *TrackController) GetTrackById(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid track id",
		})
	}

	track, err := c.service.GetTrackById(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "track not found",
			})
			return
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to fetch track",
		})
		return
	}

	ctx.JSON(http.StatusOK, track)
}

func (c *TrackController) UpdateTrackById(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid track id",
		})
		return
	}

	var request struct {
		AlbumId         int64  `json:"album_id"`
		Title           string `json:"title"`
		TrackNumber     int64  `json:"track_number"`
		DurationSeconds int64  `json:"duration_seconds"`
	}

	if err := ctx.ShouldBindBodyWithJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	track := &models.Track{
		ID:              id,
		AlbumId:         request.AlbumId,
		Title:           request.Title,
		TrackNumber:     request.TrackNumber,
		DurationSeconds: request.DurationSeconds,
	}

	err = c.service.UpdateTrackById(ctx, track)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "track not found",
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to update track",
		})

		return
	}

	ctx.JSON(http.StatusOK, track)
}

func (c *TrackController) DeleteTrackById(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid track id",
		})
		return
	}

	err = c.service.DeleteTrackById(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "track not found",
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to delete track",
		})
		return
	}

	ctx.Status(http.StatusNoContent)
}
