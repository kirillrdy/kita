package main

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
	if p.Name == "ruby" {
		return []PackageVersion{{Package: p, Version: "2.4.2"}}
	}
	return []PackageVersion{}
}

// Very naive, trying to install latest version
func (p Package) Install() {
	p.LatestVersion().Install()
}
