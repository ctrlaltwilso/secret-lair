package main

import (
	"log"
	"secretlair/config"
	"secretlair/mail"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("%v", err)
	}

	err = mail.MailClient(cfg.Username, cfg.Password, cfg.Inbox, cfg.MailAddress)
	if err != nil {
		log.Fatalf("%v", err)
	}

}
