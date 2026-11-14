package mail

import (
	"fmt"
	"io"
	"strings"
	"time"

	_ "github.com/emersion/go-message/charset"
	gomail "github.com/emersion/go-message/mail"
)

type Email struct {
	Subject string
	Text    string
	Sent    time.Time
}

func ReadEmail(r io.Reader) (Email, error) {
	mr, err := gomail.CreateReader(r)
	if err != nil {
		return Email{}, err
	}

	subject, err := mr.Header.Subject()
	if err != nil {
		return Email{}, err
	}

	sent, err := mr.Header.Date()
	if err != nil {
		return Email{}, err
	}

	for {
		p, err := mr.NextPart()
		if err == io.EOF {
			break
		} else if err != nil {
			return Email{}, err
		}

		switch h := p.Header.(type) {
		case *gomail.InlineHeader:
			t, _, _ := h.ContentType()
			if t == "text/plain" {
				b, _ := io.ReadAll(p.Body)
				return Email{Subject: subject, Text: string(b), Sent: sent}, nil
			}
		}
	}

	return Email{}, nil
}

func FirstLine(s string) (string, error) {
	before, _, found := strings.Cut(s, "(")
	if !found {
		return "", fmt.Errorf("no ( found")
	}

	fl := strings.Join(strings.Fields(before), " ")
	return fl, nil

}
