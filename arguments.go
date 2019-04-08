package gls

import (
	"errors"
	"os"
)

// Arguments parser
type Arguments struct {
	Path string
}

// ParseArguments will parse arguments from user's input
func ParseArguments() (*Arguments, error) {
	args := os.Args

	var (
		path string
	)

	if len(args) <= 1 {
		path = "."
	} else {
		path = args[1]
	}

	fi, err := os.Stat(path)
	if err != nil {
		return nil, errors.New("input is not a valid directory")
	}

	mode := fi.Mode()

	if !mode.IsDir() {
		return nil, errors.New("input is not a valid directory")
	}

	return &Arguments{
		Path: path,
	}, nil
}
