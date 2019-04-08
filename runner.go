package gls

import (
	"io/ioutil"
	"os"
)

// Run function
func Run(args *Arguments) error {

	datas, err := ioutil.ReadDir(args.Path)
	if err != nil {
		return err
	}

	var assets Assets

	for _, f := range datas {
		var asset Asset
		asset.Name = f.Name()
		asset.Size = f.Size()
		asset.ModifiedDate = f.ModTime()

		assets = append(assets, &asset)
	}

	err = parseOutput(assets, os.Stdout)
	if err != nil {
		return err
	}

	return nil
}
