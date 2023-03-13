package controllers

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/marceloamoreno87/gomail/pkg/email"
	"github.com/marceloamoreno87/gomail/pkg/rabbitmq"
)

// GetBenfitByDoc             godoc
// @Summary      Send email by HTML template
// @Description  Send email by HTML template
// @Tags         Send email
// @Param        doc  query      string  true  "Type the cpf of the user who wants to get the benefits"
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
