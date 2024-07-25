package asciiArt

func BannerFile(banner string) string {
		switch banner {
		case "standard":
			return "banner/standard.txt"
		case "shadow":
			return "banner/shadow.txt"
		case "thinkertoy":
			return "banner/thinkertoy.txt"
		default:
			return ""
		}
	
}
