package link

import "github.com/gin-gonic/gin"

type LinkHandler struct {
	service *LinkService
}

func NewHandler(service *LinkService) *LinkHandler {
	return &LinkHandler{service: service}
}

func (l *LinkHandler) Shorten(c *gin.Context) {

}

func (l *LinkHandler) Get(c *gin.Context) {

}

func (l *LinkHandler) Update(c *gin.Context) {

}

func (l *LinkHandler) Delete(c *gin.Context) {

}
