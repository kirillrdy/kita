package install

import (
	"fmt"
	"github.com/kirillrdy/kita/shell"
	"runtime"
)

func Make(what string, toWhere string) {

	prefixArg := fmt.Sprintf("--prefix=%s", toWhere)

	makeBin := "make"
	if runtime.GOOS == "freebsd" {
		makeBin = "gmake"
	}

	//TODO obviously not rely on tar binary
	shell.ExecDir(what, "sh", "configure", prefixArg)
	shell.ExecDir(what, makeBin)
	shell.ExecDir(what, makeBin, "install")
}
