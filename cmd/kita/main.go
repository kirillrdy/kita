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

	p := kita.Package{Name: flag.Arg(0)}
	log.Printf("%v Versions %v\n", p.Name, kita.Versions(p.Name))
	p.Install(*version)
}
