package main

import (
	"fmt"
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

	for day, activities := range report {
		for activity, duration := range activities {
			if duration == 0 {
				continue
			}

			fmt.Printf("%v %v %v\n", day.Local().Format(time.DateOnly), activity, duration.String())
		}
	}
}
