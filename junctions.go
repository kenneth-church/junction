package main

import (
	"github.com/rs/zerolog/log"
)

type Junction struct {
	Name    string   `yaml:"name,omitempty"`
	Apprise string   `yaml:"apprise"`
	To      JuncTo   `yaml:"to,omitempty"`
	From    JuncFrom `yaml:"from,omitempty"`
	Title   string   `yaml:"title,omitempty"`
	Body    string   `yaml:"body,omitempty"`
}

type JuncTo struct {
	Emails     []string `yaml:"emails"`
	RequireAll bool     `yaml:"require-all,omitempty"`
}

type JuncFrom struct {
	Email string `yaml:"email,omitempty"`
	IP    string `yaml:"ip,omitempty"`
}

/*
selectJunction determines which Junction should be used

Parameters:

	to   - The To field of the received email
	from - The From field of the received email
	ip        - The IP address that the email was sent from

Returns:

	int       - The index of the selected Junction
*/
func selectJunction(to []string, from string, ip string) int {
	for index, junction := range junctions {
		log.Debug().Int("junction index", index).Msg("Checking")

		// Check if the to and from blocks provided satisify the junction conditions
		toMatch := checkTo(junction.To, to)
		fromMatch := checkFrom(junction.From, from, ip)

		log.Debug().Bool("to", toMatch).Bool("from", fromMatch).Msg("Results")

		if toMatch && fromMatch {
			return index
		}
	}

	return -1
}

/*
checkTo determines if the provided junction's 'To' field matches the received email

Parameters:

	juncTo  - The 'To' block of the junction to compare with
	email   - The email addresses to check

Returns:

	bool    - Whether or not the conditions match
*/
func checkTo(juncTo JuncTo, email []string) bool {
	log.Debug().Msg("   Checking 'to email' conditions")
	// If there is no To block, match by default
	if len(juncTo.Emails) == 0 {
		log.Debug().Msg("     No condition provided, matches by default")
		return true
	}

	// Check each To value for a match
	matches := make([]bool, len(juncTo.Emails))
	for index, junctionEmail := range juncTo.Emails {
		for _, toEmail := range email {
			matched := junctionEmail == toEmail
			log.Debug().Str("provided email", junctionEmail).Str("received email", toEmail).Bool("matches", matched).Msg("     ")
			if matched {
				if !juncTo.RequireAll || len(junctionEmail) == 1 {
					log.Debug().Msg("     Only one match required, 'to' matches")
					return true
				}
				matches[index] = true
			}
		}
	}

	log.Debug().Msg("     Making sure required emails are present")
	for index, match := range matches {
		if !match {
			log.Debug().Str("email", juncTo.Emails[index]).Msg("     Required emails not present, 'to' doesn't match")
			return false
		}
	}

	log.Debug().Msg("     Required emails present, 'to' matches")
	return true
}

/*
checkFrom determines if the provided junction's 'From' field matches the received email

Parameters:

	juncFrom - The 'From' block of the junction to compare with
	email    - The email address to check
	ip       - The IP addresses to check

Returns:

	bool     - Whether or not the conditions match
*/
func checkFrom(juncFrom JuncFrom, email string, ip string) bool {
	var emailMatched bool
	var ipMatched bool

	log.Debug().Msg("   Checking 'from email' condition")
	// If there is no email, match by default
	if juncFrom.Email == "" {
		log.Debug().Msg("     No condition provided, matches by default")
		emailMatched = true
	} else {
		emailMatched = email == juncFrom.Email
		log.Debug().Str("provided email", juncFrom.Email).Str("received email", email).Bool("matches", emailMatched).Msg("     ")
	}

	log.Debug().Msg("   Checking 'from ip' condition")
	// If the IP is empty, automatically match
	if juncFrom.IP == "" {
		log.Debug().Msg("     No condition provided, matching by default")
		ipMatched = true
	} else {
		ipMatched = ip == juncFrom.IP
		log.Debug().Str("provided ip", juncFrom.IP).Str("received ip", ip).Bool("matches", ipMatched).Msg("     ")
	}

	return emailMatched && ipMatched
}
