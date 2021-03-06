package kita

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

type PackageSourceUrl string

var versions = make(map[Package][]PackageVersion)
var files = make(map[string][]string)
var urls = make(map[string][]string)

const UrlsFileName = "kita.urls"

func loadUrlsFromFile() {

	start := time.Now()

	file, err := os.Open(KitaBasePath + UrlsFileName)
	Crash(err)

	defer func() {
		err := file.Close()
		Crash(err)
	}()

	scanner := bufio.NewScanner(file)

	count := 0

	for scanner.Scan() {
		addUrl(scanner.Text())
		count += 1
	}
	log.Printf("Processed %d urls in %v", count, time.Since(start))
}

func Versions(p Package) []PackageVersion {
	return versions[p]
}

func File(version PackageVersion) string {
	//TODO deal with empty slice
	return files[version.Display()][0]
}

func Url(fileName string) string {
	//TODO panic
	return urls[fileName][0]
}

func extractVersion(url string) PackageVersion {
	for _, regex := range siteRegexes {
		if regex.MatchString(url) {
			match := regex.FindStringSubmatch(url)
			return PackageVersion{Package: Package{Name: match[1]}, Version: Version{raw: match[2]}}
		}
	}

	base := filepath.Base(url)
	base = stripThings(base)
	split := strings.Split(base, "-")
	if len(split) == 0 {
		log.Panic("failed to guess version")
	}
	version := split[len(split)-1]
	packageName := strings.TrimSuffix(base, "-"+version)
	return PackageVersion{Package: Package{Name: packageName}, Version: Version{raw: version}}
}

var siteRegexes = []*regexp.Regexp{
	regexp.MustCompile("https://github.com/(?:.*)/(.*)/archive/v(.*).tar.gz"),
	regexp.MustCompile("http://erlang.org/(?:.*)/(.*)_src_(.*).tar.gz"),
}

func addUrl(url string) {
	version := extractVersion(url)
	versions[version.Package] = append(versions[version.Package], version)

	fileName := filepath.Base(url)
	files[version.Display()] = append(files[version.Display()], fileName)
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
