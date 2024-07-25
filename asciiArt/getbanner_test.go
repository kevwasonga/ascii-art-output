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
		{[]string{"main", "standard"}, "banner/standard.txt"},
		{[]string{"main", "shadow"}, "banner/standard.txt"},
		{[]string{"main", "Hello", "standard"}, "banner/standard.txt"},
		{[]string{"main", "Three", "shadow"}, "banner/shadow.txt"},
		{[]string{"main", "Hey", "thinkertoy"}, "banner/thinkertoy.txt"},
		{[]string{"main", "HELLO", "hollow"}, "invalid bannerfile name"},
		{[]string{"main"}, ""},
	}

	for _, tt := range tests {
		os.Args = tt.args
		result := BannerFile(tt.args[len(tt.args)-1])
		if result != tt.expected {
			t.Errorf("BannerFile() with args %v; got %v, want %v", tt.args, result, tt.expected)
		}
	}
}
