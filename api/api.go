package api

import (
    "dashboard/routers"
    "fmt"
    "github.com/gin-gonic/gin"
    "net/http"
)

var app *gin.Engine

func init() {
    gin.SetMode(gin.ReleaseMode)
    app = gin.New()
    app.NoRoute(func(c *gin.Context) {
        path := c.Request.URL.Path
        c.JSON(http.StatusOK, gin.H{
            "status": http.StatusNotFound,
            "msg":    fmt.Sprintf("route %s not found", path),
        })
    })
    app.NoMethod(func(c *gin.Context) {
        path := c.Request.URL.Path
        c.JSON(http.StatusOK, gin.H{
            "status": http.StatusBadRequest,
            "msg":    fmt.Sprintf("route %s method %s not allow", path, c.Request.Method),
        })
    })

    r := app.Group("/apis")
    routers.RegisterRouter(r)
}

func Api(w http.ResponseWriter, r *http.Request) {
    app.ServeHTTP(w, r)
}
