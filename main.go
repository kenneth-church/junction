package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"net/mail"
	"os/exec"
	"strings"
	"text/template"

	"github.com/mhale/smtpd"
	"github.com/rs/zerolog/log"
)

func mailHandler(remoteAddr net.Addr, from string, to []string, data []byte) error {
	log.Info().Msg("Email Received")
	senderIP, _, err := net.SplitHostPort(remoteAddr.String())
	if err != nil {
		log.Error().Msg(fmt.Sprintf("Error getting IP: %s", err))
	}

	msg, err := mail.ReadMessage(bytes.NewReader(data))
	if err != nil {
		log.Error().Msg(fmt.Sprintf("Error getting email contents: %s", err))
	}

	index, matchedJunction := getJunction(to, from, senderIP)

	log.Print("To:", to)
	log.Print("From:", from)
	log.Print("IP:", senderIP)

	title, body := buildMessage(to, from, msg, senderIP, matchedJunction)

	var junctionID string
	if matchedJunction.Name != "" {
		junctionID = matchedJunction.Name
	} else {
		junctionID = fmt.Sprint(index)
	}
	log.Info().Msg(fmt.Sprintf("Sending notification to Junction - %s", junctionID))
	sendNotification(title, body, matchedJunction.Apprise)
	return nil
}

/*
getJunction checks the provided config for a matching Junction

Parameters:

	emailTo   - The To field of the received email
	emailFrom - The From field of the received email
	senderIP  - The IP address that the email was sent from

Returns:

	index    - The index of the matched Junction
	junction - The matched Junction
*/
func getJunction(emailTo []string, emailFrom string, senderIP string) (index int, junction Junction) {
	for index, junction := range junctions {
		// Determine if the To block matches
		toMatchFn := func() bool {
			// If there is no To block, match by default
			if len(junction.To.Emails) == 0 {
				return true
			}

			// Check each To value for a match
			matches := make([]bool, len(junction.To.Emails))
			for toIndex, junctionTo := range junction.To.Emails {
				for _, to := range emailTo {
					if junctionTo == strings.Trim(to, "<>") {
						matches[toIndex] = true
					}
				}
			}

			// Now check if the overall block matches
			allMatch := true
			anyMatch := false
			for _, match := range matches {
				if !match {
					allMatch = false
				}
			}

			if junction.To.RequireAll {
				return allMatch
			} else {
				return anyMatch
			}
		}

		fromMatchFn := func() bool {
			// If there is no email, match by default
			// If not, check it
			emailMatch := true
			if junction.From.Email != "" {
				emailMatch = (strings.Trim(emailFrom, "<>") == junction.From.Email)
			}

			// If the IP is empty, automatically match
			// If not, check it
			ipMatch := true
			if junction.From.IP != "" {
				ipMatch = senderIP == junction.From.IP
			}

			// Return result
			return emailMatch && ipMatch
		}

		toMatch := toMatchFn()
		fromMatch := fromMatchFn()

		if toMatch && fromMatch {
			return index, junction
		}
	}

	return -1, Junction{}
}

/*
buildMessage prepares the title and body of the notification

Parameters:

	emailTo   - The To field of the received email
	emailFrom - The From field of the received email
	msg       - The received email
	senderIP  - The IP address that the email was sent from
	junction  - The Junction to send to

Returns:

	title - The notification title
	body  - The notification body
*/
func buildMessage(emailTo []string, emailFrom string, msg *mail.Message, senderIP string, junction Junction) (title string, body string) {
	// Prepare the To addresses
	var toAddresses string
	for index, address := range emailTo {
		toAddresses += strings.Trim(address, "<>")
		if index != (len(emailTo) - 1) {
			toAddresses += ", "
		}
	}

	// Prepare the From address
	fromAddress := strings.Trim(emailFrom, "<>")

	// Prepare the Body
	builder := &strings.Builder{}
	io.Copy(builder, msg.Body)
	emailBody := builder.String()

	// Prepare the data used by the Template
	templateData := struct {
		Subject string
		Body    string
		To      string
		From    string
		Date    string
	}{
		Subject: msg.Header.Get("Subject"),
		Body:    emailBody,
		To:      toAddresses,
		From:    fromAddress,
		Date:    msg.Header.Get("Date"),
	}

	// If the Junction provides a Title Template, parse it
	// Else, use the Email Subject
	if junction.Title != "" {
		builder := &strings.Builder{}
		template, err := template.New("title").Parse(junction.Title)
		if err != nil {
			log.Error().Msg(fmt.Sprintf("Error building Title: %s", err))
		}

		template.Execute(builder, templateData)
		title = builder.String()
	} else {
		title = msg.Header.Get("Subject")
	}

	// If the Junction provides a Body Template, parse it
	// Else use the Email Body
	if junction.Body != "" {
		builder := &strings.Builder{}
		template, err := template.New("body").Parse(junction.Body)
		if err != nil {
			log.Error().Msg(fmt.Sprintf("Error building Body: %s", err))
		}

		template.Execute(builder, templateData)
		body = builder.String()
	} else {
		body = emailBody
	}

	return
}

/*
sendNotification sends the notification

Parameters:

	title      - The Title string
	body       - The Body string
	appriseURL - The Apprise URL to send to
*/
func sendNotification(title string, body string, appriseURL string) {
	apprise := exec.Command(apprisePath)
	apprise.Args = append(apprise.Args, "-vv", "-t", title, "-b", body, appriseURL)
	result, err := apprise.Output()
	if err != nil {
		log.Error().Msg(fmt.Sprintf("Error with Apprise: %s", err))
	}

	log.Print(string(result))
}

func main() {
	getConf()

	log.Info().Msg(fmt.Sprintf("Listening on port %s", port))
	smtpd.ListenAndServe(fmt.Sprintf(":%s", port), mailHandler, "", "")
}
