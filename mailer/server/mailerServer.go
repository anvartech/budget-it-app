package server

import (
	"context"
	"log"

	"github.com/matcornic/hermes"

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

	var emailTemplate hermes.Email

	emailTemplate = templates.WelcomeEmail(firstName, lastName, token)

	h := templates.HermesInit()

	emailHTML, err := h.GenerateHTML(emailTemplate)
	if err != nil {
		return nil, err
	}
	emailPlain, err := h.GeneratePlainText(emailTemplate)
	if err != nil {
		log.Fatalf("Plain text not generated: %s\n", err)
	}

	log.Printf("Generated Emails :%s \n\n\n Generated Plain %s \n\n email addr %s", emailHTML, emailPlain, emailAddr)

	return nil, nil
}
