package handler

import (
	docs "github.com/Dormant512/url-compressor/backend/docs"
	"github.com/Dormant512/url-compressor/backend/internal/service"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterHandlers(svc *service.Service) *gin.Engine {
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/"
	router.POST("/compressor", svc.CompressURL)
	router.GET("/compressor/:compressed_url", svc.RedirectURL)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}
