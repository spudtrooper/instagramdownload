package main

import (
	"flag"

	"github.com/spudtrooper/instagramdownload/instadl"
)

func main() {
	var (
		infile = flag.String("infile", "", "Input file")
		data   = flag.String("data", "data", "Data output directory")
		dryRun = flag.Bool("dry_run", false, "Just print the actions")
	)
	flag.Parse()
	if err := instadl.DownloadAll(*infile, *data, *dryRun); err != nil {
		panic(err)
	}
}
