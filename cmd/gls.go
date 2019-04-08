package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/wuriyanto48/gls"
)

func main() {
	args, err := gls.ParseArguments()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	datas, err := ioutil.ReadDir(args.Path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, f := range datas {
		fmt.Println(f.Name())
		fmt.Println(f.Size())
		fmt.Println(f.ModTime())
		fmt.Println(f.Sys())
	}
}
