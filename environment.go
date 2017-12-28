package kita

import (
	"os"
)

const environmentsPath = KitaBasePath + "envs/"

type Environment struct {
	name string
}

func CreateNewEnvironment(name string) Environment {
	os.MkdirAll(environmentsPath+name, os.ModePerm)
	return Environment{name: name}
}

func (environment Environment) Path() string {
	return environmentsPath + environment.name
}

func (environment Environment) Require(path string) {
	versions := versionsThatContains(path)
	p := LatestVersion(versions)
	environment.AddPackage(p)
}

//TODO rego just doing env overwrites
func (environment Environment) AddPackage(p PackageVersion) {
	archive := PackageArchive{PackageVersion: p}
	p.Install()
	archive.ExtractTo(environment.Path())
}
