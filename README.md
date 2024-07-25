## Description

Ce projet est une application Go qui utilise l'API Mailjet pour envoyer un email. L'application envoie un email de demande de réinitialisation de mot de passe en utilisant un template Mailjet.

## Prérequis

- Go 1.14 ou supérieur
- Un compte Mailjet avec une clé publique et une clé privée API
- Un template Mailjet avec votre ID 

## Installation

1. Clonez ce dépôt :

```bash
git clone <URL_DU_DEPOT>
cd <NOM_DU_DEPOT>
```

```sh
go get github.com/mailjet/mailjet-apiv3-go
```

### Configuration 

Modifiez le code pour inclure vos propres informations d'API Mailjet et les adresses email :

```go
mailjetClient := mailjet.NewMailjetClient("API_PUBLIC_MAILJET", "API_PRIVATE_MAILJET")
```

Remplacez "API_PUBLIC_MAILJET" et "API_PRIVATE_MAILJET" par vos propres clés API Mailjet.

Ajustez les informations de l'expéditeur et du destinataire :

```go
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
```

Remplacez "sender_email", "name", "destination_email" et "name" par les informations adéquates.


### Utilisation

Pour exécuter l'application, utilisez la commande suivante :

```sh
go run .
```

L'application enverra l'email en utilisant l'API Mailjet et affichera les données de réponse.

### Dépannage

Si vous rencontrez des problèmes, assurez-vous que :

- Vos clés API Mailjet sont correctes.
- L'ID du template existe dans votre compte Mailjet.
- Les adresses email sont valides.

En cas d'erreur, un message sera affiché dans la console avec des détails supplémentaires.

