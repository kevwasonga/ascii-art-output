package asciiArt

import "os"

func BannerFile() string {
	if len(os.Args) == 3 {
		switch os.Args[2] {
		case "standard":
			return "standard.txt"
		case "shadow":
			return "shadow.txt"
		case "thinkertoy":
			return "thinkertoy.txt"
		default:
			return "invalid bannerfile name"
		}
	}
	if len(os.Args) == 2 {
		return "standard.txt"
	}
	return ""
}
