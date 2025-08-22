package router

import (
	"time"

	"github.com/fredianto2405/nusapos-api/internal/auth"
	"github.com/fredianto2405/nusapos-api/internal/branch"
	"github.com/fredianto2405/nusapos-api/pkg/errors"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func SetupRouter(db *sqlx.DB) *gin.Engine {
	// init validator
	errors.InitValidator()

	r := gin.Default()

	// middleware cors
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// error handler
	r.Use(errors.ErrorHandler())

	// auth routes
	authRepo := auth.NewRepository(db)
	authService := auth.NewService(authRepo)
	authHandler := auth.NewHandler(authService)
	authGroup := r.Group("/api/v1/auth")
	RegisterAuthRoutes(authGroup, authHandler)

	// branch routes
	branchRepo := branch.NewRepository(db)
	branchService := branch.NewService(branchRepo)
	branchHandler := branch.NewHandler(branchService)
	branchGroup := r.Group("/api/v1/branches")
	RegisterBranchRoutes(branchGroup, branchHandler)

	return r
}
