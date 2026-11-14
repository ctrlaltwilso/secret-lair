package mail

import (
	"fmt"
	"io"
	"strings"

	_ "github.com/emersion/go-message/charset"
	gomail "github.com/emersion/go-message/mail"
)

func PlainText(r io.Reader) (string, string, error) {
	mr, err := gomail.CreateReader(r)
	if err != nil {
		return "", "", err
	}

	s, err := mr.Header.Subject()
	if err != nil {
		return "", "", err
	}

	for {
		p, err := mr.NextPart()
		if err == io.EOF {
			break
		} else if err != nil {
			return "", "", err
		}

		switch h := p.Header.(type) {
		case *gomail.InlineHeader:
			t, _, _ := h.ContentType()
			if t == "text/plain" {
				b, _ := io.ReadAll(p.Body)
				return string(b), string(s), nil
			}
		}
	}

	return "", "", nil
}

func FirstLine(s string) (string, error) {
	before, _, found := strings.Cut(s, "(")
	if !found {
		return "", fmt.Errorf("no ( found")
	}

	fl := strings.Join(strings.Fields(before), " ")
	return fl, nil

}
