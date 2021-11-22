package main

import (
	"flag"
	"log"
	"os"

	"github.com/spraints/cgroups-memory-experiments/parent"
	"github.com/spraints/cgroups-memory-experiments/sizes"
)

func main() {
	if _, err := os.Stat("/.dockerenv"); os.IsNotExist(err) {
		log.Fatal("must run in a docker container")
	}

	bytesPerSecond := flag.String("rate", "1m", "bytes to leak per second in each child")
	children := flag.Int("count", 1, "number of children to run")
	flag.Parse()

	bps, err := sizes.ParseBytes(*bytesPerSecond)
	if err != nil {
		log.Fatal(err)
	}

	log.SetPrefix("[parent] ")
	parent.Run(*children, bps)
}
