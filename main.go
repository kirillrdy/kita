package main

import (
	"log"
	"os"
)

//TODO check if we can hardcode user name in the prefix
//TODO detect user home dir
const KitaBasePath = "/home/kirillvr/.kita/"
const WorldPath = KitaBasePath + "world/"

//TODO move this somewehre
func ensureDir() {
	os.MkdirAll(KitaBasePath, os.ModePerm)
	os.MkdirAll(LocalSourceDir, os.ModePerm)
	os.MkdirAll(BuildPath, os.ModePerm)
	os.MkdirAll(WorldPath, os.ModePerm)
}

func main() {
	ensureDir()
	log.Println("Kita kita tokyuu")
	p := Package{Name: "ruby"}
	p.LatestVersion().Install()
}
