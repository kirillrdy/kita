package kita

import (
	"log"
)

// Package represets something like "ruby"
type Package struct {
	Name string
}

func (p Package) LatestVersion() PackageVersion {
	return LatestVersion(p.Versions())
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
