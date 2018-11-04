package main

import (
	"log"
	"os"

	sp "github.com/SparkPost/gosparkpost"

	bugsnag "github.com/bugsnag/bugsnag-go"
	// pb "github.com/anvartech/budget-it-app/mailer/proto"
	// "github.com/anvartech/budget-it-app/mailer/server"
)

// Emailclient should be available everywhere after initialisation
var Emailclient sp.Client

// SparkPostInit initialises Spark Post mailer - may be moved
func SparkPostInit() error {
	apiKey := os.Getenv("SPARKPOST_API_KEY")
	cfg := &sp.Config{
		BaseUrl:    "https://api.sparkpost.com",
		ApiKey:     apiKey,
		ApiVersion: 1,
	}

	err := Emailclient.Init(cfg)

	return err

}

func main() {
	// Get our API key from the environment; configure.

	bugsnagKey := os.Getenv("BUGSNAG_API_KEY")

	bugsnag.Configure(bugsnag.Configuration{
		APIKey: bugsnagKey,
		ProjectPackages: []string{
			"main",
			"github.com/anvartech/budget-it-app/mailer/**",
		},
	})

	err := SparkPostInit()

	if err != nil {
		bugsnag.Notify(err)
		log.Fatalf("SparkPost client init failed: %s\n", err)
	}

	log.Printf("Evrything success\n\n")

}
