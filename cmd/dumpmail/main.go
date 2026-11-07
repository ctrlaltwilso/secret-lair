// dumpmail saves every message in the configured label to testdata/ as .eml
// files, so parsing can be developed against real emails without IMAP.
package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"secretlair/config"
	"secretlair/mail"
)

const outDir = "testdata/all"

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("%v", err)
	}

	c, err := mail.Connect(cfg.Username, cfg.Password, cfg.MailAddress)
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer c.Close()

	m, err := mail.CheckMailbox(c, cfg.Inbox)
	if err != nil {
		log.Fatalf("%v", err)
	}

	messages, err := mail.FetchAllMessages(m, c, cfg.Inbox)
	if err != nil {
		log.Fatalf("%v", err)
	}

	if err := os.MkdirAll(outDir, 0755); err != nil {
		log.Fatalf("failed to create %s: %v", outDir, err)
	}

	for _, msg := range messages {
		raw := msg.BodySection[0].Bytes
		name := fmt.Sprintf("%s_%04d.eml", msg.Envelope.Date.Format("2006-01-02"), msg.SeqNum)
		if err := os.WriteFile(filepath.Join(outDir, name), raw, 0644); err != nil {
			log.Fatalf("failed to save %s: %v", name, err)
		}
	}

	log.Printf("saved %d messages to %s", len(messages), outDir)

	if err := c.Logout().Wait(); err != nil {
		log.Fatalf("failed to logout: %v", err)
	}
}
