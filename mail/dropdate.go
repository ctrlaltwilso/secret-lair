package mail

import (
	"fmt"
	"strings"
	"time"
)

func DropDate(kind Kind, text string, sent time.Time) (time.Time, error) {
	switch kind {
	case KindTomorrow:
		return sent.AddDate(0, 0, 1), nil
	case KindDated:
		findMonth := FindMonth(strings.ToLower(text))
		date, err := time.Parse("January 2", findMonth)
		if err != nil {
			return time.Time{}, err
		}

		d := time.Date(sent.Year(), date.Month(), date.Day(), 0, 0, 0, 0, sent.Location())
		if d.Before(sent) {
			d = d.AddDate(1, 0, 0)
		}

		return d, nil

	case KindDay:
		target, ok := FindDay(strings.ToLower(text))
		if !ok {
			return time.Time{}, fmt.Errorf("no weekday in %q", text)
		}
		days := (int(target) - int(sent.Weekday()) + 7) % 7
		return sent.AddDate(0, 0, days), nil

	case KindHours:
		n, ok := FindHours(text)
		if !ok {
			return time.Time{}, fmt.Errorf("no hours in %q", text)
		}

		return sent.Add(time.Duration(n) * time.Hour), nil

	}

	return time.Time{}, fmt.Errorf("no date for kind %s", kind)

}
