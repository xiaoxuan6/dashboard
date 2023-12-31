package lib

import (
    "dashboard/database"
    "dashboard/pkg/jwts"
    "encoding/json"
    "fmt"
    "github.com/golang-jwt/jwt"
    "io/ioutil"
    "net/http"
    "os"
    "strconv"
    "strings"
    "time"
)

type LoginHandler struct {
}

func (l LoginHandler) Run() *Response {
    return nil
}

type LoginRequest struct {
    Email  string `json:"email"`
    Passwd string `json:"passwd"`
}

type LoginResponse struct {
    Token string `json:"token"`
    Email string `json:"email"`
}

func (l LoginHandler) Do(r *http.Request) *Response {
    b, _ := ioutil.ReadAll(r.Body)

    var request LoginRequest
    err := json.Unmarshal(b, &request)
    if err != nil {
        return FailWithMsg(fmt.Sprintf("json 解析请求惨错误：%s", err))
    }

    if request.Email == "" || request.Passwd == "" {
        return FailWithMsg("用户名或密码为空！")
    }

    database.Init()
    passwd, ok := database.Users[request.Email]
    if !ok {
        return FailWithMsg("用户名不存在！")
    }

    if strings.Compare(request.Passwd, passwd.Password) != 0 {
        return FailWithMsg("密码错误！")
    }

    expiresAt, _ := strconv.Atoi(os.Getenv("JWT_EXPIRES_AT"))
    duration := time.Duration(expiresAt) * time.Hour
    c := jwts.MyClaims{
        Email:    request.Email,
        Password: request.Passwd,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: time.Now().Add(duration).Unix(),
        },
        Time: time.Now().Format("2006-01-02 03:04:05"),
    }

    token, err := jwts.GenerateToken(c)
    if err != nil {
        return FailWithMsg("登录失败")
    }

    response := &LoginResponse{
        Token: token,
        Email: request.Email,
    }

    return SuccessWithData(response)
}
