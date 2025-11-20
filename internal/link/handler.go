package link

import (
	"errors"

	"github.com/gin-gonic/gin"
)

type LinkHandler struct {
	service *LinkService
}

func NewHandler(service *LinkService) *LinkHandler {
	return &LinkHandler{service: service}
}

func (h *LinkHandler) Shorten(c *gin.Context) {
	var req ShortenLinkRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	shortCode, err := h.service.Shorten(c, req.URL)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"shortCode": shortCode})
}

func (h *LinkHandler) Get(c *gin.Context) {
	code := c.Param("code")

	if code == "" {
		c.JSON(400, gin.H{"error": "code is empty"})
		return
	}

	record, err := h.service.GetRecord(c, code)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if record == nil {
		c.JSON(404, gin.H{"error": "Url Not Found"})
		return
	}

	c.JSON(200, record)
}

func (h *LinkHandler) Update(c *gin.Context) {
	var req ShortenLinkRequest

	code := c.Param("code")

	if code == "" {
		c.JSON(400, gin.H{"error": "code is empty"})
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := h.service.UpdateUrl(c, req.URL, code)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.Status(200)
}

func (h *LinkHandler) Delete(c *gin.Context) {
	code := c.Param("code")

	if code == "" {
		c.JSON(400, gin.H{"error": "code is empty"})
		return
	}

	err := h.service.DeleteRecord(c, code)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.Status(204)
}
