package main

import (
	"flag"
	"github.com/kirillrdy/kita"
	"log"
)

func main() {
	log.Println("Kita kita tokyuu")

	version := flag.String("version", "", "Version of package to install")

	flag.Parse()

	command := flag.Arg(0)
	if command == "install" || command == "i" {
		p := kita.Package{Name: flag.Arg(1)}
		log.Printf("%v Versions %v\n", p.Name, kita.Versions(p))
		p.Install(*version)
	}

	if command == "init" {

	}
}
