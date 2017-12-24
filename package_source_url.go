package kita

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type PackageSourceUrl string

var versions map[string][]string = make(map[string][]string)
var files map[string][]string = make(map[string][]string)
var urls map[string][]string = make(map[string][]string)

const UrlsFileName = "kita.urls"

func loadUrlsFromFile() {

	start := time.Now()

	file, err := os.Open(KitaBasePath + UrlsFileName)
	defer file.Close()
	Crash(err)

	scanner := bufio.NewScanner(file)

	count := 0

	for scanner.Scan() {
		addUrl(scanner.Text())
		count += 1
	}
	log.Printf("Processed %d urls in %v", count, time.Since(start))
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
	endings := []string{
		".tar.gz",
		".source.tar.xz", //Hack for firefox
		".tar.xz",
		".zip",
		".tgz",
		".tar.bz2"}

	var result = fileName
	for _, ending := range endings {
		result = strings.TrimSuffix(result, ending)
	}
	return result
}

func init() {
	//TODO time that this doesnt take too long so that we don't do this every start
	loadUrlsFromFile()
}
