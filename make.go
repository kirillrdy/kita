package kita

import (
	"fmt"
	"github.com/kirillrdy/kita/shell"
	"runtime"
)

func Make(source PackageSource, env []string) {
	what := source.BuildPath()

	configure(source, env)

	makeBin := "make"
	if runtime.GOOS == "freebsd" {
		makeBin = "gmake"
	}

	shell.ExecDirEnv(what, makeBin, []string{}, env)
	shell.ExecDirEnv(what, makeBin, []string{"install"}, env)
}

func configure(source PackageSource, env []string) {
	what := source.BuildPath()
	toWhere := source.PackageVersion.WorldPath()

	prefixArg := fmt.Sprintf("--prefix=%s", toWhere)
	args := []string{"configure", prefixArg}

	if source.PackageVersion.Package.Name == "gcc" {
		args = append(args, "--disable-multilib")
	}

	shell.ExecDirEnv(what, "sh", args, env)
}
