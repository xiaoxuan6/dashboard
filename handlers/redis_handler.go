package handlers

import (
    "dashboard/pkg/redis"
    "github.com/gin-gonic/gin"
    "net/http"
    "time"
)

var RedisHandler = new(redisHandler)

type redisHandler struct {
}

func (r *redisHandler) Index(c *gin.Context) {
    res, err := redis.Remember("test", 1*time.Minute, func() []string {
        return []string{"test"}
    })

    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "status": http.StatusBadRequest,
            "data":   "",
            "msg":    err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "status": http.StatusOK,
        "data":   res,
        "msg":    "ok",
    })
}
