package shell

import (
	"github.com/kirillrdy/kita/error"
	"log"
	"os/exec"
)

func ExecDir(dir string, cmdName string, args ...string) {
	log.Printf("%% %v %v", cmdName, args)
	cmd := exec.Command(cmdName, args...)
	cmd.Dir = dir
	err := cmd.Run()
	//TODO something better than crash here
	error.Crash(err)
}

func Exec(cmdName string, args ...string) {
	ExecDir("", cmdName, args...)
}
