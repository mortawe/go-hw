package main

import (
	"flag"
	"log"
	"os"

	"hw1.1/sort"
)

func main() {
	var flags sort.Flags
	flags = sort.InitFlags(flags)
	if len(flag.Args()) == 0 {
		log.Println("No input file specified")
		os.Exit(1)
	}

	lines, err := sort.ReadLines(flag.Args()[0])

	if err != nil {
		log.Println(err)
		log.Println("Error while reading lines")
		os.Exit(2)
	}

	lines, err = sort.SortWithFlags(lines, flags)

	if sort.WriteResult(lines, flags) != nil {
		log.Println("Error while writing lines")
		os.Exit(3)
	}
}
