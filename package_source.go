package kita

import (
	"fmt"
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

//TODO prevent clashes or overwrites, eg github files only have version in the name
func (source PackageSource) Fetch() {
	log.Printf("Fetching: %v", source.URL())

	destination, err := os.Create(source.LocalPath())
	defer destination.Close() //TODO errors
	Crash(err)

	response, err := http.Get(source.URL())
	Crash(err)

	_, err = io.Copy(destination, response.Body)
	Crash(err)
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

	var requiredFiles []string

	//TODO something better than this
	if source.PackageVersion.Package.Name == "gcc" {
		requiredFiles = append(requiredFiles, "lib/libmpfr.so")
		requiredFiles = append(requiredFiles, "lib/libgmp.so")
		requiredFiles = append(requiredFiles, "include/mpc.h")
	}

	if source.PackageVersion.Package.Name == "mpc" {
		requiredFiles = append(requiredFiles, "include/gmp.h")
		requiredFiles = append(requiredFiles, "lib/libmpfr.so")
	}

	if source.PackageVersion.Package.Name == "tig" {
		requiredFiles = append(requiredFiles, "bin/gcc")
	}

	if source.PackageVersion.Package.Name == "glib" {
		requiredFiles = append(requiredFiles, "lib64/libffi.so") //TODO wtf lib64
	}

	if source.PackageVersion.Package.Name == "mpfr" {
		requiredFiles = append(requiredFiles, "include/gmp.h")
	}

	//TODO this should be a set rather than a list/array/slice
	var dependecies []PackageVersion
	for _, file := range requiredFiles {
		//TODO should we just provide requred files or the whole thingy ?
		version, err := LatestVersion(versionsThatContains(file))
		if err != nil {
			log.Panicf("Could not find the dependecy for %v", file)
		}
		dependecies = append(dependecies, version)
	}

	//TODO review if those are required
	env := []string{}

	if ldFlags(dependecies) != "" {
		env = append(env, "LDFLAGS="+ldFlags(dependecies))
	}
	if cppFlags(dependecies) != "" {
		env = append(env, "CPPFLAGS="+cppFlags(dependecies))
	}

	if ldLibraryPath(dependecies) != "" {
		env = append(env, "LD_LIBRARY_PATH="+ldLibraryPath(dependecies))
		env = append(env, "LIBRARY_PATH="+ldLibraryPath(dependecies))
	}

	Make(source, env)
}
