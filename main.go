package main

import (
	"bitacoin_client/cli"
	"flag"
	"log"
)

func main() {
	flag.Usage = cli.Usage
	flag.Parse()
	if err := cli.Dispatch(flag.Args()); err != nil {
		log.Fatal(err.Error())
	}
}
