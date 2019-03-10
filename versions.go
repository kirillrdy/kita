package kita

import (
	"errors"
	"github.com/blang/semver"
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
	var max *semver.Version
	var maxPackageVersion PackageVersion
	for _, version := range versions {
		semVersion, err := semver.Make(version.Version)
		Crash(err)
		if max == nil || semVersion.Compare(*max) > 0 {
			maxPackageVersion = version
			max = &semVersion
		}
	}
	if max != nil {
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
