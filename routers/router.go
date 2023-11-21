package routers

import (
    "dashboard/handlers"
    "github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.RouterGroup) {
    r.GET("/", func(context *gin.Context) {
        context.Redirect(301, "/apis/index")
    })

    r.GET("/docs/:id", handlers.DocsHandler.Show)
    r.GET("/docs", handlers.DocsHandler.Index)
    r.POST("/dirtyfilter", handlers.DirtryfilterHandler.Filter)
    r.POST("/collect", handlers.CollectHandler.Put)
    r.POST("/email_check", handlers.EmailHandler.Check)
    r.GET("/random_img", handlers.ImageHandler.Random)

    r.GET("/redis", handlers.RedisHandler.Index)
}
