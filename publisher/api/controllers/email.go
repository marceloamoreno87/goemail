package controllers

import (
	"encoding/json"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/marceloamoreno87/gomail/publisher/pkg/email"
	"github.com/marceloamoreno87/gomail/publisher/pkg/rabbitmq"
)

type EmailController struct{}

// SendEmail             godoc
// @Summary      Send email by HTML template
// @Description  Send email by HTML template
// @Tags         Send email
// @Param        doc  body  email.MailMessage  true  "query params"
// @Produce      json
// @Router       /mail [post]
func (ctrl EmailController) Store(c *gin.Context) {

	message := email.NewMailMessage()
	c.ShouldBind(&message)
	err := message.ValidateEmailMessage()
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	SaveAttachments(c, message.GetAttachments())

	json_data, err := json.Marshal(message)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	rabbitmq.Publish(json_data)
	c.JSON(200, "Email enviado!")
}

func SaveAttachments(c *gin.Context, attachemnts []*multipart.FileHeader) {
	for _, file := range attachemnts {
		file.Filename = uuid.NewString() + "-" + file.Filename
		filename := "/tmp/attachments/" + file.Filename
		if err := c.SaveUploadedFile(file, filename); err != nil {
			c.String(http.StatusBadRequest, "upload file err: %s", err.Error())
			return
		}
	}
}
