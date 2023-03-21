package ses

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/marceloamoreno87/gomail/consumer/pkg/email"
)

func Send(mailmessage *email.MailMessage) {
	email_template := generateSESTemplate(mailmessage.GetFrom(), mailmessage.GetTo(), mailmessage.GetCc(), mailmessage.GetSubject(), mailmessage.GetBody())
	sess, err := getSessionSES()
	service := ses.New(sess)
	_, err = service.SendEmail(email_template)

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			log.Fatal(aerr.Error())
		} else {
			log.Fatal(err)
		}
	}

}

func getSessionSES() (sess *session.Session, error error) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(os.Getenv("SES_REGION")),
		Credentials: credentials.NewStaticCredentials(os.Getenv("SES_ACCESS_KEY_ID"), os.Getenv("SES_SECRET_ACCESS_KEY"), ""),
	})

	if err != nil {
		log.Fatal(err)
	}

	return
}

func generateSESTemplate(from string, to []string, cc []string, subject string, body string) (template *ses.SendEmailInput) {
	template = &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: aws.StringSlice(cc),
			ToAddresses: aws.StringSlice(to),
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String("utf-8"),
					Data:    aws.String(body),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String("utf-8"),
				Data:    aws.String(subject),
			},
		},
		Source: aws.String(from),
	}
	return
}
