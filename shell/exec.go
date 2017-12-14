package shell

import (
	"github.com/kirillrdy/kita/error"
	"os/exec"
)

func Exec(cmdName string, args ...string) {
	cmd := exec.Command(cmdName, args...)
	err := cmd.Run()
	error.Crash(err)
}
