package email

import (
	"encoding/json"
	"os"

	"github.com/marceloamoreno87/gomail/pkg/gomail"
	"github.com/marceloamoreno87/gomail/pkg/sendgrid"
	"github.com/marceloamoreno87/gomail/pkg/ses"
)

type MailMessage struct {
	To      []string `json:"to"`
	Cc      []string `json:"cc"`
	Subject string   `json:"subject"`
	From    string   `json:"from"`
	Body    string   `json:"body"`
}

func setMailMessage(message_body []byte) *MailMessage {
	mailmessage := NewMailMessage()
	json.Unmarshal(message_body, &mailmessage)
	mailmessage.SetCc(mailmessage.Cc)
	mailmessage.SetFrom(mailmessage.From)
	mailmessage.SetSubject(mailmessage.Subject)
	mailmessage.SetBody(mailmessage.Body)
	mailmessage.SetTo(mailmessage.To)
	return mailmessage
}

func Send(message_body []byte) {
	mailmessage := setMailMessage(message_body)
	switch os.Getenv("MAIL_DRIVER") {
	case "gomail":
		gomail.Send(mailmessage.GetFrom(), mailmessage.GetTo(), mailmessage.GetCc(), mailmessage.GetSubject(), mailmessage.GetBody())
	case "sendgrid":
		sendgrid.Send(mailmessage.GetFrom(), mailmessage.GetTo(), mailmessage.GetCc(), mailmessage.GetSubject(), mailmessage.GetBody())
	case "ses":
		ses.Send(mailmessage.GetFrom(), mailmessage.GetTo(), mailmessage.GetCc(), mailmessage.GetSubject(), mailmessage.GetBody())
	default:
		gomail.Send(mailmessage.GetFrom(), mailmessage.GetTo(), mailmessage.GetCc(), mailmessage.GetSubject(), mailmessage.GetBody())
	}
}

func NewMailMessage() *MailMessage {
	return &MailMessage{}
}

func (mailmessage *MailMessage) GetTo() []string {
	return mailmessage.To
}

func (mailmessage *MailMessage) GetCc() []string {
	return mailmessage.Cc
}

func (mailmessage *MailMessage) GetSubject() string {
	return mailmessage.Subject
}

func (mailmessage *MailMessage) GetFrom() string {
	return mailmessage.From
}

func (mailmessage *MailMessage) GetBody() string {
	return mailmessage.Body
}

func (mailmessage *MailMessage) SetTo(To []string) *MailMessage {
	mailmessage.To = To
	return mailmessage
}

func (mailmessage *MailMessage) SetCc(Cc []string) *MailMessage {
	mailmessage.Cc = Cc
	return mailmessage
}

func (mailmessage *MailMessage) SetSubject(Subject string) *MailMessage {
	mailmessage.Subject = Subject
	return mailmessage
}

func (mailmessage *MailMessage) SetFrom(From string) *MailMessage {
	mailmessage.From = From
	return mailmessage
}

func (mailmessage *MailMessage) SetBody(Body string) *MailMessage {
	mailmessage.Body = Body
	return mailmessage
}
