package main

import (
	"flag"
	_ "log"
	"os"
	"strings"

)

var (
	input   string
	structs []string
)

func init() {
	flagStructs := flag.String("structs", "", "[Required] The name of schema structs to generate structs for, comma seperated\n")
	flagInput := flag.String("input", "", "[Required] The name of the input file dir\n")
	flag.Parse()

	if *flagStructs == "" || *flagInput == "" {
		flag.Usage()
		os.Exit(1)
	}

	structs = strings.Split(*flagStructs, ",")
	input = *flagInput
}

func main() {
}
