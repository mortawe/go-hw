package main

import (
	"flag"
	"fmt"
	calc "homework/calc"
	"log"
	"os"
)



func main(){
	flag.Parse()
	if len(flag.Args()) == 0 {
		log.Println("No input expression specified")
		os.Exit(2)
	}

	res, err := calc.Calculate(flag.Args()[0])
	if err != nil {
		log.Println(err)
		log.Println("Error while reading lines")
		os.Exit(1)
	}
	fmt.Println(res)
}