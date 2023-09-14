package lib

import (
    "dashboard/pkg/jwts"
    "github.com/golang-jwt/jwt"
    "io/ioutil"
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

type LoginResponse struct {
    Token string `json:"token"`
    Email string `json:"email"`
}

func (l LoginHandler) Do(r *http.Request) *Response {
    b, _ := ioutil.ReadAll(r.Body)
    return successWithData(string(b))

    email := r.PostFormValue("email")
    passwd := r.PostFormValue("passwd")

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

    response := &LoginResponse{
        Token: token,
        Email: email,
    }
    return successWithData(response)
}
