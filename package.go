package main

// Package represets something like "ruby"
type Package struct {
	Name string
}

func (p Package) LatestVersion() PackageVersion {
	//TODO somehow get latest version
	if p.Name == "ruby" {
		return PackageVersion{Package: p, Version: "2.4.2"}
	}
	panic("Dont know how to build this yet")

}

// Very naive, trying to install latest version
func (p Package) Install() {
	p.LatestVersion().Install()
}
