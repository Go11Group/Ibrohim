package api

import (
	"api-gateway/api/handler"
	"api-gateway/api/middleware"
	"api-gateway/config"

	_ "api-gateway/api/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title E-Commerce
// @version 1.0
// @description API Gateway of E-Commerce
// @host localhost:8080
// @BasePath /e-commerce
// @schemes http
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func NewRouter(cfg *config.Config) *gin.Engine {
	h := handler.NewHandler(cfg)

	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/e-commerce")
	api.Use(middleware.Check)

	admin := api.Group("/admin")
	{
		u := admin.Group("/user")
		{
			u.POST("", h.AddUser)
			u.GET("/:id", h.GetUser)
			u.PUT("/:id", h.UpdateUser)
			u.DELETE("/:id", h.DeleteUser)
			u.GET("/all", h.FetchUsers)
		}

		p := admin.Group("/product")
		{
			p.POST("", h.PostProduct)
			p.GET("/:id", h.GetProduct)
			p.PUT("/:id", h.PutProduct)
			p.DELETE("/:id", h.DeleteProduct)
			p.GET("/all", h.FetchProducts)
		}
	}

	user := api.Group("/user")
	{
		p := user.Group("/profile")
		{
			p.GET("/:id", h.GetProfile)
			p.PUT("/:id", h.UpdateProfile)
			p.DELETE("/:id", h.DeleteProfile)
		}

		prod := user.Group("/product")
		prod.GET("/all", h.FetchProducts)

		b := user.Group("/basket")
		{
			b.POST("", h.AddToBasket)
			b.GET("/:id", h.GetBasket)
			b.PUT("/:id", h.UpdateBasket)
			b.DELETE("/:id", h.RemoveFromBasket)
		}
	}

	return router
}
