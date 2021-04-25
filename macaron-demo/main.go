package main

import (
	"fmt"

	"github.com/go-macaron/binding"
	"gopkg.in/macaron.v1"
)

type ContactForm struct {
	Name           string `form:"name222" json:"name222" binding:"Required"`
	Email          string `form:"email"`
	Message        string `form:"message" binding:"Required"`
	MailingAddress string `form:"mailing_address"`
}

func main() {
	m := macaron.Classic()
	m.Get("/", func() string {
		return "Hello world!"
	})
	m.Post("/contact/submit", binding.Bind(ContactForm{}), func(contact ContactForm) string {
		return fmt.Sprintf("Name: %s\nEmail: %s\nMessage: %s\nMailing Address: %v",
			contact.Name, contact.Email, contact.Message, contact.MailingAddress)
	})
	m.Run()
}
