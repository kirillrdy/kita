package main

import (
	"github.com/kirillrdy/kita/error"
	"github.com/kirillrdy/kita/shell"
	"io"
	"net/http"
	"os"
)

// PackageSource represents something like "ruby-2.4.2.tar.gz"
type PackageSource struct {
	fileName string //TODO  Or maybe archive type or something
}

const LocalSourceDir = KitaBasePath + "sources/"

//TODO move somewhere
const BuildDir = KitaBasePath + "build/"

func (packageSource PackageSource) URL() string {
	//TODO ask some sort of URLer
	return "https://cache.ruby-lang.org/pub/ruby/2.4/ruby-2.4.2.tar.gz"
}

func (packageSource PackageSource) LocalPath() string {
	return LocalSourceDir + packageSource.fileName
}

func (packageSource PackageSource) Fetch() {
	destination, err := os.Create(packageSource.LocalPath())
	defer destination.Close() //TODO errors
	error.Crash(err)

	response, err := http.Get(packageSource.URL())
	error.Crash(err)

	_, err = io.Copy(destination, response.Body)
	error.Crash(err)
	defer response.Body.Close()
}
func (packageSource PackageSource) Extract() {
	shell.Exec("tar", "xvf", packageSource.LocalPath(), "-C", BuildDir)
}
