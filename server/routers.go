package server

import (
	"github.com/gin-gonic/gin"
	"moges/server/handlers"
)

// RegisterAPI to register apis
func RegisterAPI(router *gin.RouterGroup) {
	handlers.RegisterPhotoAPI(router.Group("/upload"))
}
