package main

import (
	"flag"
	"fmt"
	"strings"

	"ascii/asciiArt"
)

func main() {
	// Define and parse the --output flag
	var outputFile string
	flag.StringVar(&outputFile, "output", "", "Specify the output file to save the ASCII art")

	// Parse command-line flags
	flag.Parse()
	args := flag.Args()

	// Check if the number of arguments is valid
	if len(args) < 1 || len(args) > 5 {
		fmt.Printf("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard\n")
		return
	}

	// If the first argument is empty, return
	if args[0] == "" {
		return
	}

	// Print a new line and exit in case argument is a new line character only
	if args[0] == "\\n" {
		fmt.Println()
		return
	}

	// Determine the banner file name
	fileName := asciiArt.BannerFile()
	if fileName == "invalid bannerfile name" || fileName == "invalid arguments" {
		fmt.Printf("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard\n")
		return
	}

	// Load the banner map from the file
	bannerMap, err := asciiArt.LoadBannerMap(fileName)
	if err != nil {
		fmt.Println("error loading banner map:", err)
		return
	}

	if len(args) == 4 {
		// Handle case with 4 arguments
		inputString := args[1]
		inputString = strings.ReplaceAll(inputString, "\\n", "\n")
		inputString = strings.ReplaceAll(inputString, "\\t", "    ")
		lines := strings.Split(inputString, "\n")

		// Generate the ASCII art for each line
		var result strings.Builder
		for _, line := range lines {
			lineArt := asciiArt.GetBannerLine(line, bannerMap)
			result.WriteString(lineArt + "\n")
		}

		// Handle printing or saving the output
		err = asciiArt.PrintOutput(outputFile, result.String())
		if err != nil {
			fmt.Println(err)
			return
		}
	} else if len(args) == 2 {
		// Handle cases with 2 or 3 arguments
		inputString := args[1]
		inputString = strings.ReplaceAll(inputString, "\\n", "\n")
		inputString = strings.ReplaceAll(inputString, "\\t", "    ")
		lines := strings.Split(inputString, "\n")

		// Print the ASCII art for each line
		for _, line := range lines {
			asciiArt.PrintLineBanner(line, bannerMap)
		}
	}
}
