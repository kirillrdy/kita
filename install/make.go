package install

import (
	"fmt"
	"github.com/kirillrdy/kita/shell"
	"runtime"
)

func Make(what string, toWhere string) {

	confiure(what, toWhere)

	makeBin := "make"
	if runtime.GOOS == "freebsd" {
		makeBin = "gmake"
	}

	shell.ExecDir(what, makeBin)
	shell.ExecDir(what, makeBin, "install")
}

func confiure(what string, toWhere string) {
	prefixArg := fmt.Sprintf("--prefix=%s", toWhere)
	args := []string{"configure", prefixArg}
	shell.ExecDir(what, "sh", args...)
}
