package asciiArt

import (
	"fmt"
	"os"
	"strings"
)

// Print the banner for a line of text
func PrintLineBanner(line string, bannerMap map[int][]string) string {
	if line == "" {
		fmt.Println()
		os.Exit(0)
	}

	output := make([]string, 8)

	for _, char := range line {
		banner, exists := bannerMap[int(char)]
		if !exists {
			fmt.Printf("Character %c not found in banner map\n", char)
			continue
		}

		for i := 0; i < 8; i++ {
			output[i] += banner[i]
		}
	}
	return strings.Join(output, "\n")

	}
