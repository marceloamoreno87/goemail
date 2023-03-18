package email

import (
	"encoding/json"
	"errors"
	"mime/multipart"
	"os"

	"github.com/marceloamoreno87/gomail/pkg/gomail"
	"github.com/marceloamoreno87/gomail/pkg/sendgrid"
	"github.com/marceloamoreno87/gomail/pkg/ses"
)

type MailMessage struct {
	To          []string                `json:"to" form:"to" example:"test@test.com, test2@test2.com" binding:"required"`
	Cc          []string                `json:"cc" form:"cc" example:"test@test.com, test2@test2.com"`
	Subject     string                  `json:"subject" form:"subject" example:"testing" binding:"required"`
	From        string                  `json:"from" form:"from" example:"marceloamoreno87@gmail.com" binding:"required"`
	Body        string                  `json:"body" form:"body" example:"<h1>Hello, world!</h1>" binding:"required"`
	Attachments []*multipart.FileHeader `json:"attachment,omitempty" form:"attachment"`
}

func setMailMessage(message_body []byte) *MailMessage {
	mailmessage := NewMailMessage()
	json.Unmarshal(message_body, &mailmessage)
	return mailmessage
}

func (mailmessage *MailMessage) ValidateEmailMessage() error {
	if mailmessage.Body == "" {
		return errors.New("Campo body está em branco!")
	}
	if mailmessage.From == "" {
		return errors.New("Campo from está em branco!")
	}
	if mailmessage.To == nil {
		return errors.New("Campo to está em branco!")
	}
	if mailmessage.Cc == nil {
		return errors.New("Campo cc está em branco!")
	}
	if mailmessage.Subject == "" {
		return errors.New("Campo subject está em branco!")
	}
	return nil
}

func Send(message_body []byte) {
	mailmessage := setMailMessage(message_body)
	switch os.Getenv("MAIL_DRIVER") {
	case "gomail":
		gomail.Send(mailmessage.GetFrom(), mailmessage.GetTo(), mailmessage.GetCc(), mailmessage.GetSubject(), mailmessage.GetBody(), mailmessage.GetAttachments())
	case "sendgrid":
		sendgrid.Send(mailmessage.GetFrom(), mailmessage.GetTo(), mailmessage.GetCc(), mailmessage.GetSubject(), mailmessage.GetBody(), mailmessage.GetAttachments())
	case "ses":
		ses.Send(mailmessage.GetFrom(), mailmessage.GetTo(), mailmessage.GetCc(), mailmessage.GetSubject(), mailmessage.GetBody(), mailmessage.GetAttachments())
	default:
		gomail.Send(mailmessage.GetFrom(), mailmessage.GetTo(), mailmessage.GetCc(), mailmessage.GetSubject(), mailmessage.GetBody(), mailmessage.GetAttachments())
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

func (mailmessage *MailMessage) GetAttachments() []*multipart.FileHeader {
	return mailmessage.Attachments
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

func (mailmessage *MailMessage) SetAttachments(Attachments []*multipart.FileHeader) *MailMessage {
	mailmessage.Attachments = Attachments
	return mailmessage
}
