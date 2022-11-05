package services

import (
	"auth-api/models/dtos"
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"

	"github.com/spf13/viper"
)



type EmailMessage struct {
	Subject  string
	Receiver string
	Data     map[string]string
}

func DispatchEmail(message *EmailMessage) *dtos.ErrorDto {

	host := viper.Get("EMAIL_HOST_ADDRESS").(string)
	port := viper.Get("EMAIL_HOST_PORT").(string)
	account := viper.Get("EMAIL_USER_ID").(string)
	password := viper.Get("EMAIL_USER_PASSWORD").(string)


	toList := []string{message.Receiver}

	// set relative path for main.go file
	emailTemplate, _ := template.ParseFiles("./email_template.html")
	templateData := struct {
		Name string
		URL  string
	}{Name: message.Data["Name"],
		URL: message.Data["URL"]}

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: %s \n%s\n\n", message.Subject, mimeHeaders)))

	if err := emailTemplate.Execute(&body, templateData); err != nil {
		return &dtos.ErrorDto{Message: err.Error()}
	}

	auth := smtp.PlainAuth("", account, password, host)

	err := smtp.SendMail(host+":"+port, auth, account, toList, body.Bytes())

	if err != nil {
		return &dtos.ErrorDto{Message: err.Error()}
	}
	return nil
}
