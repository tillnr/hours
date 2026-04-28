package main

import (
	"os"
	"strings"

	"github.com/tillnr/hours/didfile"
)

func main() {
	if len(os.Args) <= 1 {
		usage()
	}

	didfile, err := didfile.Open()	
	if err != nil {
		println(err.Error())
		return
	}

	didfile.Append(strings.Join(os.Args[1:], " "))
}

func usage() {
	println("usage: did ENTRY ...")
	os.Exit(1)
}
