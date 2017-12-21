package main

import (
	"errors"
	e "github.com/kirillrdy/kita/error"
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
	var versions []PackageVersion
	for _, version := range Versions(p.Name) {
		versions = append(versions, PackageVersion{Package: p, Version: version})
	}
	return versions
}

func (p Package) FindVersion(requiredVersion string) (PackageVersion, error) {

	for _, version := range Versions(p.Name) {
		if version == requiredVersion {
			return PackageVersion{Package: p, Version: version}, nil
		}
	}
	return PackageVersion{}, errors.New("Required package version is not found")
}

// Very naive, trying to install latest version
func (p Package) Install(requiredVersion string) {
	log.Printf("Required to install %v", requiredVersion)
	version := p.LatestVersion()

	if requiredVersion != "" {
		var err error
		version, err = p.FindVersion(requiredVersion)
		e.Crash(err)

	}
	version.Install()

}
