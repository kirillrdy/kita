package kita

import (
	"fmt"
	"github.com/kirillrdy/kita/error"
	"github.com/kirillrdy/kita/install"
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

const LocalSourcesPath = KitaBasePath + "sources/"

//TODO move somewhere
const BuildPath = KitaBasePath + "build/"

func (source PackageSource) URL() string {
	return Url(source.fileName)
}

func (source PackageSource) LocalPath() string {
	return LocalSourcesPath + source.fileName
}

func (source PackageSource) Fetch() {
	log.Printf("Fetching: %v", source.URL())

	destination, err := os.Create(source.LocalPath())
	defer destination.Close() //TODO errors
	error.Crash(err)

	response, err := http.Get(source.URL())
	error.Crash(err)

	_, err = io.Copy(destination, response.Body)
	error.Crash(err)
	defer response.Body.Close()
}

func (source PackageSource) ExistsLocally() bool {
	_, err := os.Stat(source.LocalPath())
	return !os.IsNotExist(err)
}

func (source PackageSource) Extract() {
	//TODO obviously not rely on tar binary
	shell.Exec("tar", "xvf", source.LocalPath(), "-C", BuildPath)
}

//TODO firgure this out not based on Display of version
func (source PackageSource) BuildPath() string {
	return fmt.Sprintf("%s%s", BuildPath, source.PackageVersion.Display())
}

func (source PackageSource) prefixArgument() string {
	return fmt.Sprintf("--prefix=%s", source.PackageVersion.WorldPath())
}

func (source PackageSource) Install() {
	install.Make(source.BuildPath(), source.PackageVersion.WorldPath())
	//TODO obviously not rely on tar binary
}
