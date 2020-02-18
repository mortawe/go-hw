package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"hw1.2/calc"
)



func main(){
	flag.Parse()
	if len(flag.Args()) == 0 {
		log.Println("No input expression specified")
		os.Exit(2)
	}

	res, err := calc.Calculate(flag.Arg(0))
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Println(res)
}