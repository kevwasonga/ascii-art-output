package asciiArt

import (
	"fmt"
	"os"
	"strings"
)

// PrintOutput handles printing to the console or writing to a file.
// It creates the file if it does not exist and writes the content to it.
func PrintOutput(outputFile string, content string) error {
	//if outputFile == "" {
		// Print to console
		fmt.Print(content)
	//} else {
		// Write to file
		err := os.WriteFile(outputFile, []byte(content), 0o644)
		if err != nil {
			return fmt.Errorf("error writing to file: %w", err)
		}
	//}
	return nil
}

// GetBannerLine generates an ASCII art line for the given text string using the banner map.
func GetBannerLine(line string, bannerMap map[int][]string) string {
	if len(line) == 0 {
		return ""
	}

	// Initialize a builder to construct the ASCII art line
	var result strings.Builder

	// Iterate over each character in the line
	for i := 0; i < len(line); i++ {
		char := int(line[i]) // Convert character to its integer code
		// Check if the character is present in the banner map
		if artLines, ok := bannerMap[char]; ok {
			// Append each line of the ASCII art for the character
			for j := range artLines {
				if result.Len() > 0 && i > 0 {
					result.WriteString(" ")
				}
				// If the result already has content, add space before new character's art
				result.WriteString(artLines[j])
			}
		} else {
			// If character is not in the map, add spaces or handle accordingly
			if len(bannerMap) > 0 {
				// Default handling (e.g., using space character)
				for j := range bannerMap[' '] {
					if result.Len() > 0 && i > 0 {
						result.WriteString(" ")
					}
					result.WriteString(strings.Repeat(" ", len(bannerMap[' '][j])))
				}
			}
		}
	}

	return result.String()
}
