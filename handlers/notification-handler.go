package handlers

import (
	"bytes"
	"errors"
	"html/template"

	"gopkg.in/gomail.v2"
	"guptalibrary.com/configs"
	"guptalibrary.com/models"
)

func SendEmail(templateName, subject string, dynamicValues map[string]any, senderEmails ...string) (string, error) {

	notificationTemplate := &models.NotificationTemplate{}
	configs.DB.Where("name = ? and type = ?", templateName, "email").First(notificationTemplate)

	if notificationTemplate.Id == 0 || !notificationTemplate.IsActive {
		return "", errors.New("notification template not found")
	}

	message := gomail.NewMessage()

	emailConfig := configs.LoadEmailConfig()

	message.SetHeader("From", emailConfig.SMTPUsername)
	message.SetHeader("To", senderEmails...)
	message.SetHeader("Subject", subject)

	tmpl, tmplError := template.New("email").Parse(notificationTemplate.Content)
	if tmplError != nil {
		return "", errors.New("unable to parse notification template")
	}

	var body bytes.Buffer
	execError := tmpl.Execute(&body, dynamicValues)
	if execError != nil {
		return "", errors.New("unable to execute notification template")
	}

	message.SetBody("text/html", body.String())

	dialer := gomail.NewDialer(emailConfig.SMTPHost, emailConfig.SMTPPort, emailConfig.SMTPUsername, emailConfig.SMTPPassword)

	sendingError := dialer.DialAndSend(message)
	if sendingError != nil {
		return "", errors.New("unable to send email due to technical issue")
	}

	return "Email is sent successfully!", nil
}
