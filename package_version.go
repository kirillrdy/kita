package kita

import (
	"fmt"
	"os"
)

// Represents a particular versionf of a particular package, eg ruby-2.4.2
type PackageVersion struct {
	Package Package
	Version string //TODO for now, need some sort of version type
}

func (packageVersion PackageVersion) Display() string {
	return fmt.Sprintf("%s-%s", packageVersion.Package.Name, packageVersion.Version)
}

func FromString(packageVersion string) PackageVersion {
	//TODO for now
	return extractVersion(packageVersion)
}

func (packageVersion PackageVersion) Install() {

	os.MkdirAll(packageVersion.WorldPath(), os.ModePerm)

	archive := PackageArchive{PackageVersion: packageVersion}
	if archive.Exists() {
		archive.Extract()
		return
	}
	source := packageVersion.Source()
	//TODO roughtly
	// fetch if needed
	// verify if possible, otherwise record
	// extract
	// install --> will be subtyped eg make, cmake, cargo etc
	if !source.ExistsLocally() {
		source.Fetch()
	}

	source.Extract()
	source.Install()

	archive.create()

	RegisterContent(packageVersion)
}

func (packageVersion PackageVersion) WorldPath() string {
	return fmt.Sprintf("%s%s", WorldPath, packageVersion.Display())
}

func (packageVersion PackageVersion) Source() PackageSource {
	fileName := File(packageVersion)
	return PackageSource{fileName: fileName, PackageVersion: packageVersion}
}

//TODO this is wrong, since its very C specific
func (version PackageVersion) LibPath() string {
	return version.WorldPath() + "/lib"
}

func (version PackageVersion) IncludePath() string {
	return version.WorldPath() + "/include"
}

func (version PackageVersion) BinPath() string {
	return version.WorldPath() + "/bin"
}
