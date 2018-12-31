package server

import (
	"errors"
	"log"
	"os"
	"reflect"

	helpers "github.com/anvartech/budget-it-app/mailer/shared"
	bugsnag "github.com/bugsnag/bugsnag-go"

	sendgrid "github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

//
type emailObjects struct {
}

func (s emailObjects) isEmpty() bool {
	return reflect.DeepEqual(s, SendEmailObject{})
}

// SendEmailObject is used to hold basic sender/receiver email structure
type SendEmailObject struct {
	*emailObjects
	fullName     string
	emailAddress string
}

// SendEmailContent is used to hold email Content (plain text and html)
type SendEmailContent struct {
	*emailObjects
	plainText   string
	htmlContent string
}

// SendEmail is used to send email to certian parties
func SendEmail(subject string, emailTo SendEmailObject, emailFrom SendEmailObject, emailContent SendEmailContent) error {

	if emailFrom.isEmpty() {
		emailFrom.fullName = helpers.GetEnv("DEFAULT_WEB_EMAIL_NAME", "budget-it-app")
		emailFrom.emailAddress = helpers.GetEnv("DEFAULT_WEB_EMAIL", "no_reply@example.com")
	}
	if emailTo.isEmpty() {
		err := errors.New("No email address or name of receiver")
		bugsnag.Notify(err)

		return err
	}
	if emailContent.isEmpty() {
		err := errors.New("Nothing to send, is this a test email ?")
		bugsnag.Notify(err)

		return err
	}
	sendGridAPIKey, ok := os.LookupEnv("SENDGRID_API_KEY")
	if !ok {
		err := errors.New("SendGrid API key Environment Variables not set")
		bugsnag.Notify(err)

		log.Panicln("SendGrid API key Environment Variables not set")
		return err
	}

	client := sendgrid.NewSendClient(sendGridAPIKey)

	from := mail.NewEmail(emailFrom.fullName, emailFrom.emailAddress)

	to := mail.NewEmail(emailFrom.fullName, emailFrom.emailAddress)

	message := mail.NewSingleEmail(from, subject, to, emailContent.plainText, emailContent.htmlContent)

	response, err := client.Send(message)
	if err != nil {
		bugsnag.Notify(err)

		return err
	}
	log.Printf("email Sent: %d", response.StatusCode)
	return nil
}
