package main

import (
    "bytes"
    "log"
    "net/smtp"
    "strings"
    "text/template"
    "math"
    "net/mail"
    "strconv"
)

type EmailVars struct {
    From, To, Title string
    Coin, Direction, Fiat, Interval string
    Difference float64
}

func SendEmail(diff float64) {

    var body bytes.Buffer
    from := mail.Address{"harboly-watch", Config.Email.Username}
    emailVars := EmailVars{
        from.String(),
        Config.Email.Recipient,
        "",
        strings.ToUpper(Config.Coin),
        "",
        strings.ToUpper(Config.Fiat),
        Config.Interval,
        math.Abs(diff),
    }

    strDiff := strconv.FormatFloat(emailVars.Difference, 'g', 2, 64)
    emailVars.Title = encodeRFC2047(emailVars.Coin + " PRICE ALERT: $" + strDiff + " " + emailVars.Fiat + " CHANGE.")

    if math.Abs(diff) != diff {
        emailVars.Direction = "INCREASED"
    } else {
        emailVars.Direction = "DECREASED"
    }

    t := template.Must(template.ParseFiles("email.tmpl"))
    err := t.Execute(&body, emailVars)
    LogErr(err)

    err = smtp.SendMail(Config.Email.Server + ":" + strconv.Itoa(Config.Email.Port),
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

func encodeRFC2047(String string) string{
	addr := mail.Address{String, ""}
	return strings.Trim(addr.String(), " <>")
}