package kita

import (
	"log"
)

// Package represets something like "ruby"
type Package struct {
	Name string
}

func (p Package) LatestVersion() PackageVersion {

	versons := p.Versions()
	//TODO something better here
	if len(versons) != 0 {
		return versons[len(versons)-1]
	}
	panic("Dont know how to build this yet")

}

func (p Package) Versions() []PackageVersion {
	return Versions(p)
}

func (p Package) FindVersion(requiredVersion string) (PackageVersion, error) {
	return FindVersion(Versions(p), requiredVersion)
}

// Very naive, trying to install latest version
func (p Package) Install(requiredVersion string) {
	version := p.LatestVersion()

	if requiredVersion != "" {
		log.Printf("Required to install %v", requiredVersion)
		var err error
		version, err = p.FindVersion(requiredVersion)
		Crash(err)

	}
	version.Install()

}
