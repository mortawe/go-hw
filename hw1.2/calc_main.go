package main

import (
	"fmt"
	"log"
	"os"

	"hw1.2/calc"
)

func main() {

	if len(os.Args) == 0 {
		log.Println("No input expression specified")
		os.Exit(2)
	}
	expr := ""
	for r := 1; r < len(os.Args); r++ {
		expr += os.Args[r]
	}
	res, err := calc.Calculate(expr)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Println(res)
}
