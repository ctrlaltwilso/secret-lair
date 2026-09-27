package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Username    string
	Password    string
	Inbox       string
	MailAddress string
}

const (
	mailAddressFallback = "imap.gmail.com:993"
)

func Load() (Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		return Config{}, fmt.Errorf("failed to load .env: %w", err)
	}

	username := os.Getenv("GMAIL_USER")
	if username == "" {
		return Config{}, fmt.Errorf("GMAIL_USER is not set")
	}

	password := os.Getenv("GMAIL_APP_PASSWORD")
	if password == "" {
		return Config{}, fmt.Errorf("GMAIL_PASSWORD is not set")
	}

	inbox := os.Getenv("GMAIL_INBOX")
	if inbox == "" {
		return Config{}, fmt.Errorf("INBOX is not set")
	}

	return Config{
		Username:    username,
		Password:    password,
		Inbox:       inbox,
		MailAddress: getMailAddress(os.Getenv("MAIL_ADDRESS")),
	}, nil

}

func getMailAddress(address string) string {
	if address == "" {
		return mailAddressFallback
	}

	return address
}
