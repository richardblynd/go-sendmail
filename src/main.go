package main

import (
	"fmt"
	"net/smtp"
	"strings"
)

func main() {
	SendMail([]string{"aaa@bbb.com", "bbb@bbb.com"}, "teste", "teste")
	fmt.Println("Hello, World!")
}

func SendMail(to []string, subject string, message string) error {

	// Sender data
	from := "xyz@ddd.com"
	username := "xyz@ddd.com"
	password := "hahaha"

	// smtp server configuration
	smtpHost := "smtp.heretheservername.com"
	smtpPort := "587"

	fromText := "From: " + from + "\r\n"
	toText := "To: " + strings.Join(to[:], ",") + "\r\n"
	subject = "Subject: " + subject + "\r\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	msg := []byte(fromText + toText + subject + mime + "\r\n<html><body>" + message + "</body></html>")

	// Authentication
	auth := smtp.PlainAuth("", username, password, smtpHost)

	// Sending email
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, msg)
	return err
}
