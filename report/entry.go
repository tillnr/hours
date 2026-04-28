package report

import (
	"bufio"
	"errors"
	"io"
	"strings"
	"time"
)

type entry struct {
	activity string
	time.Time
}

func parseEntries(r io.Reader) ([]entry, error) {
	entries := []entry{}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		e, err := parse(scanner.Text())
		if err != nil {
			return []entry{}, err
		}

		entries = append(entries, e)	
	}

	return entries, nil
}

const (
	// Length of textual representation of a valid RFC3339 timestamp in UTC.
	timestampEndIndex int = 20
	// Timestamp and entry are separated by a colon.
	entryStartIndex = 21
)

func parse(row string) (entry, error) {
	if len(row) < entryStartIndex {
		return entry{}, errors.New("entry is not long enough to be valid.")
	}

	maybeTime := row[:timestampEndIndex]
	timestamp, parseErr := time.Parse(time.RFC3339, maybeTime)
	if parseErr != nil {
		return entry{}, parseErr
	}

	maybeEntry := strings.TrimSpace(row[entryStartIndex:])
	if len(maybeEntry) == 0 {
		return entry{}, errors.New("entry empty.")
	}

	return entry{string(maybeEntry), timestamp}, nil
}
