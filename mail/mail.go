// Package mail provides a light wrapper around SMTP functionality.
package mail

import (
	"errors"
	"fmt"
	"net/smtp"
	"strings"

	"github.com/motki/core/log"
)

var ErrRecipientUnsubscribed = errors.New("recipient has requested no email communication")

// Config describes a Sender's configuration.
type Config struct {
	SMTPAddress  string `toml:"smtp_address"`
	SMTPUsername string `toml:"smtp_username"`
	SMTPPassword string `toml:"smtp_password"`

	System Recipient `toml:"system"` // The system's address.
}

// A Recipient is an email recipient.
type Recipient struct {
	Name  string `toml:"name"`
	Email string `toml:"email"`
}

func (r Recipient) String() string {
	if r.Name != "" {
		return fmt.Sprintf("%s <%s>", r.Name, r.Email)
	}
	return r.Email
}

// A Sender handles sending of email.
type Sender struct {
	conf Config

	DoNotSend List
}

// NewSender creates a new Sender, ready for use.
func NewSender(conf Config, l log.Logger) *Sender {
	l.Debugf("mail: initializing sender, SMTP address: %s", conf.SMTPAddress)
	l.Debugf("mail: SMTP username: %s", conf.SMTPUsername)
	l.Debugf("mail: system Reply-to address: %s", conf.System)
	return &Sender{
		conf: conf,

		DoNotSend: nilList{},
	}
}

// Send sends an email with the given HTML message to the recipient.
func (s *Sender) Send(recipient Recipient, subject, message string) error {
	if s.DoNotSend.Exists(recipient) {
		return ErrRecipientUnsubscribed
	}
	p := strings.Split(s.conf.SMTPAddress, ":")
	a := smtp.PlainAuth("", s.conf.SMTPUsername, s.conf.SMTPPassword, p[0])
	msg := fmt.Sprintf(
		"From: %s\r\nReply-To: %s\r\nTo: %s\r\nSubject: %s\r\nContent-Type: text/html; charset=\"UTF-8\"\r\n\r\n%s",
		s.conf.System.String(),
		s.conf.System.String(),
		recipient.String(),
		subject,
		message,
	)
	return smtp.SendMail(s.conf.SMTPAddress, a, s.conf.System.Email, []string{recipient.Email}, []byte(msg))
}
