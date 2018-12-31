package server

import (
	"context"
	"log"
	"strings"

	bugsnag "github.com/bugsnag/bugsnag-go"
	"github.com/matcornic/hermes"

	helpers "github.com/anvartech/budget-it-app/mailer/shared"
	templates "github.com/anvartech/budget-it-app/mailer/templates"

	pb "github.com/anvartech/budget-it-app/mailer/proto"
)

// MailerServer is class for welcome mail
type MailerServer struct {
}

// SendWelcomeMail is used to send welcome mail
func (s *MailerServer) SendWelcomeMail(ctx context.Context, request *pb.WelcomeMailRequest) (*pb.Response, error) {
	// generate template

	userDetails := request.GetDetails()
	token := request.GetToken()

	firstName := userDetails.FirstName
	lastName := userDetails.LastName
	emailAddr := userDetails.GetEmail()

	var fullName strings.Builder
	fullName.WriteString(firstName)
	fullName.WriteString(" ")
	fullName.WriteString(lastName)

	var emailTemplate hermes.Email

	emailTemplate = templates.WelcomeEmail(firstName, lastName, token)

	h := templates.HermesInit()

	emailHTML, err := h.GenerateHTML(emailTemplate)
	if err != nil {
		bugsnag.Notify(err)

		return &pb.Response{
			SendSuccess: false,
			Errors: &pb.Error{
				Description: err.Error(),
			},
		}, err
	}
	emailPlain, err := h.GeneratePlainText(emailTemplate)
	if err != nil {
		bugsnag.Notify(err)

		log.Fatalf("Plain text not generated: %s\n", err)
	}

	emailTo := SendEmailObject{
		fullName:     fullName.String(),
		emailAddress: emailAddr,
	}
	emailFROM := SendEmailObject{
		fullName:     helpers.GetEnv("APP_NAME", "Your Budget App"),
		emailAddress: emailAddr,
	}
	emailContent := SendEmailContent{
		plainText:   emailPlain,
		htmlContent: emailHTML,
	}

	if err := SendEmail("Welcome to Budget it App! Please Confirm your Email", emailTo, emailFROM, emailContent); err != nil {
		log.Printf("Something went wrong sending the email %s\n", err)
		bugsnag.Notify(err)

		return &pb.Response{
			SendSuccess: false,
			Errors: &pb.Error{
				Description: err.Error(),
			},
		}, err
	}

	return &pb.Response{
		SendSuccess: true,
		Errors:      nil,
	}, nil
}
