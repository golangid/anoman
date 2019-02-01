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

	if err := ParseOutput(temp, os.Stdin); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
