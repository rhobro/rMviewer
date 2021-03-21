package main

import (
	"github.com/rhobro/goutils/pkg/fileio"
	"io"
	"os"
	"os/exec"
	"strings"
)

func init() {
	fileio.Init("", "rmviewer")
}

func main() {
	defer os.RemoveAll(fileio.TmpDir)

	// check python
	if !strings.Contains(Exec("python3", "--version"), "Python") {

		// check brew
		if !strings.Contains(Exec("brew", "--version"), "Homebrew") {
			Exec("/bin/bash", "-c", "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)")
		}

		// install python
		Exec("brew", "install", "python")
	}
}

func Exec(cmd string, args ...string) string {
	run := exec.Command(cmd, args...)
	stdout, _ := run.StdoutPipe()
	_ = run.Start()
	str := &strings.Builder{}
	_, _ = io.Copy(str, stdout)
	return str.String()
}
