package kita

import (
	"os"
)

const KitaBasePath = "/zroot/kita/"
const WorldPath = KitaBasePath + "world/"

//TODO move this somewehre
func ensureDir() {
	os.MkdirAll(KitaBasePath, os.ModePerm)
	os.MkdirAll(LocalSourcesPath, os.ModePerm)
	os.MkdirAll(BuildPath, os.ModePerm)
	os.MkdirAll(WorldPath, os.ModePerm)
	os.MkdirAll(LocalPackageArchivesPath, os.ModePerm)
	os.MkdirAll(environmentsPath, os.ModePerm)
}

func init() {
	ensureDir()
}
