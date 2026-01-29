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

type BandController struct {
	service *services.BandService
}

func NewBandController(service *services.BandService) *BandController {
	return &BandController{service: service}
}

func (c *BandController) CreateBand(ctx *gin.Context) {
	var request struct {
		Name        string `json:"name"`
		FormedDate  string `json:"formed_date"`
		EndedDate   string `json:"ended_date"`
		Description string `json:"description"`
		IsActive    bool   `json:"is_active"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	formedDate, err := time.Parse("2006-01-02", request.FormedDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "formed_date must be in YYYY-MM-DD format",
		})
		return
	}

	endedDate, err := time.Parse("2006-01-02", request.EndedDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "ended_date must be in YYYY-MM-DD format",
		})
		return
	}

	band, err := c.service.CreateBand(ctx, request.Name, formedDate, endedDate, request.Description, request.IsActive)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusCreated, band)

}

func (c *BandController) GetBands(ctx *gin.Context) {
	bands, err := c.service.GetBands(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, bands)
}

func (c *BandController) GetBandById(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid band id",
		})
		return
	}

	band, err := c.service.GetBandById(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "band not found",
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch band",
		})
		return
	}

	ctx.JSON(http.StatusOK, band)
}

func (c *BandController) UpdateBandById(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid band id",
		})
		return
	}

	var request struct {
		Name        string `json:"name"`
		FormedDate  string `json:"formed_date"`
		EndedDate   string `json:"ended_date"`
		Description string `json:"description"`
		IsActive    bool   `json:"is_active"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	formedDate, err := time.Parse("2006-01-02", request.FormedDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "formed_date must be in YYYY-MM-DD format",
		})
		return
	}

	endedDate, err := time.Parse("2006-01-02", request.EndedDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "ended_date must be in YYYY-MM-DD format",
		})
		return
	}

	band := &models.Band{
		ID:          id,
		Name:        request.Name,
		FormedDate:  formedDate,
		EndedDate:   endedDate,
		Description: request.Description,
		IsActive:    request.IsActive,
	}

	err = c.service.UpdateBandById(ctx, band)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "band not found",
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to update band",
		})
		return
	}

	ctx.JSON(http.StatusOK, band)

}

func (c *BandController) DeleteBandById(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid band id",
		})
		return
	}

	err = c.service.DeleteBandById(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "band not found",
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to delete band",
		})
		return
	}

	ctx.Status(http.StatusNoContent)

}
