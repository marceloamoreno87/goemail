package sendgrid

import (
	"log"
	"os"

	"github.com/marceloamoreno87/gomail/consumer/pkg/email"
	sendgrid "github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func Send(mailmessage *email.MailMessage) {
	m := setMail(mailmessage.GetFrom(), mailmessage.GetTo(), mailmessage.GetCc(), mailmessage.GetSubject(), mailmessage.GetBody())
	request := sendgrid.GetRequest(os.Getenv("SENDGRID_API_KEY"), "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	request.Body = mail.GetRequestBody(m)
	_, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	}
}

func setMail(from string, to []string, cc []string, subject string, body string) (m *mail.SGMailV3) {
	m = mail.NewV3Mail()
	mail_from := mail.NewEmail(from, from)
	content := mail.NewContent("text/html", body)
	m.SetFrom(mail_from)
	m.AddContent(content)
	personalization := mail.NewPersonalization()
	tos := getTos(to)
	ccs := getCcs(cc)
	personalization.AddTos(tos...)
	personalization.AddCCs(ccs...)
	personalization.Subject = subject
	m.AddPersonalizations(personalization)
	return
}

func getTos(to []string) (tos []*mail.Email) {
	for _, mail_to := range to {
		m := mail.NewEmail(mail_to, mail_to)
		tos = append(tos, m)
	}
	return
}

func getCcs(cc []string) (ccs []*mail.Email) {
	for _, mail_to := range cc {
		m := mail.NewEmail(mail_to, mail_to)
		ccs = append(ccs, m)
	}
	return
}
