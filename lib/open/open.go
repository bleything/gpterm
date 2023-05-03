package open

import (
	"os/exec"
	"runtime"
)

func Browser(url string) *exec.Cmd {
	switch os := runtime.GOOS; os {
	case "darwin":
		return exec.Command("/usr/bin/open", url)
	default:
		return exec.Command("xdg-open", url)
	}
}
