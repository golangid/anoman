package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	questions = []string{
		"project name ?",
		"subtitle ?",
		"author name ?",
		"year ?",
	}
)

func main() {
	var (
		outputName string
	)

	flag.StringVar(&outputName, "output", "out.md", "output markdown name")
	flag.StringVar(&outputName, "o", "out.md", "output markdown name")

	flag.Parse()

	reader := bufio.NewReader(os.Stdin)

	var (
		temp TemplateMD

		responses []string
	)

	for _, question := range questions {

		fmt.Printf("%s : ", question)

		response, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		response = strings.TrimSpace(response)

		responses = append(responses, response)
	}

	temp.ProjectName = responses[0]
	temp.Subtitle = responses[1]
	temp.Author = responses[2]
	temp.Year = responses[3]

	// for testing you can use os.Stdin

	// if err := ParseOutput(temp, os.Stdin); err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	file, err := os.Create(outputName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	if err := ParseOutput(temp, file); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
