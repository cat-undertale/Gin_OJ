package test

import (
	"crypto/tls"
	"github.com/jordan-wright/email"
	"net/smtp"
	"testing"
)

func TestSendEmail(t *testing.T) {
	e := email.NewEmail()
	e.From = "Get <979919125@qq.com>"
	e.To = []string{"2473296046@qq.com"}
	e.Subject = "验证码发送测试"
	e.HTML = []byte("您的验证码<b>123456</b>")
	//err := e.Send("smtp.163.com:456", smtp.PlainAuth("", "979919125@qq.com", "#@gmz20001027", "smtp.163.com"))
	//返回EOF时，关闭SSL重试
	err := e.SendWithTLS("smtp.qq.com:465",
		smtp.PlainAuth("", "979919125@qq.com", "veuiokyamvovbbeb", "smtp.qq.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.qq.com"})
	if err != nil {
		t.Fatal(err)
	}
}
