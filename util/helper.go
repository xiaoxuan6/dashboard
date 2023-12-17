package util

import (
    "github.com/gin-gonic/gin"
    "math/rand"
    "net/http"
    "strconv"
    "strings"
    "time"
)

func Contains(s string) func(item string) bool {
    return func(item string) bool {
        return strings.Contains(s, item)
    }
}

func Random(num int, incr int) string {
    rand.Seed(time.Now().UnixNano())

    randomNum := rand.Intn(num) + incr

    return strconv.Itoa(randomNum)
}

func Success(ctx *gin.Context) {
    SuccessWithData(ctx, "")
}

func SuccessWithData(ctx *gin.Context, data interface{}) {
    ctx.JSON(http.StatusOK, gin.H{
        "status": http.StatusOK,
        "data":   data,
        "msg":    "ok",
    })
}

func Fail(ctx *gin.Context, msg string) {
    ctx.JSON(http.StatusOK, gin.H{
        "status": http.StatusBadRequest,
        "data":   "",
        "msg":    msg,
    })
}
