package main

import (
	"fmt"
	"os"
	"time"

	"github.com/timematic/anytime"
)

func usage() string {
	return "Usage: anytime-parse date [zone name or abbreviation]\nExample:\n anytime-parse '2006-01-02'\n anytime-parse '2006-01-02' UTC"
}

func main() {
	if len(os.Args) == 1 {
		println(usage())
		os.Exit(1)
	}

	loc := time.Local
	input := os.Args[1]
	if len(os.Args) > 2 {
		var err error
		if loc, err = time.LoadLocation(os.Args[2]); err != nil {
			fmt.Printf("%s %v", os.Args[2], err)
			os.Exit(1)
		}
	}

	output, err := anytime.ParseInLocation(input, loc)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(output)
}
