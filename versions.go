package kita

import (
	"errors"
	"strings"
)

func FindVersion(versions []PackageVersion, requiredVersion string) (PackageVersion, error) {

	for _, version := range versions {
		if version.Version == requiredVersion {
			return version, nil
		}
	}
	return PackageVersion{}, errors.New("Required package version is not found")
}

func LatestVersion(versions []PackageVersion) (PackageVersion, error) {
	//TODO something better here
	if len(versions) != 0 {
		return versions[len(versions)-1], nil
	}
	return PackageVersion{}, errors.New("Can't find latest version")
}

func ldFlags(versions []PackageVersion) string {
	var flags []string
	for _, version := range versions {
		flags = append(flags, "-L"+version.LibPath())
	}
	return strings.Join(flags, " ")
}

func ldLibraryPath(versions []PackageVersion) string {
	var flags []string
	for _, version := range versions {
		flags = append(flags, version.LibPath())
	}
	return strings.Join(flags, ":")
}

func cppFlags(versions []PackageVersion) string {
	var flags []string
	for _, version := range versions {
		flags = append(flags, "-I"+version.IncludePath())
	}
	return strings.Join(flags, " ")
}
