package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/achsanit/my-gram/internal/middleware"
	"github.com/achsanit/my-gram/internal/model"
	"github.com/achsanit/my-gram/internal/service"
	"github.com/gin-gonic/gin"
)

type PhotoHandler interface {
	PostPhoto(ctx *gin.Context)
	GetPhotosUser(ctx *gin.Context)

	GetPhotoByID(ctx *gin.Context)
}

type photoHandlerImpl struct {
	svc service.PhotoService
}

// GetPhotoByID implements PhotoHandler.
func (p *photoHandlerImpl) GetPhotoByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	photo, err := p.svc.GetPhotoByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   photo,
	})
}

// GetPhotosUser implements PhotoHandler.
func (p *photoHandlerImpl) GetPhotosUser(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Query("user_id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	photos, err := p.svc.GetPhotosUser(ctx, userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   photos,
	})
}

// PostPhoto implements PhotoHandler.
func (p *photoHandlerImpl) PostPhoto(ctx *gin.Context) {
	photoReq := model.CreatePhoto{}

	if err := ctx.Bind(&photoReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := photoReq.ValidateInput(); err != nil {
		var errData map[string]string
		errRes := json.Unmarshal([]byte(err.Error()), &errData)
		if errRes != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"message": errData})
		return
	}

	photo := model.Photo{
		Title:   photoReq.Title,
		Url:     photoReq.Url,
		Caption: photoReq.Caption,
		UserID:  uint64(int(ctx.GetFloat64(middleware.CLAIM_USER_ID))),
	}

	res, err := p.svc.PostPhoto(ctx, photo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"data":   res,
	})
}

func NewPhotoHandler(service service.PhotoService) PhotoHandler {
	return &photoHandlerImpl{
		svc: service,
	}
}
