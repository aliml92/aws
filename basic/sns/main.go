package main

import (
	"log"
	"time"

	"github.com/aliml92/aws/basic/sqs/internal/message"
	"github.com/aliml92/aws/basic/sqs/internal/pkg/cloud/aws"
)

func main() {
	ses, err := aws.New(aws.Config{
		Address: "http://localhost:4566",
		Region:  "eu-central-1",
		Profile: "localstack",
		ID: "test",
		Secret: "test",
	})
	if err != nil {
		log.Fatalln(err)
	}

	message.Message(aws.NewSQS(ses, time.Second*5))
}