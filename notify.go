package main

import (
	"fmt"
	"os/exec"
	"strings"
	"text/template"

	"github.com/rs/zerolog/log"
)

/*
buildMessage prepares the title and body of the notification

Parameters:

	email    - Data from the received email
	junction - The Junction to send to

Returns:

	title - The notification title
	body  - The notification body
*/
func buildMessage(email EmailData, junction Junction) (title string, body string, url string) {
	// Prepare the data used by the Template
	templateData := struct {
		Subject string   // The received email's subject line
		Body    string   // The received email's body
		To      string   // The received email's to field preformatted
		From    string   // The received email's from field
		Date    string   // The date the received email was sent
		IP      string   // The IP of the machine that sent the received email
		RawTo   []string // The raw slice of the email's to field
	}{
		Subject: email.Subject,
		Body:    email.Body,
		To:      strings.Join(email.To, ","),
		From:    email.From,
		Date:    email.Date,
		IP:      email.IP,
		RawTo:   email.To,
	}

	// If the Junction provides a Title Template, parse it
	// Else, use the Email Subject
	if junction.Title != "" {
		builder := &strings.Builder{}
		template, err := template.New("title").Parse(junction.Title)
		if err != nil {
			log.Error().Err(err).Msg("Can't parse the title")
		}

		template.Execute(builder, templateData)
		title = builder.String()
	} else {
		title = email.Subject
	}

	// If the Junction provides a Body Template, parse it
	// Else use the Email Body
	if junction.Body != "" {
		builder := &strings.Builder{}
		template, err := template.New("body").Parse(junction.Body)
		if err != nil {
			log.Error().Err(err).Msg("Can't parse the body")
		}

		template.Execute(builder, templateData)
		body = builder.String()
	} else {
		body = email.Body
	}

	// Build the URL template
	builder := &strings.Builder{}
	template, err := template.New("url").Parse(junction.Apprise)
	if err != nil {
		log.Error().Err(err).Msg("Can't parse the url")
	}

	template.Execute(builder, templateData)
	url = builder.String()

	return
}

/*
sendNotification sends the notification using the Apprise CLI

Parameters:

	title      - The Title string
	body       - The Body string
	appriseURL - The Apprise URL to send to
*/
func sendNotification(title string, body string, appriseURL string) {
	apprise := exec.Command(apprisePath)
	apprise.Args = append(apprise.Args, "-vv", "-t", title, "-b", body, fmt.Sprintf("%s?overflow=split", appriseURL))
	result, err := apprise.Output()
	if err != nil {
		log.Error().Err(err).Msg("Apprise returned an error")
	}

	log.Print(string(result))
}
