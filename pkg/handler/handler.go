package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) initRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.Post("/sign-up")
		auth.Post("/sign-in")
	}
}
