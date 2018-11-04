package templates

import (
	"strings"

	helpers "github.com/anvartech/budget-it-app/mailer/shared"

	"github.com/matcornic/hermes"
)

// WelcomeEmail is the generated email containing a body
func WelcomeEmail(firstName string, lastName string, token string) hermes.Email {

	webAddr := helpers.GetEnv("WEB_ADDR", "http://example.com/")

	var resetLink strings.Builder

	resetLink.WriteString(webAddr)
	resetLink.WriteString("confirm?token=")
	resetLink.WriteString(token)

	return hermes.Email{
		Body: hermes.Body{
			Name: firstName,
			Intros: []string{
				"Welcome! We are so excited to have you on board.",
			},
			Actions: []hermes.Action{
				{
					Instructions: "To get started with Budget, you need to confirm your account",
					Button: hermes.Button{
						Text: "Confirm your account",
						Link: resetLink.String(),
					},
				},
			},
			Outros: []string{
				"If you did not create an account, no further action is required on your part.",
			},
		},
	}
}
