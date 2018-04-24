package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
)

var fromBase = flag.Int("f", 10, "the base the number is represented in")
var toBase = flag.Int("t", 16, "the base the number will be converted to")
var value = flag.String("v", "", "the value to convert")

func main() {
	flag.Parse()
	if *value == "" && flag.Arg(0) != "" {
		*value = flag.Arg(0)
	}
	if *fromBase < 2 || *fromBase > 36 || *toBase < 2 || *toBase > 36 {
		log.Fatal("base not supported, base should be between 2 and 36 included")
	}
	v, err := strconv.ParseInt(*value, *fromBase, 64)
	if err != nil {
		log.Fatal("Value is invalid or not specified")
	}
	fmt.Println(strconv.FormatInt(v, *toBase))
}
