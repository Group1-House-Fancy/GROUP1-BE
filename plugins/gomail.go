package plugins

import (
	"log"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

var config_smtp_host = os.Getenv("CONFIG_SMTP_HOST")
var config_smtp_port = os.Getenv("CONFIG_SMTP_PORT")
var config_sender_name = os.Getenv("CONFIG_SENDER_NAME")
var config_auth_email = os.Getenv("CONFIG_AUTH_EMAIL")
var config_auth_password = os.Getenv("CONFIG_AUTH_PASSWORD")

func SendingEmail(to string) {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", config_sender_name)
	mailer.SetHeader("To", to)
	// mailer.SetAddressHeader("Cc", "tralalala@gmail.com", "Tra Lala La")
	mailer.SetHeader("Subject", "Selamat")
	mailer.SetBody("text/html", "Hello, <b>Terima kasih sudah membeli rumah di House Fancy</b>")
	// mailer.Attach("./bro.png")
	Config_Smtp_Port, _ := strconv.Atoi(config_smtp_port)
	dialer := gomail.NewDialer(
		config_smtp_host,
		Config_Smtp_Port,
		config_auth_email,
		config_auth_password,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Mail sent!")
}
