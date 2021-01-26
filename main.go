package main

import (
	"log"

	"github.com/ganboonhong/gotp/cmd"
)

func main() {
	log.SetFlags(log.Lshortfile)
	cmd.Execute()
}
