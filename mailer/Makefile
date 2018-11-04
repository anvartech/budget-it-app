build:
	protoc -I. --go_out=plugins=grpc:. proto/*.proto
    GOOS=linux GOARCH=amd64 go build
    docker build -t anvartech/mailer-service .
