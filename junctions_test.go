package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/rs/zerolog"
)

var testJuncs = []Junction{
	{
		Apprise: "json://localhost",
	},
	{
		Apprise: "json://localhost",
		To: JuncTo{
			Emails: []string{"testto@test.com"},
		},
	},
	{
		Apprise: "json://localhost",
		To: JuncTo{
			Emails: []string{"testto@test.com", "testto2@test.com"},
		},
	},
	{
		Apprise: "json://localhost",
		To: JuncTo{
			Emails:     []string{"testto@test.com", "testto2@test.com"},
			RequireAll: true,
		},
	},
	{
		Apprise: "json://localhost",
		From: JuncFrom{
			Email: "testfrom@test.com",
		},
	},
	{
		Apprise: "json://localhost",
		From: JuncFrom{
			IP: "1.1.1.1",
		},
	},
	{
		Apprise: "json://localhost",
		From: JuncFrom{
			Email: "testfrom@test.com",
			IP:    "1.1.1.1",
		},
	},
	{
		Apprise: "json://localhost",
		To: JuncTo{
			Emails: []string{"testto@test.com"},
		},
		From: JuncFrom{
			Email: "testfrom@test.com",
		},
	},
	{
		Apprise: "json://localhost",
		To: JuncTo{
			Emails: []string{"testto@test.com", "testto2@test.com"},
		},
		From: JuncFrom{
			Email: "testfrom@test.com",
		},
	},
	{
		Apprise: "json://localhost",
		To: JuncTo{
			Emails:     []string{"testto@test.com", "testto2@test.com"},
			RequireAll: true,
		},
		From: JuncFrom{
			Email: "testfrom@test.com",
		},
	},
	{
		Apprise: "json://localhost",
		To: JuncTo{
			Emails: []string{"testto@test.com"},
		},
		From: JuncFrom{
			IP: "1.1.1.1",
		},
	},
	{
		Apprise: "json://localhost",
		To: JuncTo{
			Emails: []string{"testto@test.com", "testto2@test.com"},
		},
		From: JuncFrom{
			IP: "1.1.1.1",
		},
	},
	{
		Apprise: "json://localhost",
		To: JuncTo{
			Emails:     []string{"testto@test.com", "testto2@test.com"},
			RequireAll: true,
		},
		From: JuncFrom{
			IP: "1.1.1.1",
		},
	},
	{
		Apprise: "json://localhost",
		To: JuncTo{
			Emails: []string{"testto@test.com"},
		},
		From: JuncFrom{
			Email: "testfrom@test.com",
			IP:    "1.1.1.1",
		},
	},
	{
		Apprise: "json://localhost",
		To: JuncTo{
			Emails: []string{"testto@test.com", "testto2@test.com"},
		},
		From: JuncFrom{
			Email: "testfrom@test.com",
			IP:    "1.1.1.1",
		},
	},
	{
		Apprise: "json://localhost",
		To: JuncTo{
			Emails:     []string{"testto@test.com", "testto2@test.com"},
			RequireAll: true,
		},
		From: JuncFrom{
			Email: "testfrom@test.com",
			IP:    "1.1.1.1",
		},
	},
}

var testEmails = []EmailData{
	{
		To:   []string{"testto@test.com"},
		From: "testfrom@test.com",
		IP:   "1.1.1.1",
	},
	{
		To:   []string{"testto@test.com"},
		From: "testfromwrong@test.com",
		IP:   "1.1.1.1",
	},
	{
		To:   []string{"testto@test.com"},
		From: "testfrom@test.com",
		IP:   "8.8.8.8",
	},
	{
		To:   []string{"testtowrong@test.com"},
		From: "testfrom@test.com",
		IP:   "1.1.1.1",
	},
	{
		To:   []string{"testtowrong@test.com"},
		From: "testfromwrong@test.com",
		IP:   "1.1.1.1",
	},
	{
		To:   []string{"testtowrong@test.com"},
		From: "testfrom@test.com",
		IP:   "8.8.8.8",
	},
	{
		To:   []string{"testtowrong@test.com"},
		From: "testfromwrong@test.com",
		IP:   "8.8.8.8",
	},
	{
		To:   []string{"testto@test.com", "testto2@test.com"},
		From: "testfrom@test.com",
		IP:   "1.1.1.1",
	},
	{
		To:   []string{"testtowrong@test.com", "testto2@test.com"},
		From: "testfrom@test.com",
		IP:   "1.1.1.1",
	},
	{
		To:   []string{"testto@test.com", "testto2wrong@test.com"},
		From: "testfrom@test.com",
		IP:   "1.1.1.1",
	},
	{
		To:   []string{"testto@test.com", "testto2@test.com"},
		From: "testfromwrong@test.com",
		IP:   "1.1.1.1",
	},
	{
		To:   []string{"testtowrong@test.com", "testto2wrong@test.com"},
		From: "testfromwrong@test.com",
		IP:   "1.1.1.1",
	},
	{
		To:   []string{"testtowrong@test.com", "testto2wrong@test.com"},
		From: "testfromwrong@test.com",
		IP:   "8.8.8.8",
	},
	{
		To:   []string{"testto@test.com", "testto2@test.com"},
		From: "testfrom@test.com",
		IP:   "8.8.8.8",
	},
}

func TestMain(m *testing.M) {
	zerolog.SetGlobalLevel(zerolog.Disabled)
	os.Exit(m.Run())
}

func TestCheckTo(t *testing.T) {
	var tests = []struct {
		junc   Junction
		email  EmailData
		result bool
	}{
		{testJuncs[0], testEmails[0], true},
		{testJuncs[0], testEmails[1], true},
		{testJuncs[0], testEmails[2], true},
		{testJuncs[0], testEmails[3], true},
		{testJuncs[0], testEmails[4], true},
		{testJuncs[0], testEmails[5], true},
		{testJuncs[0], testEmails[6], true},
		{testJuncs[0], testEmails[7], true},
		{testJuncs[0], testEmails[8], true},
		{testJuncs[0], testEmails[9], true},
		{testJuncs[0], testEmails[10], true},
		{testJuncs[0], testEmails[11], true},
		{testJuncs[0], testEmails[12], true},
		{testJuncs[0], testEmails[13], true},

		{testJuncs[1], testEmails[0], true},
		{testJuncs[1], testEmails[1], true},
		{testJuncs[1], testEmails[2], true},
		{testJuncs[1], testEmails[3], false},
		{testJuncs[1], testEmails[4], false},
		{testJuncs[1], testEmails[5], false},
		{testJuncs[1], testEmails[6], false},
		{testJuncs[1], testEmails[7], true},
		{testJuncs[1], testEmails[8], false},
		{testJuncs[1], testEmails[9], true},
		{testJuncs[1], testEmails[10], true},
		{testJuncs[1], testEmails[11], false},
		{testJuncs[1], testEmails[12], false},
		{testJuncs[1], testEmails[13], true},

		{testJuncs[2], testEmails[0], true},
		{testJuncs[2], testEmails[1], true},
		{testJuncs[2], testEmails[2], true},
		{testJuncs[2], testEmails[3], false},
		{testJuncs[2], testEmails[4], false},
		{testJuncs[2], testEmails[5], false},
		{testJuncs[2], testEmails[6], false},
		{testJuncs[2], testEmails[7], true},
		{testJuncs[2], testEmails[8], true},
		{testJuncs[2], testEmails[9], true},
		{testJuncs[2], testEmails[10], true},
		{testJuncs[2], testEmails[11], false},
		{testJuncs[2], testEmails[12], false},
		{testJuncs[2], testEmails[13], true},

		{testJuncs[3], testEmails[0], false},
		{testJuncs[3], testEmails[1], false},
		{testJuncs[3], testEmails[2], false},
		{testJuncs[3], testEmails[3], false},
		{testJuncs[3], testEmails[4], false},
		{testJuncs[3], testEmails[5], false},
		{testJuncs[3], testEmails[6], false},
		{testJuncs[3], testEmails[7], true},
		{testJuncs[3], testEmails[8], false},
		{testJuncs[3], testEmails[9], false},
		{testJuncs[3], testEmails[10], true},
		{testJuncs[3], testEmails[11], false},
		{testJuncs[3], testEmails[12], false},
		{testJuncs[3], testEmails[13], true},

		{testJuncs[4], testEmails[0], true},
		{testJuncs[4], testEmails[1], true},
		{testJuncs[4], testEmails[2], true},
		{testJuncs[4], testEmails[3], true},
		{testJuncs[4], testEmails[4], true},
		{testJuncs[4], testEmails[5], true},
		{testJuncs[4], testEmails[6], true},
		{testJuncs[4], testEmails[7], true},
		{testJuncs[4], testEmails[8], true},
		{testJuncs[4], testEmails[9], true},
		{testJuncs[4], testEmails[10], true},
		{testJuncs[4], testEmails[11], true},
		{testJuncs[4], testEmails[12], true},
		{testJuncs[4], testEmails[13], true},

		{testJuncs[5], testEmails[0], true},
		{testJuncs[5], testEmails[1], true},
		{testJuncs[5], testEmails[2], true},
		{testJuncs[5], testEmails[3], true},
		{testJuncs[5], testEmails[4], true},
		{testJuncs[5], testEmails[5], true},
		{testJuncs[5], testEmails[6], true},
		{testJuncs[5], testEmails[7], true},
		{testJuncs[5], testEmails[8], true},
		{testJuncs[5], testEmails[9], true},
		{testJuncs[5], testEmails[10], true},
		{testJuncs[5], testEmails[11], true},
		{testJuncs[5], testEmails[12], true},
		{testJuncs[5], testEmails[13], true},

		{testJuncs[6], testEmails[0], true},
		{testJuncs[6], testEmails[1], true},
		{testJuncs[6], testEmails[2], true},
		{testJuncs[6], testEmails[3], true},
		{testJuncs[6], testEmails[4], true},
		{testJuncs[6], testEmails[5], true},
		{testJuncs[6], testEmails[6], true},
		{testJuncs[6], testEmails[7], true},
		{testJuncs[6], testEmails[8], true},
		{testJuncs[6], testEmails[9], true},
		{testJuncs[6], testEmails[10], true},
		{testJuncs[6], testEmails[11], true},
		{testJuncs[6], testEmails[12], true},
		{testJuncs[6], testEmails[13], true},

		{testJuncs[7], testEmails[0], true},
		{testJuncs[7], testEmails[1], true},
		{testJuncs[7], testEmails[2], true},
		{testJuncs[7], testEmails[3], false},
		{testJuncs[7], testEmails[4], false},
		{testJuncs[7], testEmails[5], false},
		{testJuncs[7], testEmails[6], false},
		{testJuncs[7], testEmails[7], true},
		{testJuncs[7], testEmails[8], false},
		{testJuncs[7], testEmails[9], true},
		{testJuncs[7], testEmails[10], true},
		{testJuncs[7], testEmails[11], false},
		{testJuncs[7], testEmails[12], false},
		{testJuncs[7], testEmails[13], true},

		{testJuncs[8], testEmails[0], true},
		{testJuncs[8], testEmails[1], true},
		{testJuncs[8], testEmails[2], true},
		{testJuncs[8], testEmails[3], false},
		{testJuncs[8], testEmails[4], false},
		{testJuncs[8], testEmails[5], false},
		{testJuncs[8], testEmails[6], false},
		{testJuncs[8], testEmails[7], true},
		{testJuncs[8], testEmails[8], true},
		{testJuncs[8], testEmails[9], true},
		{testJuncs[8], testEmails[10], true},
		{testJuncs[8], testEmails[11], false},
		{testJuncs[8], testEmails[12], false},
		{testJuncs[8], testEmails[13], true},

		{testJuncs[9], testEmails[0], false},
		{testJuncs[9], testEmails[1], false},
		{testJuncs[9], testEmails[2], false},
		{testJuncs[9], testEmails[3], false},
		{testJuncs[9], testEmails[4], false},
		{testJuncs[9], testEmails[5], false},
		{testJuncs[9], testEmails[6], false},
		{testJuncs[9], testEmails[7], true},
		{testJuncs[9], testEmails[8], false},
		{testJuncs[9], testEmails[9], false},
		{testJuncs[9], testEmails[10], true},
		{testJuncs[9], testEmails[11], false},
		{testJuncs[9], testEmails[12], false},
		{testJuncs[9], testEmails[13], true},

		{testJuncs[10], testEmails[0], true},
		{testJuncs[10], testEmails[1], true},
		{testJuncs[10], testEmails[2], true},
		{testJuncs[10], testEmails[3], false},
		{testJuncs[10], testEmails[4], false},
		{testJuncs[10], testEmails[5], false},
		{testJuncs[10], testEmails[6], false},
		{testJuncs[10], testEmails[7], true},
		{testJuncs[10], testEmails[8], false},
		{testJuncs[10], testEmails[9], true},
		{testJuncs[10], testEmails[10], true},
		{testJuncs[10], testEmails[11], false},
		{testJuncs[10], testEmails[12], false},
		{testJuncs[10], testEmails[13], true},

		{testJuncs[11], testEmails[0], true},
		{testJuncs[11], testEmails[1], true},
		{testJuncs[11], testEmails[2], true},
		{testJuncs[11], testEmails[3], false},
		{testJuncs[11], testEmails[4], false},
		{testJuncs[11], testEmails[5], false},
		{testJuncs[11], testEmails[6], false},
		{testJuncs[11], testEmails[7], true},
		{testJuncs[11], testEmails[8], true},
		{testJuncs[11], testEmails[9], true},
		{testJuncs[11], testEmails[10], true},
		{testJuncs[11], testEmails[11], false},
		{testJuncs[11], testEmails[12], false},
		{testJuncs[11], testEmails[13], true},

		{testJuncs[12], testEmails[0], false},
		{testJuncs[12], testEmails[1], false},
		{testJuncs[12], testEmails[2], false},
		{testJuncs[12], testEmails[3], false},
		{testJuncs[12], testEmails[4], false},
		{testJuncs[12], testEmails[5], false},
		{testJuncs[12], testEmails[6], false},
		{testJuncs[12], testEmails[7], true},
		{testJuncs[12], testEmails[8], false},
		{testJuncs[12], testEmails[9], false},
		{testJuncs[12], testEmails[10], true},
		{testJuncs[12], testEmails[11], false},
		{testJuncs[12], testEmails[12], false},
		{testJuncs[12], testEmails[13], true},

		{testJuncs[13], testEmails[0], true},
		{testJuncs[13], testEmails[1], true},
		{testJuncs[13], testEmails[2], true},
		{testJuncs[13], testEmails[3], false},
		{testJuncs[13], testEmails[4], false},
		{testJuncs[13], testEmails[5], false},
		{testJuncs[13], testEmails[6], false},
		{testJuncs[13], testEmails[7], true},
		{testJuncs[13], testEmails[8], false},
		{testJuncs[13], testEmails[9], true},
		{testJuncs[13], testEmails[10], true},
		{testJuncs[13], testEmails[11], false},
		{testJuncs[13], testEmails[12], false},
		{testJuncs[13], testEmails[13], true},

		{testJuncs[14], testEmails[0], true},
		{testJuncs[14], testEmails[1], true},
		{testJuncs[14], testEmails[2], true},
		{testJuncs[14], testEmails[3], false},
		{testJuncs[14], testEmails[4], false},
		{testJuncs[14], testEmails[5], false},
		{testJuncs[14], testEmails[6], false},
		{testJuncs[14], testEmails[7], true},
		{testJuncs[14], testEmails[8], true},
		{testJuncs[14], testEmails[9], true},
		{testJuncs[14], testEmails[10], true},
		{testJuncs[14], testEmails[11], false},
		{testJuncs[14], testEmails[12], false},
		{testJuncs[14], testEmails[13], true},

		{testJuncs[15], testEmails[0], false},
		{testJuncs[15], testEmails[1], false},
		{testJuncs[15], testEmails[2], false},
		{testJuncs[15], testEmails[3], false},
		{testJuncs[15], testEmails[4], false},
		{testJuncs[15], testEmails[5], false},
		{testJuncs[15], testEmails[6], false},
		{testJuncs[15], testEmails[7], true},
		{testJuncs[15], testEmails[8], false},
		{testJuncs[15], testEmails[9], false},
		{testJuncs[15], testEmails[10], true},
		{testJuncs[15], testEmails[11], false},
		{testJuncs[15], testEmails[12], false},
		{testJuncs[15], testEmails[13], true},
	}

	for i, test := range tests {
		name := fmt.Sprint(i)
		t.Run(name, func(t *testing.T) {
			res := checkTo(test.junc.To, test.email.To)
			if res != test.result {
				t.Errorf("received '%t', wanted '%t'", res, test.result)
			}
		})
	}
}

func TestCheckFrom(t *testing.T) {
	var tests = []struct {
		junc   Junction
		email  EmailData
		result bool
	}{
		{testJuncs[0], testEmails[0], true},
		{testJuncs[0], testEmails[1], true},
		{testJuncs[0], testEmails[2], true},
		{testJuncs[0], testEmails[3], true},
		{testJuncs[0], testEmails[4], true},
		{testJuncs[0], testEmails[5], true},
		{testJuncs[0], testEmails[6], true},
		{testJuncs[0], testEmails[7], true},
		{testJuncs[0], testEmails[8], true},
		{testJuncs[0], testEmails[9], true},
		{testJuncs[0], testEmails[10], true},
		{testJuncs[0], testEmails[11], true},
		{testJuncs[0], testEmails[12], true},
		{testJuncs[0], testEmails[13], true},

		{testJuncs[1], testEmails[0], true},
		{testJuncs[1], testEmails[1], true},
		{testJuncs[1], testEmails[2], true},
		{testJuncs[1], testEmails[3], true},
		{testJuncs[1], testEmails[4], true},
		{testJuncs[1], testEmails[5], true},
		{testJuncs[1], testEmails[6], true},
		{testJuncs[1], testEmails[7], true},
		{testJuncs[1], testEmails[8], true},
		{testJuncs[1], testEmails[9], true},
		{testJuncs[1], testEmails[10], true},
		{testJuncs[1], testEmails[11], true},
		{testJuncs[1], testEmails[12], true},
		{testJuncs[1], testEmails[13], true},

		{testJuncs[2], testEmails[0], true},
		{testJuncs[2], testEmails[1], true},
		{testJuncs[2], testEmails[2], true},
		{testJuncs[2], testEmails[3], true},
		{testJuncs[2], testEmails[4], true},
		{testJuncs[2], testEmails[5], true},
		{testJuncs[2], testEmails[6], true},
		{testJuncs[2], testEmails[7], true},
		{testJuncs[2], testEmails[8], true},
		{testJuncs[2], testEmails[9], true},
		{testJuncs[2], testEmails[10], true},
		{testJuncs[2], testEmails[11], true},
		{testJuncs[2], testEmails[12], true},
		{testJuncs[2], testEmails[13], true},

		{testJuncs[3], testEmails[0], true},
		{testJuncs[3], testEmails[1], true},
		{testJuncs[3], testEmails[2], true},
		{testJuncs[3], testEmails[3], true},
		{testJuncs[3], testEmails[4], true},
		{testJuncs[3], testEmails[5], true},
		{testJuncs[3], testEmails[6], true},
		{testJuncs[3], testEmails[7], true},
		{testJuncs[3], testEmails[8], true},
		{testJuncs[3], testEmails[9], true},
		{testJuncs[3], testEmails[10], true},
		{testJuncs[3], testEmails[11], true},
		{testJuncs[3], testEmails[12], true},
		{testJuncs[3], testEmails[13], true},

		{testJuncs[4], testEmails[0], true},
		{testJuncs[4], testEmails[1], false},
		{testJuncs[4], testEmails[2], true},
		{testJuncs[4], testEmails[3], true},
		{testJuncs[4], testEmails[4], false},
		{testJuncs[4], testEmails[5], true},
		{testJuncs[4], testEmails[6], false},
		{testJuncs[4], testEmails[7], true},
		{testJuncs[4], testEmails[8], true},
		{testJuncs[4], testEmails[9], true},
		{testJuncs[4], testEmails[10], false},
		{testJuncs[4], testEmails[11], false},
		{testJuncs[4], testEmails[12], false},
		{testJuncs[4], testEmails[13], true},

		{testJuncs[5], testEmails[0], true},
		{testJuncs[5], testEmails[1], true},
		{testJuncs[5], testEmails[2], false},
		{testJuncs[5], testEmails[3], true},
		{testJuncs[5], testEmails[4], true},
		{testJuncs[5], testEmails[5], false},
		{testJuncs[5], testEmails[6], false},
		{testJuncs[5], testEmails[7], true},
		{testJuncs[5], testEmails[8], true},
		{testJuncs[5], testEmails[9], true},
		{testJuncs[5], testEmails[10], true},
		{testJuncs[5], testEmails[11], true},
		{testJuncs[5], testEmails[12], false},
		{testJuncs[5], testEmails[13], false},

		{testJuncs[6], testEmails[0], true},
		{testJuncs[6], testEmails[1], false},
		{testJuncs[6], testEmails[2], false},
		{testJuncs[6], testEmails[3], true},
		{testJuncs[6], testEmails[4], false},
		{testJuncs[6], testEmails[5], false},
		{testJuncs[6], testEmails[6], false},
		{testJuncs[6], testEmails[7], true},
		{testJuncs[6], testEmails[8], true},
		{testJuncs[6], testEmails[9], true},
		{testJuncs[6], testEmails[10], false},
		{testJuncs[6], testEmails[11], false},
		{testJuncs[6], testEmails[12], false},
		{testJuncs[6], testEmails[13], false},

		{testJuncs[7], testEmails[0], true},
		{testJuncs[7], testEmails[1], false},
		{testJuncs[7], testEmails[2], true},
		{testJuncs[7], testEmails[3], true},
		{testJuncs[7], testEmails[4], false},
		{testJuncs[7], testEmails[5], true},
		{testJuncs[7], testEmails[6], false},
		{testJuncs[7], testEmails[7], true},
		{testJuncs[7], testEmails[8], true},
		{testJuncs[7], testEmails[9], true},
		{testJuncs[7], testEmails[10], false},
		{testJuncs[7], testEmails[11], false},
		{testJuncs[7], testEmails[12], false},
		{testJuncs[7], testEmails[13], true},

		{testJuncs[8], testEmails[0], true},
		{testJuncs[8], testEmails[1], false},
		{testJuncs[8], testEmails[2], true},
		{testJuncs[8], testEmails[3], true},
		{testJuncs[8], testEmails[4], false},
		{testJuncs[8], testEmails[5], true},
		{testJuncs[8], testEmails[6], false},
		{testJuncs[8], testEmails[7], true},
		{testJuncs[8], testEmails[8], true},
		{testJuncs[8], testEmails[9], true},
		{testJuncs[8], testEmails[10], false},
		{testJuncs[8], testEmails[11], false},
		{testJuncs[8], testEmails[12], false},
		{testJuncs[8], testEmails[13], true},

		{testJuncs[9], testEmails[0], true},
		{testJuncs[9], testEmails[1], false},
		{testJuncs[9], testEmails[2], true},
		{testJuncs[9], testEmails[3], true},
		{testJuncs[9], testEmails[4], false},
		{testJuncs[9], testEmails[5], true},
		{testJuncs[9], testEmails[6], false},
		{testJuncs[9], testEmails[7], true},
		{testJuncs[9], testEmails[8], true},
		{testJuncs[9], testEmails[9], true},
		{testJuncs[9], testEmails[10], false},
		{testJuncs[9], testEmails[11], false},
		{testJuncs[9], testEmails[12], false},
		{testJuncs[9], testEmails[13], true},

		{testJuncs[10], testEmails[0], true},
		{testJuncs[10], testEmails[1], true},
		{testJuncs[10], testEmails[2], false},
		{testJuncs[10], testEmails[3], true},
		{testJuncs[10], testEmails[4], true},
		{testJuncs[10], testEmails[5], false},
		{testJuncs[10], testEmails[6], false},
		{testJuncs[10], testEmails[7], true},
		{testJuncs[10], testEmails[8], true},
		{testJuncs[10], testEmails[9], true},
		{testJuncs[10], testEmails[10], true},
		{testJuncs[10], testEmails[11], true},
		{testJuncs[10], testEmails[12], false},
		{testJuncs[10], testEmails[13], false},

		{testJuncs[11], testEmails[0], true},
		{testJuncs[11], testEmails[1], true},
		{testJuncs[11], testEmails[2], false},
		{testJuncs[11], testEmails[3], true},
		{testJuncs[11], testEmails[4], true},
		{testJuncs[11], testEmails[5], false},
		{testJuncs[11], testEmails[6], false},
		{testJuncs[11], testEmails[7], true},
		{testJuncs[11], testEmails[8], true},
		{testJuncs[11], testEmails[9], true},
		{testJuncs[11], testEmails[10], true},
		{testJuncs[11], testEmails[11], true},
		{testJuncs[11], testEmails[12], false},
		{testJuncs[11], testEmails[13], false},

		{testJuncs[12], testEmails[0], true},
		{testJuncs[12], testEmails[1], true},
		{testJuncs[12], testEmails[2], false},
		{testJuncs[12], testEmails[3], true},
		{testJuncs[12], testEmails[4], true},
		{testJuncs[12], testEmails[5], false},
		{testJuncs[12], testEmails[6], false},
		{testJuncs[12], testEmails[7], true},
		{testJuncs[12], testEmails[8], true},
		{testJuncs[12], testEmails[9], true},
		{testJuncs[12], testEmails[10], true},
		{testJuncs[12], testEmails[11], true},
		{testJuncs[12], testEmails[12], false},
		{testJuncs[12], testEmails[13], false},

		{testJuncs[13], testEmails[0], true},
		{testJuncs[13], testEmails[1], false},
		{testJuncs[13], testEmails[2], false},
		{testJuncs[13], testEmails[3], true},
		{testJuncs[13], testEmails[4], false},
		{testJuncs[13], testEmails[5], false},
		{testJuncs[13], testEmails[6], false},
		{testJuncs[13], testEmails[7], true},
		{testJuncs[13], testEmails[8], true},
		{testJuncs[13], testEmails[9], true},
		{testJuncs[13], testEmails[10], false},
		{testJuncs[13], testEmails[11], false},
		{testJuncs[13], testEmails[12], false},
		{testJuncs[13], testEmails[13], false},

		{testJuncs[14], testEmails[0], true},
		{testJuncs[14], testEmails[1], false},
		{testJuncs[14], testEmails[2], false},
		{testJuncs[14], testEmails[3], true},
		{testJuncs[14], testEmails[4], false},
		{testJuncs[14], testEmails[5], false},
		{testJuncs[14], testEmails[6], false},
		{testJuncs[14], testEmails[7], true},
		{testJuncs[14], testEmails[8], true},
		{testJuncs[14], testEmails[9], true},
		{testJuncs[14], testEmails[10], false},
		{testJuncs[14], testEmails[11], false},
		{testJuncs[14], testEmails[12], false},
		{testJuncs[14], testEmails[13], false},

		{testJuncs[15], testEmails[0], true},
		{testJuncs[15], testEmails[1], false},
		{testJuncs[15], testEmails[2], false},
		{testJuncs[15], testEmails[3], true},
		{testJuncs[15], testEmails[4], false},
		{testJuncs[15], testEmails[5], false},
		{testJuncs[15], testEmails[6], false},
		{testJuncs[15], testEmails[7], true},
		{testJuncs[15], testEmails[8], true},
		{testJuncs[15], testEmails[9], true},
		{testJuncs[15], testEmails[10], false},
		{testJuncs[15], testEmails[11], false},
		{testJuncs[15], testEmails[12], false},
		{testJuncs[15], testEmails[13], false},
	}

	for i, test := range tests {
		name := fmt.Sprint(i)
		t.Run(name, func(t *testing.T) {
			res := checkFrom(test.junc.From, test.email.From, test.email.IP)
			if res != test.result {
				t.Errorf("received '%t', wanted '%t'", res, test.result)
			}
		})
	}
}

func TestSelectJunction(t *testing.T) {
	junctions = make([]Junction, len(testJuncs))

	for i, junc := range testJuncs {
		junctions[len(junctions)-i-1] = junc
	}

	var tests = []struct {
		email  EmailData
		result int
	}{
		{testEmails[0], 1},
		{testEmails[1], 4},
		{testEmails[2], 7},
		{testEmails[3], 9},
		{testEmails[4], 10},
		{testEmails[5], 11},
		{testEmails[6], 15},
		{testEmails[7], 0},
		{testEmails[8], 1},
		{testEmails[9], 1},
		{testEmails[10], 3},
		{testEmails[11], 10},
		{testEmails[12], 15},
		{testEmails[13], 6},
	}

	for i, test := range tests {
		name := fmt.Sprint(i)
		t.Run(name, func(t *testing.T) {
			res := selectJunction(test.email.To, test.email.From, test.email.IP)
			if res != test.result {
				t.Errorf("received '%d', wanted '%d'", res, test.result)
			}
		})
	}
}
