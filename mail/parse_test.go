package mail

import (
	"os"
	"path/filepath"
	"testing"
)

func TestReadEmail(t *testing.T) {
	entries, err := os.ReadDir("../testdata/all/")
	if err != nil {
		t.Fatal(err)
	}

	for _, e := range entries {

		f, err := os.Open(filepath.Join("../testdata/all/", e.Name()))
		if err != nil {
			t.Logf("%s: %v", e.Name(), err)
			continue
		}

		defer f.Close()

		email, err := ReadEmail(f)
		if err != nil {
			t.Logf("%s: %v", e.Name(), err)
			continue
		}

		firstline, err := FirstLine(email.Text)
		if err != nil {
			t.Logf("%s: %v", e.Name(), err)
			continue
		}

		kind := Classify(firstline)

		if kind == KindUnknown {
			kind = Classify(email.Subject)
		}

		t.Logf("%s  [%s]\n  line:    %q\n  subject: %q\n  sent:    %s",
			e.Name(), kind, firstline, email.Subject, email.Sent.Format("2006-01-02 15:04"))
	}
}
