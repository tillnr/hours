package main

import (
	"os"
	"strings"

	"github.com/tillnr/hours"
)

func main() {
	if len(os.Args) <= 1 {
		usage()
	}

	err := hours.Add(strings.Join(os.Args[1:], " "))
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
}

func usage() {
	println("usage: did ENTRY ...")
	os.Exit(1)
}
