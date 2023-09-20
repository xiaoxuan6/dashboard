package routers

import (
    "dashboard/handlers"
    "github.com/gin-gonic/gin"
    "net/http"
)

func RegisterRouter(r *gin.RouterGroup) {
    r.GET("/", func(context *gin.Context) {
        context.HTML(http.StatusOK, "/api_index", nil)
    })

    r.POST("/dirtyfilter", handlers.DirtryfilterHandler.Filter)
}
