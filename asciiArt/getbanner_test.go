package asciiArt

import (
	"os"
	"testing"
)

func TestBannerFile(t *testing.T) {
	tests := []struct {
		args     []string
		expected string
	}{
		{[]string{"main", "standard"}, "standard.txt"},
		{[]string{"main", "shadow"}, "standard.txt"},
		{[]string{"main", "Hello", "standard"}, "standard.txt"},
		{[]string{"main", "Three", "shadow"}, "shadow.txt"},
		{[]string{"main", "Hey", "thinkertoy"}, "thinkertoy.txt"},
		{[]string{"main", "HELLO", "hollow"}, "invalid bannerfile name"},
		{[]string{"main"}, ""},
	}

	for _, tt := range tests {
		os.Args = tt.args
		result := BannerFile()
		if result != tt.expected {
			t.Errorf("BannerFile() with args %v; got %v, want %v", tt.args, result, tt.expected)
		}
	}
}
