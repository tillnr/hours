package main

import (
	"flag"
	"os"
	"strings"
	"time"

	"github.com/tillnr/hours/didfile"
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

	now := time.Now()
	t, err := time.ParseInLocation(time.TimeOnly, maybeTime, loc)
	if err != nil {
		return time.ParseInLocation(time.RFC3339, maybeTime, loc)
	}

	return time.Date(now.Year(), now.Month(), now.Day(), t.Hour(), t.Minute(), t.Second(), 0, loc), nil
}

func main() {
	until := time.Time{}
	flag.Var(timeValue{&until}, "until", `
Make an entry lasting until the time specified.
Must be after time of last entry and before the present moment.
Must be either a full RFC3339 location, or simple timestamp in HH:MM:SS format.
The simple timestamp will be intrepreted as refering to today in local time.`)
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	didfile, err := didfile.Open()
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	entry := strings.Join(args, " ")
	if !until.IsZero() {
		err = didfile.AppendUntil(entry, until)
	} else {
		err = didfile.Append(entry)
	}

	if err != nil {
		println(err.Error())
	}
}
