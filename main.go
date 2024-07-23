package main

import (
	"fmt"
	"os"
	"strings"

	"ascii/asciiArt"
)

func main() {
	// Check if an argument is provided
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Printf("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard\n")
		return
	}
	if os.Args[1] == "" {
		return
	}
	// print a new line and exit in case argument is a new line character only
	if os.Args[1] == "\\n" {
		fmt.Println()
		return
	}

	fileName := asciiArt.BannerFile()

	// Load the banner map from the file
	bannerMap, err := asciiArt.LoadBannerMap(fileName)
	if err != nil {
		fmt.Println("error loading banner map:", err)
		return
	}

	// Process the provided argument
	args := strings.ReplaceAll(os.Args[1], "\\n", "\n")
	args = strings.ReplaceAll(args, "\\t", "    ")
	lines := strings.Split(args, "\n")

	// Generate the ASCII art for each line
	for _, line := range lines {
		asciiArt.PrintLineBanner(line, bannerMap)
	}
}
