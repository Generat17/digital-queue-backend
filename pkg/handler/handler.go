package handler

import (
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"server/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Use(cors.Default()) // отключяем CORS политику

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // Swagger

	/*auth := router.Group("/auth")
	{
		//auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}*/

	// доработать
	//api := router.Group("/api", h.userIdentity)
	api := router.Group("/api")
	{
		// api для операций с сотрудниками
		employee := api.Group("/employee")
		{
			employee.GET("", h.getAllLists)
			//employee.POST("", h.createEmployee)
			//employee.DELETE("", h.deleteEmployee)
			//employee.PATCH("", h.updateEmployee)
		}

		// api для операций с обязанностями (услугами)
		responsibility := api.Group("/responsibility")
		{
			responsibility.GET("", h.getResponsibilityList)
		}

		// api для операций с очередью (ticket - это элемент массива очереди)
		queue := api.Group("/queue")
		{
			queue.GET("", h.getQueueLists)
			queue.GET(":service", h.addQueueItem)
			//ticket.GET(":service", h.createTicket)
		}
	}

	return router
}
