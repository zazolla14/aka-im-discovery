package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"

	"github.com/1nterdigital/aka-im-discover/internal/api/http"
	discovermw "github.com/1nterdigital/aka-im-discover/internal/api/mw"
	discoverapi "github.com/1nterdigital/aka-im-discover/internal/service"
	middleware "github.com/1nterdigital/aka-im-tools/mw"
)

func SetRouter(svcName string, api *discoverapi.Api, mw *discovermw.MW) *gin.Engine {
	r := gin.New()

	handler := http.NewDiscoverHandler(api)
	r.GET("/health", handler.Health)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Use(gin.Recovery(), middleware.CorsHandler(), middleware.GinParseOperationID())
	r.Use(mw.GinParseToken())
	r.Use(otelgin.Middleware(svcName))

	article := r.Group("/discover/article")
	article.GET("/find", handler.FindArticles)

	carousel := r.Group("/discover/carousel")
	carousel.GET("/find", handler.FindCarousels)

	bo := r.Group("/bo", mw.CheckAdmin)

	carouselAdmin := bo.Group("/discover/carousel")
	carouselAdmin.GET("/find", handler.FindCarousels)
	carouselAdmin.POST("/add", handler.CreateCarousel)
	carouselAdmin.DELETE("/del", handler.DeleteCarousel)
	carouselAdmin.POST("/edit", handler.EditCarousel)

	articleAdmin := bo.Group("/discover/article")
	articleAdmin.GET("/find", handler.FindArticles)
	articleAdmin.POST("/add", handler.CreateArticle)
	articleAdmin.DELETE("/del", handler.DeleteArticle)
	articleAdmin.POST("/edit", handler.EditArticle)

	return r
}
