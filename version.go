package kita

import (
	"strings"
)

type Version struct {
	raw string
}

//TODO very hacky for now
func (version Version) Compare(version2 Version) int {
	return strings.Compare(version.raw, version2.raw)
}
