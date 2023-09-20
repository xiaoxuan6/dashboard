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
    app.NoMethod(func(c *gin.Context) {
        path := c.Request.URL.Path
        c.JSON(http.StatusBadRequest, gin.H{
            "status": http.StatusNotFound,
            "msg":    fmt.Sprintf("router %s method %s not found", path, c.Request.Method),
        })
    })
    app.NoRoute(func(c *gin.Context) {
        path := c.Request.URL.Path
        c.JSON(http.StatusOK, gin.H{
            "status": http.StatusNotFound,
            "msg":    fmt.Sprintf("router %s not found", path),
        })
    })

    app.LoadHTMLGlob("html/*")
    r := app.Group("/apis")
    routers.RegisterRouter(r)
}

func Api(w http.ResponseWriter, r *http.Request) {
    app.ServeHTTP(w, r)
}
