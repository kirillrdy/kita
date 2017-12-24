package shell

import (
	"log"
	"os/exec"
)

func ExecDir(dir string, cmdName string, args ...string) {
	log.Printf("%% %v %v", cmdName, args)
	cmd := exec.Command(cmdName, args...)
	cmd.Dir = dir
	err := cmd.Run()
	//TODO something better than crash here
	if err != nil {
		log.Panic(err)
	}
}

func Exec(cmdName string, args ...string) {
	ExecDir("", cmdName, args...)
}
