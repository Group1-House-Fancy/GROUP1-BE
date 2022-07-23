package plugins

import (
	"log"

	"gopkg.in/gomail.v2"
)

const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587
const CONFIG_SENDER_NAME = "PT. House Fancy <housefancy950@gmail.com>"
const CONFIG_AUTH_EMAIL = "housefancy950@gmail.com"
const CONFIG_AUTH_PASSWORD = "lgra xuae hopo qehn"

func SendingEmail(to string) {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", to)
	mailer.SetHeader("Subject", "Selamat")
	mailer.SetBody("text/html", "Hello, <b>Terima kasih sudah membeli rumah di House Fancy</b>")

	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Mail sent!")
}
