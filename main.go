package main

import (
	"fmt"
	"net/mail"

	config "github.com/OzqurYalcin/mailer/config"
	mailer "github.com/OzqurYalcin/mailer/src"
)

func init() {
	config.MailHost = "" // Mail Host
	config.MailPort = "" // Mail Port
	config.MailUser = "" // Kullanıcı Adı
	config.MailPass = "" // Şifre
}

func main() {
	api := new(mailer.API)
	api.Lock()
	defer api.Unlock()
	request := mailer.Request{}
	request.Body.From = mail.Address{"Name", "mail@example.com"}
	request.Body.To = mail.Address{"Name", "mail@example.com"}
	request.Body.Subject = "Title"
	request.Body.Msg = "Message"
	send := api.Mail(request)
	if send {
		fmt.Println("e-posta iletildi")
	} else {
		fmt.Println("hata oluştu")
	}
}