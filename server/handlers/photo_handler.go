package handlers

import (
	"github.com/gin-gonic/gin"
	"moges/common"
	"moges/domain/service"
	"moges/server/handlers/validators"
	"net/http"
)

func RegisterPhotoAPI(router *gin.RouterGroup) {

	router.POST("", Upload)
}

func Upload(c *gin.Context) {

	photoRequest := validators.NewPhotoValidator()

	if err := photoRequest.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": err.Error(),
		})
		return
	}


	if err := service.StorePhoto(photoRequest); err != nil {
		c.JSON(http.StatusInternalServerError, common.NewError("message", err))
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": true,
		"data": "Successful",
	})
}

