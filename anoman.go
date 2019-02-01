package main

import (
	"flag"
	"fmt"
)

func main() {
	var (
		outputName string
	)

	flag.StringVar(&outputName, "output", "out.md", "output markdown name")
	flag.StringVar(&outputName, "o", "out.md", "output markdown name")

	flag.Parse()

	fmt.Println(outputName)
}
