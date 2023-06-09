// Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Marcelo Moreno",
            "email": "marceloamoreno87@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/send-email": {
            "post": {
                "description": "Send email by HTML template",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Send email"
                ],
                "summary": "Send email by HTML template",
                "parameters": [
                    {
                        "description": "query params",
                        "name": "doc",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/email.MailMessage"
                        }
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "email.MailMessage": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "string",
                    "example": "\u003ch1\u003eHello, world!\u003c/h1\u003e"
                },
                "cc": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "test@test.com",
                        " test2@test2.com"
                    ]
                },
                "from": {
                    "type": "string",
                    "example": "marceloamoreno87@gmail.com"
                },
                "subject": {
                    "type": "string",
                    "example": "testing"
                },
                "to": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "test@test.com",
                        " test2@test2.com"
                    ]
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "GOMAIL",
	Description:      "Serviço de API para enviar email.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
