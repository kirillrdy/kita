package main

// Package represets something like "ruby"
type Package struct {
	Name string
}

func (p Package) LatestVersion() PackageVersion {
	//TODO somehow get latest version
	return PackageVersion{Package: p, Version: "2.4.2"}
}