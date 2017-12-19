package main

import (
	"log"
	"path/filepath"
	"strings"
)

type PackageSourceUrl string

var versions map[string][]string = make(map[string][]string)
var files map[string][]string = make(map[string][]string)
var urls map[string][]string = make(map[string][]string)

func AllUrls() []string {
	return []string{"https://cache.ruby-lang.org/pub/ruby/2.4/ruby-2.4.2.tar.gz"}
}

func Versions(packageName string) []string {
	return versions[packageName]
}

func File(version PackageVersion) string {
	//TODO deal with empty slice
	return files[version.Display()][0]
}

func Url(fileName string) string {
	//TODO panic
	return urls[fileName][0]
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
	versions[packageName] = append(versions[packageName], version)

	fileName := filepath.Base(url)
	packagePlusVersion := packageName + "-" + version
	files[packagePlusVersion] = append(files[packagePlusVersion], fileName)
	urls[fileName] = append(urls[fileName], url)

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
