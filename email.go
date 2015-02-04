package main

import (
	"bytes"
	"log"
	"math"
	"net/mail"
	"net/smtp"
	"os"
	"path"
	"strconv"
	"strings"
	"text/template"
)

type emailVars struct {
	From, To, Title                 string
	Coin, Direction, Fiat, Interval string
	Difference                      float64
}

//SendEmail sends email to configured user, using
//pre-configured SMTP auth settings, about the latest
//price delta as diff.
func SendEmail(diff float64) {

	var body bytes.Buffer
	from := mail.Address{Name: "harboly-watch", Address: Config.Email.Username}
	email := emailVars{
		from.String(),
		Config.Email.Recipient,
		"",
		strings.ToUpper(Config.Coin),
		"",
		strings.ToUpper(Config.Fiat),
		Config.Interval,
		math.Abs(diff),
	}

	strDiff := strconv.FormatFloat(email.Difference, 'g', 2, 64)
	email.Title = encodeRFC2047(email.Coin + " PRICE ALERT: $" + strDiff + " " + email.Fiat + " CHANGE.")

	if math.Abs(diff) != diff {
		email.Direction = "INCREASED"
	} else {
		email.Direction = "DECREASED"
	}

	file := path.Join(GetDir(), "email.tmpl")
	if _, err := os.Stat(file); os.IsNotExist(err) {
	    file = path.Join(GetDir(), "..", "etc", "email.tmpl")
	}
	t := template.Must(template.ParseFiles(file))
	err := t.Execute(&body, email)
	LogErr(err)

	err = smtp.SendMail(Config.Email.Server+":"+strconv.Itoa(Config.Email.Port),
		smtp.PlainAuth("",
			Config.Email.Username,
			Config.Email.Password,
			Config.Email.Server,
		),
		Config.Email.Username,
		[]string{Config.Email.Recipient},
		body.Bytes())

	if err == nil {
		log.Println("email sent!")
	} else {
		log.Printf("ERR sending email: %v", err.Error())
	}
}

func encodeRFC2047(String string) string {
	addr := mail.Address{Name: String, Address: ""}
	return strings.Trim(addr.String(), " <>")
}
