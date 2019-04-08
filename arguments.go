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
	Path     string
	ShowSize bool
	Help     func()
}

// ParseArguments will parse arguments from user's input
func ParseArguments() (*Arguments, error) {

	var (
		path     = "."
		showSize bool
	)

	flag.BoolVar(&showSize, "s", false, "show size")
	flag.BoolVar(&showSize, "size", false, "show size")

	flag.Usage = func() {
		PrintGreenColor("   gls (ls for human)   ")
		fmt.Println()
		PrintGreenColor("	-h    | --help show help and usage")
	}

	flag.Parse()

	args := flag.Args()

	if len(args) >= 1 {
		path = args[0]
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
		Path:     path,
		Help:     flag.Usage,
		ShowSize: showSize,
	}, nil
}
