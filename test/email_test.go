package test

import (
	"github.com/jordan-wright/email"
	"net/smtp"
	"testing"
)

func TestSendEmail(t *testing.T) {
	host := "smtp.qq.com"
	port := "25"
	username := "meet.yuzhang@qq.com"
	password := "ahcgdolgclwohica"

	e := &email.Email{
		From:    username,
		To:      []string{"yu1.zhang@icloud.com"},
		Subject: "Verification Code Sending Test",
		HTML:    []byte("Your Verification Code is<h1>123456</h1>"),
	}

	err := e.Send(host+":"+port, smtp.PlainAuth("", username, password, host))

	if err != nil {
		t.Error(err)
		return
	}
}
