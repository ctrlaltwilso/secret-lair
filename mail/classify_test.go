package mail

import (
	"testing"
)

func TestClassify(t *testing.T) {
	var messages = []struct {
		kind   Kind
		phrase string
	}{

		{KindDated, "Secret Lair Commander Deck: Odds and Ends arrives on September 28"},
		{KindDated, "Explore three drops coming November 9"},
		{KindDay, "Chaos Vault Opens Monday"},
		{KindTomorrow, "Corn Maze Superdrop arrives tomorrow"},
		{KindHours, "In 25 hours, the Secret Lair Commander Deck: Odds and Ends arrives"},
		{KindHours, "Only 24 hours until Festival in a Box: Atlanta 2026!"},
		{KindLaunch, "The Corn Maze Superdrop is available now"},
		{KindLastChance, "Available while supplies last"},
		{KindUnknown, "Are you lost in here too?"},
	}

	for _, m := range messages {
		got := Classify(m.phrase)
		if got != m.kind {
			t.Errorf("Classify(%q) = %s, want %s", m.phrase, got, m.kind)
		}
	}
}
