package mail

import (
	"fmt"
	"io"
	"regexp"
	"strings"
	"time"

	_ "github.com/emersion/go-message/charset"
	gomail "github.com/emersion/go-message/mail"
)

func PlainText(r io.Reader) (string, error) {
	mr, err := gomail.CreateReader(r)
	if err != nil {
		return "", err
	}

	for {
		p, err := mr.NextPart()
		if err == io.EOF {
			break
		} else if err != nil {
			return "", err
		}

		switch h := p.Header.(type) {
		case *gomail.InlineHeader:
			t, _, _ := h.ContentType()
			if t == "text/plain" {
				b, _ := io.ReadAll(p.Body)
				return string(b), nil
			}
		}
	}

	return "", nil
}

func FirstLine(s string) (string, error) {
	before, _, found := strings.Cut(s, "(")
	if !found {
		return "", fmt.Errorf("no ( found")
	}

	fl := strings.Join(strings.Fields(before), " ")
	return fl, nil

}

type Kind string

const (
	KindUnknown  Kind = "unknown"
	KindTomorrow Kind = "tomorrow"
	KindLaunch   Kind = "launch"
	KindDay      Kind = "day"
	KindDated    Kind = "dated"
)

func Classify(s string) Kind {
	s = strings.ToLower(s)
	foundMonth := FindMonth(s)
	foundDay := FindDay(s)

	switch {
	case foundMonth != "":
		return KindDated
	case foundDay != "":
		return KindDay
	case strings.Contains(s, "tomorrow"):
		return KindTomorrow
	case strings.Contains(s, "available now"):
		return KindLaunch
	default:
		return KindUnknown
	}
}

func monthNames() []string {
	var months []string
	for m := time.January; m <= time.December; m++ {
		months = append(months, strings.ToLower(m.String()))
	}

	return months
}

var monthRe = regexp.MustCompile(`(` + strings.Join(monthNames(), "|") + `) \d{1,2}`)

func FindMonth(s string) string {
	date := monthRe.FindString(s)
	return date
}

func FindDay(s string) string {
	for d := time.Sunday; d <= time.Saturday; d++ {
		name := strings.ToLower(d.String())
		if strings.Contains(s, name) {
			return name
		}
	}

	return ""
}
