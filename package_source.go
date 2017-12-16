package main

import (
	"fmt"
	"github.com/kirillrdy/kita/error"
	"github.com/kirillrdy/kita/shell"
	"io"
	"log"
	"net/http"
	"os"
)

// PackageSource represents something like "ruby-2.4.2.tar.gz"
type PackageSource struct {
	fileName       string //TODO  Or maybe archive type or something
	PackageVersion PackageVersion
}

const LocalSourceDir = KitaBasePath + "sources/"

//TODO move somewhere
const BuildPath = KitaBasePath + "build/"

func (packageSource PackageSource) URL() string {
	//TODO ask some sort of URLer
	return "https://cache.ruby-lang.org/pub/ruby/2.4/ruby-2.4.2.tar.gz"
}

func (packageSource PackageSource) LocalPath() string {
	return LocalSourceDir + packageSource.fileName
}

func (packageSource PackageSource) Fetch() {
	log.Printf("Fetching: %v", packageSource.URL())

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
	//TODO obviously not rely on tar binary
	shell.Exec("tar", "xvf", packageSource.LocalPath(), "-C", BuildPath)
}

func (packageSource PackageSource) BuildPath() string {
	return fmt.Sprintf("%s%s", BuildPath, packageSource.PackageVersion.Display())
}

func (packageSource PackageSource) prefixArgument() string {
	return fmt.Sprintf("--prefix=%s", packageSource.PackageVersion.WorldPath())
}

func (packageSource PackageSource) Install() {
	//TODO obviously not rely on tar binary
	shell.ExecDir(packageSource.BuildPath(), "sh", "configure", packageSource.prefixArgument())
	shell.ExecDir(packageSource.BuildPath(), "make")
	shell.ExecDir(packageSource.BuildPath(), "make", "install")
}
