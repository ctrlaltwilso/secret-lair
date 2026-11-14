package mail

import (
	"testing"
	"time"
)

func TestDropDate(t *testing.T) {
	var messages = []struct {
		kind   Kind
		phrase string
		sent   time.Time
		want   time.Time
	}{
		{
			KindTomorrow,
			"Corn Maze Superdrop arrives tomorrow",
			time.Date(2026, time.October, 11, 16, 0, 0, 0, time.UTC),
			time.Date(2026, time.October, 12, 16, 0, 0, 0, time.UTC),
		},
		{
			KindDated,
			"Secret Lair Commander Deck: Odds and Ends arrives on September 28",
			time.Date(2026, time.September, 26, 16, 0, 0, 0, time.UTC),
			time.Date(2026, time.September, 28, 0, 0, 0, 0, time.UTC),
		},
		{
			KindDated,
			"Secret Lair Commander Deck: Odds and Ends arrives on January 5",
			time.Date(2026, time.December, 20, 16, 0, 0, 0, time.UTC),
			time.Date(2027, time.January, 5, 0, 0, 0, 0, time.UTC),
		},
		{
			KindDay,
			"Chaos Vault Opens Monday",
			time.Date(2026, time.September, 25, 16, 0, 0, 0, time.UTC),
			time.Date(2026, time.September, 28, 16, 0, 0, 0, time.UTC),
		},
		{
			KindHours,
			"In 25 hours, the Secret Lair Commander Deck: Odds and Ends arrives",
			time.Date(2026, time.September, 27, 16, 0, 0, 0, time.UTC),
			time.Date(2026, time.September, 28, 17, 0, 0, 0, time.UTC),
		},
	}

	for _, m := range messages {
		got, err := DropDate(m.kind, m.phrase, m.sent)
		if err != nil {
			t.Errorf("%s", err)
			continue
		}

		if !got.Equal(m.want) {
			t.Errorf("%s: got: %q\nwant: %q", m.kind, got, m.want)
		}

	}
}
