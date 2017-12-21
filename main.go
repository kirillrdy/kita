package main

import (
	"flag"
	"log"
	"os"
)

const KitaBasePath = "/kita/"
const WorldPath = KitaBasePath + "world/"

//TODO move this somewehre
func ensureDir() {
	os.MkdirAll(KitaBasePath, os.ModePerm)
	os.MkdirAll(LocalSourcesPath, os.ModePerm)
	os.MkdirAll(BuildPath, os.ModePerm)
	os.MkdirAll(WorldPath, os.ModePerm)
	os.MkdirAll(LocalPackageArchivesPath, os.ModePerm)
}

func main() {
	log.Println("Kita kita tokyuu")
	ensureDir()

	version := flag.String("version", "", "Version of package to install")

	flag.Parse()

	p := Package{Name: flag.Arg(0)}
	log.Printf("%v Versions %v\n", p.Name, versions[p.Name])
	p.Install(*version)
}
