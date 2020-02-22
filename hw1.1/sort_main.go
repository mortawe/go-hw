package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"hw1.1/sort"
)

func initFlags(flags sort.Flags) sort.Flags {

	flag.BoolVar(&flags.IsFold, "f", false, "ignore register")
	flag.BoolVar(&flags.IsUnique, "u", false, "shows first of equals")
	flag.BoolVar(&flags.IsReversed, "r", false, "reverse")
	flag.StringVar(&flags.OutputFile, "o", "", "output file")
	flag.BoolVar(&flags.IsNumbers, "n", false, "number sort")
	flag.IntVar(&flags.ColumnNum, "k", 1, "column number")

	flag.Parse()
	return flags
}
func writeResult(lines string, flags sort.Flags) error {
	if flags.OutputFile != "" {
		return writeFile(lines, flags)
	}
	fmt.Println(lines)
	return nil
}

func writeFile(lines string, flags sort.Flags) error {

	file, err := os.Create(flags.OutputFile)
	defer file.Close()
	if err != nil {
		return err
	}
	_, err = file.Write([]byte(lines))
	if err != nil {
		return err
	}

	return nil
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	var flags sort.Flags
	flags = initFlags(flags)
	if len(flag.Args()) == 0 {
		log.Println("No input file specified")
		os.Exit(1)
	}

	lines, err := readLines(flag.Args()[0])

	if err != nil {
		log.Println(err)
		log.Println("Error while reading lines")
		os.Exit(2)
	}

	output, err := sort.SortWithFlags(lines, flags)
	if err != nil {
		log.Println(err)
		os.Exit(2)
	}
	if writeResult(output, flags) != nil {
		log.Println("Error while writing lines")
		os.Exit(3)
	}
}
