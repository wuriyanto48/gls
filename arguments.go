package gls

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

const (
	// Version const
	Version = "0.0.0"
)

// Arguments parser
type Arguments struct {
	Path string
	Help func()
}

// ParseArguments will parse arguments from user's input
func ParseArguments() (*Arguments, error) {
	args := os.Args

	var (
		path string
	)

	flag.Usage = func() {
		PrintGreenColor("   gls (ls for human)   ")
		fmt.Println()
		PrintGreenColor("	-h    | --help show help and usage")
	}

	flag.Parse()

	if len(args) <= 1 {
		path = "."
	} else {
		path = args[1]
	}

	fi, err := os.Stat(path)
	if err != nil {
		return &Arguments{Help: flag.Usage}, errors.New("input is not a valid directory")
	}

	mode := fi.Mode()

	if !mode.IsDir() {
		return &Arguments{Help: flag.Usage}, errors.New("input is not a valid directory")
	}

	return &Arguments{
		Path: path,
		Help: flag.Usage,
	}, nil
}
