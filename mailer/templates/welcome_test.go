package templates

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWelcomeEmail(t *testing.T) {
	h := HermesInit()
	type args struct {
		firstName string
		lastName  string
		token     string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Testing welcome email generatedHTML and generatedText",
			args: args{
				firstName: "TestUser",
				lastName:  "TestUserLastName",
				token:     "TestToken",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := WelcomeEmail(tt.args.firstName, tt.args.lastName, tt.args.token)
			r, err := h.GenerateHTML(got)

			t.Log(r)
			assert.Nil(t, err)
			assert.NotEmpty(t, r)

			assert.Contains(t, r, "Hi "+tt.args.firstName, "Name: Should test that a Proper Introduction existed")
			assert.Contains(t, r, "Welcome! We are so excited to have you on board.", "Intro: Should have good intro")
			assert.Contains(t, r, "http://example.com/confirm?token="+tt.args.token, "Action: Button should have link")
			assert.Contains(t, r, "If you did not create an account, no further action is required on your part.", "Outro: Should have outro")
			// 	t.Errorf("WelcomeEmail() = %v, want %v", got, tt.want)
			// }
		})
	}
}
