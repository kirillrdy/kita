package main

import (
	"fmt"
	"github.com/kirillrdy/kita/shell"
	"log"
	"os"
)

//TODO different OS and architectures
type PackageArchive struct {
	PackageVersion PackageVersion
}

const LocalPackageArchivesPath = KitaBasePath + "packages/"

func (archive PackageArchive) fileName() string {
	return fmt.Sprintf("%s.tar.xz", archive.PackageVersion.Display())
}

func (archive PackageArchive) path() string {
	return LocalPackageArchivesPath + archive.fileName()
}

func (archive PackageArchive) create() {
	log.Printf("Making Package for %v", archive.PackageVersion.Display())
	//TODO usage of Display here
	//TODO correct flags for compression and compressort type
	shell.ExecDir(WorldPath, "tar", "cvf", archive.path(), archive.PackageVersion.Display())
}

func (archive PackageArchive) Extract() {
	//TODO obviously not rely on tar binary
	shell.Exec("tar", "xvf", archive.path(), "-C", WorldPath)
}

func (archive PackageArchive) Exists() bool {
	_, err := os.Stat(archive.path())
	return !os.IsNotExist(err)
}
