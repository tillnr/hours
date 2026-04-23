package hours

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"slices"
	"time"
)

// A report maps the activities to time spent on them.
type Report map[string]time.Duration

// Generate a report mapping days to their respective report.
func FullReport() (map[time.Time]Report, error) {
	file, fileErr := didfile(openRead)
	if fileErr != nil {
		return map[time.Time]Report{}, fileErr
	}

	fileBytes, readErr := io.ReadAll(file)
	file.Close()
	if readErr != nil {
		return map[time.Time]Report{}, readErr
	}
	

	return report(fileBytes)
}

func report(b []byte) (map[time.Time]Report, error) {
	entries, parseErr := parse(b)
	if parseErr != nil {
		return map[time.Time]Report{}, parseErr
	}

	if len(entries) == 0 {
		return map[time.Time]Report{}, errors.New("nothing to report.")
	}
	
	if !slices.IsSortedFunc(entries, compareEntriesByTime) {
		return map[time.Time]Report{}, errors.New("didfile timestamps not in ascending order.")
	}
	
	// We append a dummy entry to calculate how long the current task has been going on.
	entries = append(entries, entry{"", time.Now()})
	report := map[time.Time]Report{}
	for i := 0; i < len(entries) - 1; i++ {
		entry := entries[i]
		durationToday, durationTomorrow := durations(entry, entries[i+1])
		if todayEntries, ok := report[entry.Today()]; ok {
			todayEntries[entry.Activity] += durationToday
		} else {
			report[entry.Today()] = Report{}
			report[entry.Today()][entry.Activity] = durationToday
		}

		if tomorrowEntries, ok := report[entry.Tomorrow()]; ok && durationTomorrow != 0 {
			tomorrowEntries[entry.Activity] += durationTomorrow
		} else {
			report[entry.Tomorrow()] = Report{}
			report[entry.Tomorrow()][entry.Activity] = durationTomorrow
		}

	}

	return report, nil
}

type entry struct {
	Activity string
	time.Time
}

func (e entry) Today() time.Time {
	zone := time.FixedZone(e.Time.Zone())
	year, month, day := e.Time.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, zone).Round(0).UTC()
}

func (e entry) Tomorrow() time.Time {
	return e.Today().AddDate(0,0,1)
}

func parse(b []byte) ([]entry, error) {
	if len(b) == 0 {
		return []entry{}, nil
	}

	maybeEntries := bytes.Split(b, []byte("\n"))
	entries := []entry{}
	for line, maybeEntry := range maybeEntries {
		trimmed := bytes.TrimSpace(maybeEntry)
		if len(trimmed) == 0 {
			continue
		}

		entry, parseErr := parseEntry(trimmed)
		if parseErr != nil {
			return entries, errors.Join(DidfileEntryError(line), parseErr)
		}

		entries = append(entries, entry)
	}

	return entries, nil
}

func parseEntry(b []byte) (entry, error) {
	// An entry consists of a valid RFC3339 timestamp in UTC, a colon and non empty
	// entry text. Splitting by colon must therefore yield 4 non empty components.
	components := bytes.FieldsFunc(b, isColon)
	if len(components) < 4 {
		return entry{}, errors.New("not enough components in entry.")
	}

	maybeTime := string(bytes.Join(components[:3], []byte{':'}))
	timestamp, parseErr := time.Parse(time.RFC3339, maybeTime)
	if parseErr != nil {
		return entry{}, parseErr
	}

	if len(bytes.TrimSpace(components[3])) == 0 {
		return entry{}, errors.New("entry empty.")
	}
	
	return entry{string(components[3]), timestamp.Round(0).UTC()}, nil
}

func durations(start, end entry) (time.Duration, time.Duration) {
	if end.Time.Before(start.Tomorrow()) {
		return end.Time.Sub(start.Time), time.Duration(0)
	}

	return start.Tomorrow().Sub(start.Time), end.Time.Sub(start.Tomorrow())
}

func isColon(r rune) bool {
	return r == ':'
}

func compareEntriesByTime(a, b entry) int {
	return a.Time.Compare(b.Time)
}

type DidfileEntryError int

func (e DidfileEntryError) Error() string {
	return "invalid entry at line " + fmt.Sprint(int(e)) + ":"
}
