package main

import (
	"log"
	"path/filepath"
	"strings"
)

type PackageSourceUrl string

func AllUrls() []string {
	return []string{"https://cache.ruby-lang.org/pub/ruby/2.4/ruby-2.4.2.tar.gz"}
}

func extractVersion(url string) (string, string) {
	base := filepath.Base(url)
	base = stripThings(base)
	split := strings.Split(base, "-")
	if len(split) == 0 {
		log.Panic("failed to guess version")
	}
	version := split[len(split)-1]
	return strings.TrimSuffix(base, "-"+version), version
}

func addUrl(url string) {
	packageName, version := extractVersion(url)
	log.Println(packageName)
	log.Println(version)

}

func stripThings(fileName string) string {
	endings := []string{".tar.gz", ".tar.xz", ".tar.bz2"}

	var result = fileName
	for _, ending := range endings {
		result = strings.TrimSuffix(result, ending)
	}
	return result
}

func init() {
	for _, url := range AllUrls() {
		addUrl(url)
	}
}
