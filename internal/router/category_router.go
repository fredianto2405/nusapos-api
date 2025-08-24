package router

import (
	"github.com/fredianto2405/nusapos-api/internal/category"
	"github.com/fredianto2405/nusapos-api/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterCategoryRoutes(rg *gin.RouterGroup, handler *category.Handler) {
	rg.Use(middleware.JWTAuthMiddleware())
	rg.POST("", handler.Add)
	rg.GET("", handler.GetAll)
	rg.PUT("/:id", handler.Edit)
	rg.DELETE("/:id", handler.Delete)
}
