package main

import (
	"fmt"

	"github.com/tillnr/hours/didfile"
	"github.com/tillnr/hours/report"
)

func main() {
	didfile, err := didfile.Open()
	if err != nil {
		println(err.Error())
		return
	}
	defer didfile.Close()

	report, err := report.Full(didfile.Reader())
	if err != nil {
		println(err.Error())
		return
	}

	for activity, duration := range report {
		fmt.Printf("%v %v \n", activity, duration.Hours())
	}
}
