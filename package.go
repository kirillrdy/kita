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
	var versions []PackageVersion
	for _, version := range Versions(p.Name) {
		versions = append(versions, PackageVersion{Package: p, Version: version})
	}
	return versions
}

// Very naive, trying to install latest version
func (p Package) Install() {
	p.LatestVersion().Install()
}
