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
