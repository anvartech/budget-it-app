package templates

import (
	"os"

	"github.com/matcornic/hermes"
)

// HermesInit initialises Hermes generator
func HermesInit() hermes.Hermes {
	webAddr := os.Getenv("WEB_ADDR")
	logo := os.Getenv("WEB_LOGO")

	return hermes.Hermes{
		// Optional Theme
		Theme: new(hermes.Flat),
		Product: hermes.Product{
			// Appears in header & footer of e-mails
			Name:      "Budget it",
			Link:      webAddr,
			Logo:      logo,
			Copyright: "Copyright © 2018 Anvar Tech. All rights reserved.",
		},
	}
}
