package kita

import (
	"errors"
)

func FindVersion(versions []PackageVersion, requiredVersion string) (PackageVersion, error) {

	for _, version := range versions {
		if version.Version == requiredVersion {
			return version, nil
		}
	}
	return PackageVersion{}, errors.New("Required package version is not found")
}

func LatestVersion(versions []PackageVersion) PackageVersion {
	//TODO something better here
	if len(versions) != 0 {
		return versions[len(versions)-1]
	}
	panic("Dont know how to build this yet")
}
