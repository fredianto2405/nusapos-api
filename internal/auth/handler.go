package auth

import (
	"github.com/fredianto2405/nusapos-api/pkg/errors"
	"github.com/fredianto2405/nusapos-api/pkg/jwt"
	"github.com/fredianto2405/nusapos-api/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Login(c *gin.Context) {
	var request LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(err)
		return
	}

	if err := errors.Validate.Struct(request); err != nil {
		c.Error(err)
		return
	}

	user, err := h.service.Login(&request)
	if err != nil {
		c.Error(err)
		return
	}

	token, err := jwt.GenerateJWT(user.ID, user.Username, user.Role)
	if err != nil {
		c.Error(err)
		return
	}

	data := LoginResponse{
		AccessToken: token,
		TokenType:   "Bearer",
		ExpiresIn:   43200000,
	}

	response.Respond(c, http.StatusOK, true, "login berhasil", data, nil)
}
