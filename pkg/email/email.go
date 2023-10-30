package email

import (
    "fmt"
    "github.com/jordan-wright/email"
    "net/smtp"
    "os"
)

type Option struct {
    From    string
    To      []string
    Subject string
    Text    []byte
    Html    []byte
}

func Send(opt Option) error {
    addr := fmt.Sprintf("%s:%s", os.Getenv("MAIL_HOST"), os.Getenv("MAIL_PORT"))

    if opt.From == "" {
        opt.From = os.Getenv("MAIL_USERNAME")
    }

    if len(opt.To) == 0 {
        opt.To = []string{os.Getenv("MAIL_TO")}
    }

    if opt.Subject == "" {
        opt.Subject = os.Getenv("MAIL_SUBJECT")
    }

    e := email.NewEmail()
    e.From = opt.From
    e.To = opt.To
    e.Subject = opt.Subject
    //e.Text = opt.Text
    e.HTML = opt.Html
    err := e.Send(addr, smtp.PlainAuth("ssl", os.Getenv("MAIL_USERNAME"), os.Getenv("MAIL_PASSWORD"), os.Getenv("MAIL_HOST")))
    return err
}
