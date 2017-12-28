package shell

import (
	"log"
	"os"
	"os/exec"
)

const verboseBuilding = true

func ExecDirEnv(dir string, cmdName string, args []string, env []string) {
	log.Printf("%% %v %v %v", cmdName, args, env)
	cmd := exec.Command(cmdName, args...)
	cmd.Dir = dir
	cmd.Env = append(os.Environ(), env...)
	if verboseBuilding {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	err := cmd.Run()
	//TODO something better than crash here
	if err != nil {
		log.Panic(err)
	}
}
func ExecDir(dir string, cmdName string, args []string) {
	ExecDirEnv(dir, cmdName, args, []string{})
}

func Exec(cmdName string, args ...string) {
	ExecDir("", cmdName, args)
}
