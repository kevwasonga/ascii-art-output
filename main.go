package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"

	"ascii/asciiArt"
)

func main() {
	// Defines the command-line flags
	outputfile := flag.String("output", "", "Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
	flag.Parse()

	// Check if the output flag is provided
	if *outputfile != "" {
		// Validate the output file format using regex
		validOutput := regexp.MustCompile(`^[a-zA-Z0-9_-]+\.txt$`)
		args := os.Args[1:]

		// Check if the output flag is in the correct format
		if len(args) < 1 || !strings.HasPrefix(args[0], "--output=") || !validOutput.MatchString(args[0]) {
			printUsageAndExit()
		}
	}

	// Remaining arguments after the flag
	args := flag.Args()

	// Validate the number of the rest of arguments
	if len(args) < 1 || len(args) > 2 {
		printUsageAndExit()
	}

	word := args[0]
	fileName := "standard"
	if len(args) == 2 {
		fileName = args[1]
	}

	// Print a new line and exit in case argument is a new line character only
	if os.Args[1] == "\\n" {
		fmt.Println()
		return
	}

	fileName = asciiArt.BannerFile(fileName)
	// Load the banner map from the file
	bannerMap, err := asciiArt.LoadBannerMap(fileName)
	if err != nil {
		fmt.Println("error loading banner map:", err)
		return
	}

	// Process the provided argument
	word = strings.ReplaceAll(word, "\\n", "\n")
	word = strings.ReplaceAll(word, "\\t", "    ")
	lines := strings.Split(word, "\n")

	// Generate the ASCII art for each line
	for _, line := range lines {
		asciiOutput := asciiArt.PrintLineBanner(line, bannerMap)

		if *outputfile != "" {
			err = os.WriteFile(*outputfile, []byte(asciiOutput), 0o666)
			if err != nil {
				fmt.Println("error writing to file:", err)
				return
			}
		} else {
			fmt.Println(asciiOutput)
		}
	}
}

func printUsageAndExit() {
	fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
	fmt.Println() //print space between the usage message

	fmt.Println("EX: go run . --output=<fileName.txt> something standard")
	os.Exit(0)
}
