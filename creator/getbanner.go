package creator

import "log"

// GetBannerPath create a map of banner names (standard, shadow, thinkertoy)
// and link them to their corresponding banner files.
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
