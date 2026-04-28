package report

import "time"

func aggregate(t time.Time, entries []entry) Report {
	// To be able to calculate how long the current activity has been going on for, we
	// need to insert a dummy entry.
	entries = append(entries, entry{"", time.Now()})

	report := make(map[string]time.Duration)
	for i := 0; i < len(entries) - 1; i++ {
		// We want to account for the whole time spent on the current activity, if
		// part of it falls into the the 
		activity := entries[i].activity
		start := entries[i].Time
		end := entries[i+1].Time

		if start.After(t) && end.After(t) {
			report[activity] += end.Sub(start)
			continue
		}

		if end.After(t) {
			report[activity] += end.Sub(t)
		}
	}
	
	return Report(report)
}
