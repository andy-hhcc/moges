package security

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"moges/common"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		ignoreURLs := []string{"/authenticate"}
		if common.ContainStr(c.Request.URL.Path, ignoreURLs) {
			return
		}

		auth := c.Request.Header.Get("Authorization")

		validToken := viper.Get("server.valid_token")
		if auth != validToken {
			_ = c.AbortWithError(http.StatusForbidden, errors.New("User is not authorization"))
			return
		}
	}
}

func PhotoMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}
