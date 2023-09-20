package routers

import (
    "dashboard/handlers"
    "github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.RouterGroup) {
    r.GET("/", func(context *gin.Context) {
        context.Redirect(301, "/apis/index")
    })

    r.GET("/docs", handlers.DocsHandler.Index)
    r.POST("/dirtyfilter", handlers.DirtryfilterHandler.Filter)
}
