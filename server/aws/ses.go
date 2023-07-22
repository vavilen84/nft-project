package main

import (
	"fmt"
	"github.com/vavilen84/nft-project/constants"
	"github.com/vavilen84/nft-project/helpers"
	"os"

	//go get -u github.com/aws/aws-sdk-go
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

func SendResetPasswordEmail(recipient, token string) error {
	// TODO: replace with real domain
	link := fmt.Sprintf(
		constants.ResetPasswordHtmlBodyFormat,
		os.Getenv("DOMAIN")+"/reset-password&token="+token,
	)
	return sendEmail(
		recipient,
		constants.NoReplySenderEmail,
		constants.ResetPasswordSubject,
		link,
	)
}

func sendEmail(recipient, sender, subject, htmlBody string) error {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION"))},
	)
	svc := ses.New(sess)
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{},
			ToAddresses: []*string{
				aws.String(recipient),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(constants.EmailCharSet),
					Data:    aws.String(htmlBody),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(constants.EmailCharSet),
				Data:    aws.String(subject),
			},
		},
		Source: aws.String(sender),
	}
	_, err = svc.SendEmail(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ses.ErrCodeMessageRejected:
				err := errors.New(fmt.Sprintf("%s: %s", ses.ErrCodeMessageRejected, aerr.Error()))
				helpers.LogError(err)
				return err
			case ses.ErrCodeMailFromDomainNotVerifiedException:
				err := errors.New(fmt.Sprintf("%s: %s", ses.ErrCodeMailFromDomainNotVerifiedException, aerr.Error()))
				helpers.LogError(err)
				return err
			case ses.ErrCodeConfigurationSetDoesNotExistException:
				err := errors.New(fmt.Sprintf("%s: %s", ses.ErrCodeConfigurationSetDoesNotExistException, aerr.Error()))
				helpers.LogError(err)
				return err
			default:
				err := errors.New(aerr.Error())
				helpers.LogError(err)
				return err
			}
		} else {
			err := errors.New(aerr.Error())
			helpers.LogError(err)
			return err
		}
	}
	return nil
}
