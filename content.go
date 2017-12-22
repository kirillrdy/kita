package kita

import (
	"github.com/kirillrdy/kita/error"
	"log"
	"path/filepath"
)

func RegisterContent(version PackageVersion) {
	files, err := filepath.Glob(version.WorldPath())
	error.Crash(err)
	log.Println(files)
}
