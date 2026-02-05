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

type BandMemberController struct {
	service *services.BandMemberService
}

func NewBandmemberController(service *services.BandMemberService) *BandMemberController {
	return &BandMemberController{service: service}
}

func (c *BandMemberController) CrateBandMember(ctx *gin.Context) {
	var request struct {
		BandId    int64  `json:"band_id"`
		ArtistId  int64  `json:"artist_id"`
		RoleId    int64  `json:"role_id"`
		StartDate string `json:"start_date"`
		EndDate   string `json:"end_date"`
	}

	if err := ctx.ShouldBindBodyWithJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	startDate, err := time.Parse("2006-01-02", request.StartDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "start_date must be in YYYY-MM-DD format",
		})
		return
	}

	endDate, err := time.Parse("2006-01-02", request.EndDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "end_date must be in YYYY-MM-DD format",
		})
		return
	}

	DTObandMember := &models.BandMember{
		BandId:    request.BandId,
		ArtistId:  request.ArtistId,
		RoleId:    request.RoleId,
		StartDate: startDate,
		EndDate:   endDate,
	}

	bandMember, err := c.service.CreateBandMember(ctx, DTObandMember)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, bandMember)
}

func (c *BandMemberController) GetBandMembers(ctx *gin.Context) {
	bandMembers, err := c.service.GetBandMembers(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, bandMembers)
}

func (c *BandMemberController) GetBandMemberById(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid band member id",
		})
		return
	}

	bandMember, err := c.service.GetBandMemberById(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "band member not found",
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch band member",
		})
		return
	}

	ctx.JSON(http.StatusOK, bandMember)
}

func (c *BandMemberController) UpdateBandMemberById(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid band member id",
		})
		return
	}

	var request struct {
		BandId    int64  `json:"band_id"`
		ArtistId  int64  `json:"artist_id"`
		RoleId    int64  `json:"role_id"`
		StartDate string `json:"start_date"`
		EndDate   string `json:"end_date"`
	}

	if err := ctx.ShouldBindBodyWithJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	startDate, err := time.Parse("2006-01-02", request.StartDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "start_date must be in YYYY-MM-DD format",
		})
		return
	}

	endDate, err := time.Parse("2006-01-02", request.EndDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "end_date must be in YYYY-MM-DD format",
		})
		return
	}

	bandMember := &models.BandMember{
		ID:        id,
		BandId:    request.BandId,
		ArtistId:  request.ArtistId,
		RoleId:    request.RoleId,
		StartDate: startDate,
		EndDate:   endDate,
	}

	err = c.service.UpdateBandMemberById(ctx, bandMember)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "band member not found",
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to update band member",
		})
		return
	}

	ctx.JSON(http.StatusOK, bandMember)

}

func (c *BandMemberController) DeleteBandMemberById(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid band member id",
		})
		return
	}

	err = c.service.DeleteBandmemberById(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "band member not found",
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to delete band member",
		})
		return
	}

	ctx.Status(http.StatusNoContent)
}
