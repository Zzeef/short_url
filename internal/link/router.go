package link

import "github.com/gin-gonic/gin"

func NewRouter(handler *LinkHandler) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/shorten")
	{
		// POST /api/shorten
		api.POST("/", handler.Shorten)

		// GET /api/shorten/:id
		api.GET("/:id", handler.Get)

		// PUT /api/shorten/:id
		api.PUT("/:id", handler.Update)

		// DELETE /api/links/:id
		api.DELETE("/:id", handler.Delete)
	}

	return r
}
