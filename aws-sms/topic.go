package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

func main() {

	fmt.Println("creating session")
	sess := session.Must(session.NewSession())
	fmt.Println("session created")

	svc := sns.New(sess)
	fmt.Println("service created")

	out, err := svc.Subscribe(&sns.SubscribeInput{
		Endpoint: aws.String(""),
		Protocol: aws.String("email"),
		TopicArn: aws.String(""),
	})
	if err != nil {
		fmt.Println("---", err.Error())
		return
	}
	fmt.Println("Subscription Arn:", out.SubscriptionArn)

	params := &sns.PublishInput{
		Message:     aws.String("This is a test SMS!"),
		PhoneNumber: aws.String("+86 18980501737"),
	}
	resp, err := svc.Publish(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
