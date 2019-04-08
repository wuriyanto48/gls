package main

import (
	"os"

	"github.com/wuriyanto48/gls"
)

func main() {
	args, err := gls.ParseArguments()
	if err != nil {
		gls.PrintRedColor(err.Error())
		args.Help()
		os.Exit(1)
	}

	if err := gls.Run(args); err != nil {
		gls.PrintRedColor(err.Error())
		os.Exit(1)
	}
}
