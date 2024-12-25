package router

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mss-boot-io/mss-boot/pkg/response"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Init(r *gin.RouterGroup) {
	v1 := r.Group("/v1")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	configCors := cors.DefaultConfig()
	configCors.AllowOrigins = []string{"*"}
	configCors.AllowCredentials = true
	configCors.AddAllowMethods("Authorization")
	v1.Use(cors.New(configCors))
	v1.OPTIONS("/*path", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	for i := range response.Controllers {
		response.Controllers[i].Other(r.Group("/api", cors.New(configCors)))
		e := v1.Group(response.Controllers[i].Path(), response.Controllers[i].Handlers()...)
		if action := response.Controllers[i].GetAction(response.Get); action != nil {
			e.GET("/:"+response.Controllers[i].GetKey(), action.Handler()...)
		}
		if action := response.Controllers[i].GetAction(response.Control); action != nil {
			e.POST("", action.Handler()...)
			e.PUT("/:"+response.Controllers[i].GetKey(), action.Handler()...)
		}
		if action := response.Controllers[i].GetAction(response.Create); action != nil {
			e.POST("", action.Handler()...)
		}
		if action := response.Controllers[i].GetAction(response.Update); action != nil {
			e.PUT("/:"+response.Controllers[i].GetKey(), action.Handler()...)
		}
		if action := response.Controllers[i].GetAction(response.Delete); action != nil {
			e.DELETE("/:"+response.Controllers[i].GetKey(), action.Handler()...)
		}
		if action := response.Controllers[i].GetAction(response.Search); action != nil {
			e.GET("", action.Handler()...)
		}
	}
}
