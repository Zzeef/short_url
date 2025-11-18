package link

import "github.com/gin-gonic/gin"

func NewRouter(handler *LinkHandler) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/shorten")
	{
		// POST /api/shorten
		api.POST("/", handler.Shorten)

		// GET /api/shorten/:code
		api.GET("/:code", handler.Get)

		// PUT /api/shorten/:code
		api.PUT("/:code", handler.Update)

		// DELETE /api/links/:code
		api.DELETE("/:code", handler.Delete)
	}

	return r
}
