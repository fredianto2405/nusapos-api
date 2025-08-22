package router

import (
	"github.com/fredianto2405/nusapos-api/internal/branch"
	"github.com/fredianto2405/nusapos-api/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterBranchRoutes(rg *gin.RouterGroup, handler *branch.Handler) {
	rg.Use(middleware.JWTAuthMiddleware())
	rg.POST("", handler.Add)
	rg.GET("", handler.GetAll)
	rg.PUT("/:id", handler.Edit)
	rg.DELETE("/:id", handler.Delete)
}
