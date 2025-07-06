package middleware

import (
	"backend/dao"
	"backend/pkg/response"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	tokenExpireDuration = 7 * time.Hour * 24
)

var (
	CustomSecret = []byte("hahaha")
	d            *dao.Dao
)

// InitJWT 初始化JWT中间件
func InitJWT(dao *dao.Dao) {
	d = dao
}

type UserClaims struct {
	Username string `json:"username"`
	ID       uint   `json:"id"`
}

type CustomClaims struct {
	UserClaims
	jwt.RegisteredClaims
}

// GenToken 生成token
func GenToken(Id uint, username string) (string, error) {
	user := UserClaims{
		Username: username,
		ID:       Id,
	}
	claims := CustomClaims{
		user,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenExpireDuration)),
			Issuer:    "xunxun",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(CustomSecret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return CustomSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			response.Fail(c, response.TokenErrCode)
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			response.Fail(c, response.TokenErrCode)
			return
		}

		tokenString := parts[1]

		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := ParseToken(tokenString)
		if err != nil {
			response.Fail(c, response.TokenErrCode)
			return
		}
		// 将当前请求的username信息保存到请求的上下文c上
		c.Set("username", mc.Username)
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}
