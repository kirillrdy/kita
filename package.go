package kita

import (
	"log"
)

// Package represets something like "ruby"
type Package struct {
	Name string
}

func (p Package) LatestVersion() (PackageVersion, error) {
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
	version, err := p.LatestVersion()

	if err != nil {
		log.Panicf("Failed to find latest version for %v", p)
	}

	if requiredVersion != "" {
		log.Printf("Required to install %v", requiredVersion)
		var err error
		version, err = p.FindVersion(requiredVersion)
		Crash(err)

	}
	version.Install()

}
