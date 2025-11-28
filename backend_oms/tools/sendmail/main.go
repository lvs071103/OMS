package main

import (
	"time"

	"gopkg.in/gomail.v2"
)

func main() {
	// 邮箱地址
	from := "zabbix@dominos.com.cn"
	password := "**************"
	toEmail := "Dean.Wen@dominos.com.cn"
	smtpHost := "outlook.office365.com"

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", toEmail)
	m.SetAddressHeader("Cc", from, "Dan")
	m.SetHeader("Subject", "Go语言发送邮件测试")
	m.SetBody("text/html", "测试日期:"+time.Now().Local().String())

	d := gomail.NewDialer(smtpHost, 587, from, password)

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
