package asciiArt

import (
	"testing"
)

func TestLoadBannerMap(t *testing.T) {
	// Test with a valid file
	bannerMap, err := LoadBannerMap("../standard.txt")
	if err != nil {
		t.Errorf("Error loading banner map: %v", err)
	}

	// Check if the banners are correctly loaded

	if bannerMap[32][0] != "      " { 
		t.Errorf("Expected banner for empty space to be '      ', got %s", bannerMap[32][0])
	}
	if bannerMap[33][0] != " _  " { 
		t.Errorf("Expected banner for exclamation mark to be '   ', got %s", bannerMap[33][0])
	}

	// Test with an invalid file
	_, err = LoadBannerMap("shaddy.doc")
	if err == nil {
		t.Errorf("Expected error loading invalid file, got nil")
	}

	// Test with a file that does not exist
	_, err = LoadBannerMap("shaddy.txt")
	if err == nil {
		t.Errorf("Expected error loading non-existent file, got nil")
	}
}
