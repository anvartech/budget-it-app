package main

import (
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/anvartech/budget-it-app/mailer/proto"
	server "github.com/anvartech/budget-it-app/mailer/server"
	helpers "github.com/anvartech/budget-it-app/mailer/shared"

	bugsnag "github.com/bugsnag/bugsnag-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

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

	port := helpers.GetEnv("MAILER_PORT", "3081")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		bugsnag.Notify(err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterMailerServiceServer(grpcServer, &server.MailerServer{})

	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
