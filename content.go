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

//TODO currently is a bit broken doesnt do nested files matching
//TODO very waistful way of storing and loading db for now
func RegisterContent(version PackageVersion) {

	var gradBase = make(map[string][]string)

	if _, err := os.Stat(contentFile); !os.IsNotExist(err) {

		file, err := os.Open(contentFile)
		error.Crash(err)
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&gradBase)
		error.Crash(err)
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
