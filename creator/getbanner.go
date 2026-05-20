package creator

import "log"

func GetBannerPath(name string) string {
	banners := map[string]string{
		"standard":   "creator/banners/standard.txt",
		"shadow":     "creator/banners/shadow.txt",
		"thinkertoy": "creator/banners/thinkertoy.txt",
	}

	path, ok := banners[name]
	if !ok {
		log.Fatal("Invalid banner. Use: standard | shadow | thinkertoy")
	}
	return path
}
