package jwts

import (
    "fmt"
    "github.com/golang-jwt/jwt"
    "os"
)

type MyClaims struct {
    Email    string `json:"email"`
    Password string `json:"password"`
    Time     string `json:"time"` // 生成 token 时的时间，防止生成的 token 不变
    jwt.StandardClaims
}

var key = os.Getenv("JWT_SIGNED_TOKEN")

func GenerateToken(claims MyClaims) (string, error) {
    t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return t.SignedString([]byte(key))
}

func ParseTokenString(tokenString string) (Claims interface{}, er error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return []byte(key), nil
    })

    if err != nil {
        return nil, err
    }

    if !token.Valid {
        return nil, fmt.Errorf("请求令牌无效")
    }

    return token.Claims, nil
}

func ParseWithClaims(tokenString string, claims *MyClaims) (*MyClaims, error) {
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return []byte(key), nil
    })

    if err != nil {
        return nil, err
    }

    if !token.Valid {
        return nil, fmt.Errorf("请求令牌无效")
    }

    if cl, ok := token.Claims.(*MyClaims); ok {
        return cl, nil
    }

    return nil, fmt.Errorf("请求令牌格式有误")
}
