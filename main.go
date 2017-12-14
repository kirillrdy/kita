package main

import (
	"log"
	"os"
)

const KitaBasePath = "/home/kirillvr/.kita/"

//TODO also move to util
func ensureDir() {
	os.MkdirAll(KitaBasePath, os.ModePerm)
	os.MkdirAll(LocalSourceDir, os.ModePerm)
	os.MkdirAll(BuildDir, os.ModePerm)
}

func main() {
	ensureDir()
	log.Println("Hello from go")
	source := PackageSource{fileName: "ruby-2.4.2.tar.gz"}
	source.Fetch()
	source.Extract()
}
