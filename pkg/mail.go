package pkg

import (
	"net/http"

	// (1)
	"appengine"
	"appengine/mail"
)

// Subject
const subject = "Test Send"

// Mail body
const mailBody = `
Thank you!

This is demo of Sending Email.
`

func init() {
	http.HandleFunc("/sendEmail", sendEmail)
}

func sendEmail(w http.ResponseWriter, r *http.Request) {
	// Context
	ctx := appengine.NewContext(r)
	// Specify Sender Email
	from := "sender@gmail.com"
	// Specify Destination Emails
	to := []string{"to1@gmail.com", "to2@gmail.com"}
	// Create Message
	msg := &mail.Message{
		Sender:  from,
		To:      to,
		Subject: subject,
		Body:    mailBody,
	}
	// Send Email
	if err := mail.Send(ctx, msg); err != nil {
		// Error Handler
	}
}
