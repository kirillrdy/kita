package main

import (
	"log"
	"os"
)

const KitaBasePath = "/home/kirillvr/.kita/"

//TODO move to util
func Crash(err error) {
	if err != nil {
		log.Panic(err)
	}
}

//TODO also move to util
func ensureDir() {
	os.MkdirAll(KitaBasePath, os.ModePerm)
	os.MkdirAll(LocalSourceDir, os.ModePerm)
}

func main() {
	ensureDir()
	log.Println("Hello from go")
	source := PackageSource{fileName: "ruby-2.4.2.tar.gz"}
	source.Fetch()
}
