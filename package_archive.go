package kita

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
	shell.ExecDir(archive.PackageVersion.WorldPath(), "tar", []string{"cvf", archive.path(), "."})
}

func (archive PackageArchive) Extract() {
	archive.ExtractTo(archive.PackageVersion.WorldPath())
}

func (archive PackageArchive) ExtractTo(path string) {
	//TODO obviously not rely on tar binary
	shell.Exec("tar", "xvf", archive.path(), "-C", path)
}

func (archive PackageArchive) Exists() bool {
	_, err := os.Stat(archive.path())
	return !os.IsNotExist(err)
}
