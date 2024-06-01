package support

import (
	"fmt"
	"log"
	"net/smtp"
	e "sample-gin-ddd/pkg/errors"

	"sample-gin-ddd/pkg/infrastracture/config"
	"sample-gin-ddd/pkg/util"
	"strings"
)

type SendMailParam struct {
	From    string
	To      []string
	Subject string
	Body    string
}

func send(p *SendMailParam) error {
	conf := config.Mail()
	server := fmt.Sprintf("%s:%s", conf.HostName, conf.Port)
	auth := smtp.CRAMMD5Auth(conf.Username, conf.Password)
	msg := []byte(fmt.Sprintf("To: %s\nSubject: %s\n\n%s", strings.Join(p.To, ","), p.Subject, p.Body))

	if err := smtp.SendMail(server, auth, p.From, p.To, msg); err != nil {
		log.Printf("Send mail error: %s", err)
		return err
	}

	return nil
}

func SendRegisterMail(mail string) error {
	url := config.GetEnv("TEMPORARY_REGISTER_ACCOUNT_BASE_URL") + util.EmailToId(mail)
	param := &SendMailParam{
		From:    config.MAIL_FROM,
		To:      []string{mail},
		Subject: config.MAIL_SUBJECT,
		Body:    fmt.Sprintf(config.MAIL_BODY, url),
	}
	if err := send(param); err != nil {
		return e.INTERNAL_SERVER_ERROR
	}

	return nil
}
