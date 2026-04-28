package report

import(
	"io"
	"time"
)

// A Report maps activities to the time spent on them.
type Report map[string]time.Duration

// Read didfile and generate a report aggregating duration per activity from t onwards.
func From(t time.Time, didfile io.Reader) (Report, error) {
	entries, err := parseEntries(didfile)
	if err != nil {
		return Report{}, err
	}
	
	return aggregate(t, entries), nil
}

// Aggregate all entries in didfile.
func Full(didfile io.Reader) (Report, error) {
	return From(time.Time{}, didfile)
}

// Generate a report for the current day from didfile.
func Today(didfile io.Reader) (Report, error) {
	return From(dateOnly(time.Now()), didfile)
}

func dateOnly(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}
