package mail

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestPlainText(t *testing.T) {
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

		p, err := PlainText(f)
		if err != nil {
			t.Logf("%s: %v", e.Name(), err)
			continue
		}

		fmt.Println("Email: Captured")

		firstline, err := FirstLine(p)
		if err != nil {
			t.Logf("%s: %v", e.Name(), err)
			continue
		}

		fmt.Println("First Line:", firstline)
		fmt.Println("Kind:", Classify(firstline))
	}

}
