package asciiArt

import (
	"bufio"
	"fmt"
	"os"
)

// LoadBannerMap loads the banner map from the file provided
func LoadBannerMap(fileName string) (map[int][]string, error) {
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		fmt.Printf("Error getting file info: %q\n%v\n", fileName, err)
		os.Exit(1)
	}
	//  Checks if the file has been tampered with
	fileSize := fileInfo.Size()

	switch fileName {
	case "standard.txt":
		if fileSize != 6623 {
			fmt.Println("Standard filesize tampered")
			os.Exit(1)
		}
	case "shadow.txt":
		if fileSize != 7463 {
			fmt.Println("Shadow filesize tampered")
			os.Exit(1)
		}
	case "thinkertoy.txt":
		if fileSize != 5558 {
			fmt.Println("Thinkertoy filesize tampered")
			os.Exit(1)
		}

	}
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	bannerMap := make(map[int][]string)
	key := 32
	lineCount := 0
	chunk := []string{}

	for scanner.Scan() {
		lines := scanner.Text()

		if lines != "" {
			chunk = append(chunk, lines)
			lineCount++
		}

		if lineCount == 8 {
			bannerMap[key] = chunk
			key++
			chunk = []string{}
			lineCount = 0
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	return bannerMap, nil
}
