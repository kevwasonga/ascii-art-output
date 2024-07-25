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
	outputfile := flag.String("output", "", "Specify the output file in the format --output=<fileName.txt>")
	flag.Parse()

	// Check if the output flag is provided
	if *outputfile != "" {
		// Validate the output flag format
		if !isValidOutputFlag(os.Args) {
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

func isValidOutputFlag(args []string) bool {
	// Ensure the flag starts with "--output="
	if len(args) < 1 || !strings.HasPrefix(args[1], "--output=") {
		return false
	}

	// Extract the flag value
	flagValue := args[1][len("--output="):]

	// Validate the flag format using regex
	validOutput := regexp.MustCompile(`^[a-zA-Z0-9_-]+\.txt$`)
	return validOutput.MatchString(flagValue)
}

func printUsageAndExit() {
	fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
	fmt.Println()
	fmt.Println("EX: go run . --output=<fileName.txt> something standard")
	os.Exit(0)
}
