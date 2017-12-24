package kita

import (
	"encoding/json"
	"github.com/kirillrdy/kita/error"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const contentFile = KitaBasePath + "content.db"

func loadGrandContentDb() map[string][]string {

	//TODO poorely named
	var gradBase = make(map[string][]string)

	file, err := os.Open(contentFile)
	error.Crash(err)
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&gradBase)
	error.Crash(err)
	return gradBase
}

//TODO currently is a bit broken doesnt do nested files matching
//TODO very waistful way of storing and loading db for now
func RegisterContent(version PackageVersion) {

	var gradBase = make(map[string][]string)

	if _, err := os.Stat(contentFile); !os.IsNotExist(err) {
		gradBase = loadGrandContentDb()
	}

	files, err := filepath.Glob(version.WorldPath() + "/**/*")
	error.Crash(err)
	for i := range files {
		files[i] = strings.TrimPrefix(files[i], version.WorldPath()+"/")
	}
	log.Println(files)

	gradBase[version.Display()] = files
	file, err := os.Create(contentFile)
	error.Crash(err)
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "\t")
	encoder.Encode(gradBase)

}

//TODO this should return latest or specific version
func VersionThatContains(requiredFile string) PackageVersion {
	grandDb := loadGrandContentDb()
	for packageVersion, files := range grandDb {
		for _, file := range files {
			if file == requiredFile {
				return FromString(packageVersion)
			}
		}
	}
	//TODO wrong
	return PackageVersion{}
}
