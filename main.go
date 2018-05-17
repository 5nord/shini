package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	file    string
	doQuery bool
)

func init() {
	flag.StringVar(&file, "f", "-", "Name of the TOML file (default is stdin)")
	flag.BoolVar(&doQuery, "q", false, "Perform TOML query")
}

func fatal(msg string) {
	fmt.Fprintf(os.Stderr, "shini: %s\n", msg)
	os.Exit(1)
}

func main() {

	flag.Parse()

	config, err := readConfig(file)
	if err != nil {
		fatal(err.Error())
	}

	var result interface{}
	if doQuery {
		if len(flag.Args()) != 1 {
			flag.Usage()
			os.Exit(1)
		}
		result, err = config.Query(flag.Arg(0))
		if err != nil {
			fatal(err.Error())
		}
	} else {
		result = config.Get(flag.Args())
	}

	err = Print(result)
	if err != nil {
		os.Exit(1)
	}
}
