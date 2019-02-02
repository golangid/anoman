package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"runtime"
	"strings"
)

const (
	//Version of anoman
	Version = "1.0.0"
	// Banner of anoman
	Banner = `
	_________________________________
	| |_*_| | \| |  | |\/| |_*_| | \| |
	| |   | |\ \ |__| |  | |   | |\ | |
	***********************************
`
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
		outputName  string
		showVersion bool
		templateMD  TemplateMD
		responses   []string
	)

	flag.StringVar(&outputName, "output", "out.md", "output markdown name")
	flag.StringVar(&outputName, "o", "out.md", "output markdown name")
	flag.BoolVar(&showVersion, "v", false, "show anoman version")
	flag.BoolVar(&showVersion, "version", false, "show anoman version")

	flag.Usage = func() {
		printGreenText(Banner)
		fmt.Println()
		printGreenText("   **-----------------------------------------------**   \n")
		fmt.Println()
		printGreenText("    Anoman (Markdown Generator for your project)   \n")
		fmt.Println()
		printGreenText("	-o  | --output output markdown name eg: -o README.md\n")
		printGreenText("	-v    | --version show anoman version\n")
		printGreenText("	-h    | --help show help and usage\n")
		fmt.Println()
		printGreenText("   **-----------------------------------------------**   \n")
		fmt.Println()
	}

	flag.Parse()

	if showVersion {
		printGreenText(fmt.Sprintf("%s version %s (runtime: %s) %s", os.Args[0], Version, runtime.Version(), "\n"))
		os.Exit(0)
	}

	reader := bufio.NewReader(os.Stdin)

	for _, question := range questions {

		printGreenText(fmt.Sprintf("%s : ", question))

		response, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		response = strings.TrimSpace(response)

		fmt.Println()

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
