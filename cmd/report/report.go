package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/tillnr/hours/didfile"
	"github.com/tillnr/hours/report"
)

type timeValue struct {
	*time.Time
}

func (t timeValue) String() string {
	if t.Time == nil {
		return time.Time{}.Local().Format(time.RFC3339)
	}
	return t.Time.Local().Format(time.RFC3339)
}

func (t timeValue) Set(maybeTime string) error {
	parsed, err := parse(maybeTime)
	*t.Time = parsed
	return err
}

func parse(maybeTime string) (time.Time, error) {
	loc, err := time.LoadLocation("Local")
	if err != nil {
		return time.Time{}, err
	}

	t, err := time.ParseInLocation(time.DateOnly, maybeTime, loc)
	if err != nil {
		return time.ParseInLocation(time.RFC3339, maybeTime, loc)
	}

	return t, nil
}

func main() {
	var from = time.Time{}
	flag.Var(timeValue{&from}, "from", `
	`)
	flag.Parse()

	didfile, err := didfile.Open()
	if err != nil {
		println(err.Error())
		return
	}
	defer didfile.Close()

	report, err := report.From(from, didfile.Reader())
	if err != nil {
		println(err.Error())
		return
	}

	for activity, duration := range report {
		fmt.Printf("%v %v \n", activity, duration.Hours())
	}
}
