package main

import (
	"log"
	"os"
)

const KitaBasePath = "/home/kirillvr/.kita/"
const WorldPath = KitaBasePath + "world/"

//TODO move this somewehre
func ensureDir() {
	os.MkdirAll(KitaBasePath, os.ModePerm)
	os.MkdirAll(LocalSourceDir, os.ModePerm)
	os.MkdirAll(BuildDir, os.ModePerm)
	os.MkdirAll(WorldPath, os.ModePerm)
}

func main() {
	ensureDir()
	log.Println("Kita kita tokyuu")
	p := Package{Name: "ruby"}
	p.LatestVersion().Install()
}
