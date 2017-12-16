package shell

import (
	"github.com/kirillrdy/kita/error"
	"os/exec"
)

func ExecDir(dir string, cmdName string, args ...string) {
	cmd := exec.Command(cmdName, args...)
	cmd.Dir = dir
	err := cmd.Run()
	error.Crash(err)
}

func Exec(cmdName string, args ...string) {
	ExecDir("", cmdName, args...)
}
