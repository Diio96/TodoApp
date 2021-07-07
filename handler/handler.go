package handler

import (
	"github.com/Diio96/todo-app/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}
func NewHandler(services *service.Service) *Handler{
	return &Handler{services: services}
}
func (h *Handler) InitRoutes() *gin.Engine{
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getListById)
			lists.DELETE("/:id", h.deleteList)
			lists.PUT("/:id", h.updateList)

			items := lists.Group(":id/items")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getAllItems)
				items.GET("/:id", h.getItemById)
				items.PUT("/:id", h.updateItem)
				items.DELETE("/:id", h.deleteItem)

			}
		}
	}
	return router
}