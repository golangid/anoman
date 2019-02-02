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
		templateMD TemplateMD
		responses  []string
	)

	flag.StringVar(&outputName, "output", "out.md", "output markdown name")
	flag.StringVar(&outputName, "o", "out.md", "output markdown name")

	flag.Parse()

	reader := bufio.NewReader(os.Stdin)

	for _, question := range questions {

		printGreenText(fmt.Sprintf("%s : ", question))

		response, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		response = strings.TrimSpace(response)

		responses = append(responses, response)
	}

	templateMD.ProjectName = responses[0]
	templateMD.Subtitle = responses[1]
	templateMD.Author = responses[2]
	templateMD.Year = responses[3]

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

	if err := ParseOutput(templateMD, file); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func printGreenText(s string) {
	fmt.Printf("\033[32m%s\033[0m", s)
}
