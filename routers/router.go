package routers

import (
    "dashboard/handlers"
    "github.com/gin-gonic/gin"
    "net/http"
)

func RegisterRouter(r *gin.RouterGroup) {
    r.GET("/", func(context *gin.Context) {
        context.JSON(http.StatusOK, gin.H{
            "status": 200,
            "msg":    "ok",
        })
    })

    r.POST("/dirtyfilter", handlers.DirtryfilterHandler.Filter)
}
