package serializers

import (
	"github.com/gin-gonic/gin"
	"moges/domain/model"
)

type PhotoResponse struct {
	ID   uint    `json:"id"`
	Name string `json:"name"`
	Path string `json:"path"`
}

type PhotoSerializer struct {
	C *gin.Context
	model.Photo
}


func (s PhotoSerializer) Response() PhotoResponse {

	response := PhotoResponse{
		ID:   s.ID,
		Name: s.Name,
		Path: s.Path,
	}
	return response
}
