package router

import (
	"github.com/fredianto2405/nusapos-api/pkg/errors"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"time"
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

	return r
}
