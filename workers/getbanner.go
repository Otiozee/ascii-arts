package workers

import "log"

func GetBannerPath(name string) string {
	banners := map[string]string{
		"standard":   "workers/banners/standard.txt",
		"shadow":     "workers/banners/shadow.txt",
		"thinkertoy": "workers/banners/thinkertoy.txt",
	}

	path, ok := banners[name]
	if !ok {
		log.Fatal("Invalid banner. Use: standard | shadow | thinkertoy")
	}
	return path
}
