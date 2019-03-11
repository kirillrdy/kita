package kita

import (
	"errors"
	"strings"
)

func FindVersion(versions []PackageVersion, requiredVersion string) (PackageVersion, error) {

	for _, version := range versions {
		if version.Version.raw == requiredVersion {
			return version, nil
		}
	}
	return PackageVersion{}, errors.New("Required package version is not found")
}

func LatestVersion(versions []PackageVersion) (PackageVersion, error) {
	var maxPackageVersion PackageVersion
	empty := PackageVersion{}
	for _, version := range versions {

		if maxPackageVersion == empty || version.Version.Compare(maxPackageVersion.Version) > 0 {
			maxPackageVersion = version
		}
	}
	if maxPackageVersion != empty {
		return maxPackageVersion, nil
	} else {
		return PackageVersion{}, errors.New("Can't find latest version")
	}
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
