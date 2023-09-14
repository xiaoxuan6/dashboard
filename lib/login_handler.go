package lib

import (
    "dashboard/pkg/jwts"
    "github.com/golang-jwt/jwt"
    "net/http"
    "os"
    "strconv"
    "time"
)

type LoginHandler struct {
}

func (l LoginHandler) Run() *Response {
    return nil
}

type LoginRepose struct {
    Token string `json:"token"`
    Email string `json:"email"`
}

func (l LoginHandler) Do(r *http.Request) *Response {
    _ = r.ParseForm()

    email := r.FormValue("email")
    passwd := r.FormValue("passwd")

    if email == "" || passwd == "" {
        return failWithMsg("用户名或密码错误！")
    }

    // todo: 判断用户是否存在

    expiresAt, _ := strconv.Atoi(os.Getenv("JWT_EXPIRES_AT"))
    duration := time.Duration(expiresAt) * time.Hour
    c := jwts.MyClaims{
        Email:    email,
        Password: passwd,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: time.Now().Add(duration).Unix(),
        },
    }

    token, err := jwts.GenerateToken(c)
    if err != nil {
        return failWithMsg("登录失败")
    }

    var response LoginRepose
    response.Email = email
    response.Token = token

    return successWithData(response)
}
