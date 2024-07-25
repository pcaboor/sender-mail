package main

import (
	"fmt"
	"log"

	mailjet "github.com/mailjet/mailjet-apiv3-go"
)

func main() {
	mailjetClient := mailjet.NewMailjetClient("API_PUBLIC_MAILJET", "API_PRIVATE_MAILJET")
	messagesInfo := []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Email: "sender_email",
				Name:  "name",
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: "destination_email",
					Name:  "name",
				},
			},
			TemplateID:       6165565,
			TemplateLanguage: true,
			Subject:          "Demande de réinitialisation du mot de passe.",
		},
	}
	messages := mailjet.MessagesV31{Info: messagesInfo}
	res, err := mailjetClient.SendMailV31(&messages)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data: %+v\n", res)
}
