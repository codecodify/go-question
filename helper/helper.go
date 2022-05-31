package helper

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/codecodify/go-question/define"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleError(ctx *gin.Context, err error, code int) {
	ctx.JSON(code, gin.H{
		"status": "error",
		"error":  err.Error(),
	})
}

func HandleSuccess(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   data,
	})
}

func Md5(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

type UserClaims struct {
	Identity string `json:"identity"`
	Name     string `json:"name"`
	jwt.StandardClaims
}

var key = []byte(define.DefaultTokenKey)

// GetUserToken 生成token
func GetUserToken(identity, name string) (string, error) {
	claim := &UserClaims{
		Identity: identity,
		Name:     name,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(key)
}

// CheckUserToken 校验token
func CheckUserToken(tokenString string) (*UserClaims, error) {
	claims := new(UserClaims)
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("token is invalid")
	}
	return claims, nil
}
