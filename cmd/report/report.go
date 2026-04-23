package main

import (
	"os"
	"time"

	"github.com/tillnr/hours"
)

func main() {
	report, err := hours.FullReport() 
	if err != nil {
		println("Could not make report:", err.Error())
		os.Exit(1)
	}
	
	for day, entries := range(report) {
		for key, duration := range(entries) {
			if duration == 0 {
				continue
			}

			println(day.Local().Format(time.DateOnly), key, duration.String())
		}
	}
}
