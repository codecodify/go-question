package helper

import (
	"crypto/md5"
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/codecodify/go-question/define"
	"github.com/codecodify/go-question/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jordan-wright/email"
	uuid "github.com/satori/go.uuid"
	"math/rand"
	"net/http"
	"net/smtp"
	"time"
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

// SendMail 发送邮件
func SendMail(to, code string) error {
	e := email.NewEmail()
	e.From = "shaoxingliu@126.com"
	e.To = []string{to}
	e.Subject = "注册验证码"
	e.Text = []byte(fmt.Sprintf("您好，您的验证码是：%s", code))
	// todo 邮箱测试
	return e.SendWithTLS("smtp.126.com:465",
		smtp.PlainAuth("", "shaoxingliu@126.com", "FBJDYSEHDCVUNHRW", "smtp.126.com"),
		&tls.Config{
			InsecureSkipVerify: true,
			ServerName:         "smtp.126.com",
		})
}

// GetUUID 获取uuid
func GetUUID() string {
	return uuid.NewV4().String()
}

// GetRandomCode 生成随机验证码
func GetRandomCode() string {
	code := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 6; i++ {
		code += fmt.Sprintf("%d", rand.Intn(10))
	}
	return code
}

// SetRedisString 在redis存储字符串
func SetRedisString(key, value string) error {
	return models.Redis.Set(models.RedisCtx, key, value, time.Duration(define.DefaultTokenExpire)*time.Second).Err()
}

// GetRedisString 在redis获取字符串
func GetRedisString(key string) (string, error) {
	return models.Redis.Get(models.RedisCtx, key).Result()
}
