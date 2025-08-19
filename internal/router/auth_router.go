package router

import (
	"github.com/fredianto2405/nusapos-api/internal/auth"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(rg *gin.RouterGroup, handler *auth.Handler) {
	rg.POST("/login", handler.Login)
}
