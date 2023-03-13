package controllers

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/marceloamoreno87/gomail/pkg/email"
	"github.com/marceloamoreno87/gomail/pkg/rabbitmq"
)

// SendEmail             godoc
// @Summary      Send email by HTML template
// @Description  Send email by HTML template
// @Tags         Send email
// @Param        doc  body  email.MailMessage  true  "query params"
// @Produce      json
// @Router       /send-email [post]
func SendEmail(c *gin.Context) {

	message := email.NewMailMessage()
	c.ShouldBind(&message)
	json_data, err := json.Marshal(message)
	if err != nil {
		return
	}

	rabbitmq.Publish(json_data)
	c.JSON(200, "Email enviado!")
}
