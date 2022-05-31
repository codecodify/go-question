package test

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"testing"
)

var key = []byte("golang")

type MyCustomClaims struct {
	Identity string `json:"identity"`
	Name     string `json:"name"`
	jwt.StandardClaims
}

func TestGenerateToken(t *testing.T) {
	claim := new(MyCustomClaims)
	claim.Identity = "123456789"
	claim.Name = "codecodify"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	fmt.Println(token.SignedString(key))
}

func TestValidToken(t *testing.T) {
	tokenStr := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZGVudGl0eSI6IjEyMzQ1Njc4OSIsIm5hbWUiOiJjb2RlY29kaWZ5In0.Uf4lpxNAFjEacGgpM2WnfWDKZgpMSMOikwn8FhWLaBs"
	claims := new(MyCustomClaims)
	token, _ := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if token.Valid {
		fmt.Println("Token is valid", claims)
	}

}
