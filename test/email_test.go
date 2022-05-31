package test

import (
	"crypto/tls"
	"github.com/jordan-wright/email"
	"net/smtp"
	"testing"
)

func TestEmailSend(t *testing.T) {
	e := email.NewEmail()
	e.From = "shaoxingliu@126.com"
	e.To = []string{"79627084@qq.com"}
	e.Subject = "放假通知"
	e.Text = []byte("端午节放假3天")
	err := e.SendWithTLS("smtp.126.com:465",
		smtp.PlainAuth("", "shaoxingliu@126.com", "FBJDYSEHDCVUNHRW", "smtp.126.com"),
		&tls.Config{
			InsecureSkipVerify: true,
			ServerName:         "smtp.126.com",
		})
	if err != nil {
		panic(err.Error())
	}
}
