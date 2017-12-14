package main

import (
	"io"
	"net/http"
	"os"
)

// PackageSource represents something like "ruby-2.4.2.tar.gz"
type PackageSource struct {
	fileName string //TODO  Or maybe archive type or something
}

const LocalSourceDir = KitaBasePath + "sources/"

func (packageSource PackageSource) URL() string {
	//TODO ask some sort of URLer
	return "https://cache.ruby-lang.org/pub/ruby/2.4/ruby-2.4.2.tar.gz"
}

func (packageSource PackageSource) LocalPath() string {
	return LocalSourceDir + packageSource.fileName
}

func (packageSource PackageSource) Fetch() {
	destination, err := os.Create(packageSource.LocalPath())
	Crash(err)

	response, err := http.Get(packageSource.URL())
	Crash(err)

	_, err = io.Copy(destination, response.Body)
	Crash(err)
	defer response.Body.Close()
}
