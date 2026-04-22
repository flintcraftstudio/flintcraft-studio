package view

import (
	"fmt"
	"strings"
)

// ledgerNumeral returns the editorial "№ 01" style numeral for an index.
func ledgerNumeral(i int) string {
	return fmt.Sprintf("№%02d", i)
}

// ledgerPage returns a fictional page number for the table-of-contents
// ledger. Each report claims four pages of editorial space.
func ledgerPage(i int) string {
	return fmt.Sprintf("%d", i*4)
}

// runInLabel returns the run-in editorial label for a given paragraph index
// in a Field Report's body. Falls through to a quiet em-dash for any
// paragraphs beyond the first three.
func runInLabel(i int) string {
	switch i {
	case 0:
		return "The brief —"
	case 1:
		return "What we did —"
	case 2:
		return "The outcome —"
	default:
		return "Note —"
	}
}

// joinTags renders a project's tags as a comma-separated metadata string.
func joinTags(tags []string) string {
	return strings.Join(tags, ", ")
}

// trimTrailingPeriod strips a single trailing period — the ledger row uses
// an em-dash separator, so the project Title's terminal period reads as
// punctuation noise.
func trimTrailingPeriod(s string) string {
	return strings.TrimSuffix(s, ".")
}
