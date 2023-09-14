package lib

import (
    "dashboard/pkg/jwts"
    "encoding/json"
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
    _ = json.Unmarshal(b, &request)

    if request.Email == "" || request.Passwd == "" {
        return failWithMsg("用户名或密码错误！")
    }

    // todo: 判断用户是否存在

    expiresAt, _ := strconv.Atoi(os.Getenv("JWT_EXPIRES_AT"))
    duration := time.Duration(expiresAt) * time.Hour
    c := jwts.MyClaims{
        Email:    request.Email,
        Password: request.Passwd,
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
        Email: request.Email,
    }
    return successWithData(response)
}
