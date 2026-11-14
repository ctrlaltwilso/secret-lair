package mail

import (
	"regexp"
	"strings"
	"time"
)

type Kind string

const (
	KindUnknown    Kind = "unknown"
	KindTomorrow   Kind = "tomorrow"
	KindLaunch     Kind = "launch"
	KindDay        Kind = "day"
	KindDated      Kind = "dated"
	KindHours      Kind = "hours"
	KindLastChance Kind = "last"
)

var rules = []struct {
	kind    Kind
	phrases []string
}{
	{KindTomorrow, []string{"tomorrow"}},
	{KindLaunch, []string{"available now", "out now", "launch day"}},
	{KindLastChance, []string{"available until", "supplies last"}},
}

var hoursRe = regexp.MustCompile(`\d+ ?(hours?|hr)\b`)

func Classify(s string) Kind {
	s = strings.ToLower(s)
	foundMonth := FindMonth(s)
	foundDay := FindDay(s)

	switch {
	case foundMonth != "":
		return KindDated
	case foundDay != "":
		return KindDay
	case hoursRe.MatchString(s):
		return KindHours
	}

	for _, r := range rules {
		for _, p := range r.phrases {
			if strings.Contains(s, p) {
				return r.kind
			}
		}
	}

	return KindUnknown
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
