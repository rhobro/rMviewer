package main

import (
	"fmt"
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

	// check rmview
	if strings.Contains(Exec("where", "rmview"), "not found") {
		// check python
		fmt.Println("Checking for a Python 3 installation")
		if !strings.Contains(Exec("python3", "--version"), "Python") {

			// check brew
			fmt.Println("Checking for a Homebrew installation")
			if !strings.Contains(Exec("brew", "--version"), "Homebrew") {
				fmt.Println("Installing Homebrew")
				Exec("/bin/bash", "-c", "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)")
			}

			// install python
			fmt.Println("Installing Python 3")
			Exec("brew", "install", "python")
		}

		fmt.Println("Checking for a Git installation")
		if !strings.Contains(Exec("git", "--version"), "version") {
			fmt.Println("Installing git")
			Exec("brew", "install", "git")
		}

		fmt.Println("Downloading rMview")
		Exec("git", "clone", "https://github.com/bordaigorl/rmview.git", fileio.TmpDir)
		fmt.Println("Installing rMview")
		Exec("pip3", "install", fileio.TmpDir)
	}

	// launch rmview
	fmt.Println("Launching rMview")
	Exec("rmview")
}

func Exec(cmd string, args ...string) string {
	run := exec.Command(cmd, args...)
	stdout, _ := run.StdoutPipe()
	_ = run.Start()
	str := &strings.Builder{}
	_, _ = io.Copy(str, stdout)
	return str.String()
}
