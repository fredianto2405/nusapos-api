package branch

import (
	"net/http"
	"strconv"

	"github.com/fredianto2405/nusapos-api/pkg/errors"
	"github.com/fredianto2405/nusapos-api/pkg/response"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetAll(c *gin.Context) {
	search := c.DefaultQuery("search", "")
	pagination := c.DefaultQuery("pagination", "false") == "true"
	if !pagination {
		branches, err := h.service.GetAll(search)
		if err != nil {
			c.Error(err)
			return
		}

		response.Respond(c, http.StatusOK, true, "berhasil mengambil data cabang", branches, nil)
		return
	}

	pageQuery := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageQuery)
	if err != nil {
		c.Error(err)
		return
	}

	sizeQuery := c.DefaultQuery("size", "10")
	size, err := strconv.Atoi(sizeQuery)
	if err != nil {
		c.Error(err)
		return
	}

	branches, total, err := h.service.GetAllPageable(page, size, search)
	if err != nil {
		c.Error(err)
		return
	}

	meta := response.NewMeta(total, page, size)
	response.Respond(c, http.StatusOK, true, "berhasil mengambil data cabang", branches, meta)
}

func (h *Handler) Add(c *gin.Context) {
	var request Request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(err)
		return
	}

	if err := errors.Validate.Struct(&request); err != nil {
		c.Error(err)
		return
	}

	if err := h.service.Create(&request); err != nil {
		c.Error(err)
		return
	}

	response.Respond(c, http.StatusOK, true, "cabang berhasil ditambahkan", nil, nil)
}

func (h *Handler) Edit(c *gin.Context) {
	id := c.Param("id")
	var request Request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(err)
		return
	}

	if err := errors.Validate.Struct(&request); err != nil {
		c.Error(err)
		return
	}

	if err := h.service.Update(id, &request); err != nil {
		c.Error(err)
		return
	}

	response.Respond(c, http.StatusOK, true, "data cabang berhasil diubah", nil, nil)
}

func (h *Handler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.Delete(id); err != nil {
		c.Error(err)
		return
	}
	
	response.Respond(c, http.StatusOK, true, "cabang berhasil dihapus", nil, nil)
}
