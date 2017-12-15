package main

// Represents a particular versionf of a particular package, eg ruby-2.4.2
type PackageVersion struct {
	Package Package
	Version string //TODO for now, need some sort of version type
}

func (packageVersion PackageVersion) Install() {
	source := packageVersion.Source()
	//TODO roughtly
	// fetch if needed
	// verify if possible, otherwise record
	// extract
	// install --> will be subtyped eg make, cmake, cargo etc
	source.Fetch()
	source.Extract()
}

func (PackageVersion PackageVersion) Source() PackageSource {
	//TODO find source for a given version
	return PackageSource{fileName: "ruby-2.4.2.tar.gz"}
}
