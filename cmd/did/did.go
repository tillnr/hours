package main

import (
	"flag"
	"os"
	"strings"
	"time"

	"github.com/tillnr/hours/didfile"
)

type timeValue struct {
	time.Time
}

func (t timeValue) String() string {
	return t.Time.Local().Format(time.RFC3339)
}

func (t timeValue) Set(maybeTime string) error {
	parsed, err := parse(maybeTime)
	t.Time = parsed
	return err
}

func parse(maybeTime string) (time.Time, error) {
	loc, err := time.LoadLocation("Local")
	if err != nil {
		return time.Time{}, err
	}

	t, err := time.ParseInLocation(time.RFC3339, maybeTime, loc)
	if err == nil {
		return t, nil
	}

	return time.ParseInLocation(time.DateOnly, maybeTime, loc)
}

var until = timeValue{}

func main() {
	flag.Var(until, "until", "Make an entry lasting until the time specified. Must be after time of last entry and before the present moment. Must be in DateOnly or RFC3339 format.")
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
	if !until.Time.IsZero() {
		err = didfile.AppendUntil(entry, until.Time)
	} else {
		err = didfile.Append(entry)
	}

	if err != nil {
		println(err.Error())
	}
}
