package smtp

import (
	netsmtp "net/smtp"
)

// MailConfig smtp config for send an email
type MailConfig struct {
	User       string
	Password   string
	Host       string
	Port       string
	Authorised netsmtp.Auth
}

// Message required info in the email
type Message struct {
	From        string
	To          []string
	Cc          []string
	Bcc         []string
	Subject     string
	Body        string
	ContentType string
}
