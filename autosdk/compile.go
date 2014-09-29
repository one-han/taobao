package autosdk

import (
	"os/exec"
)

func Format(path string) error {
	cmd := exec.Command("go fmt", path)
	return cmd.Run()
}

func Install(path string) error {
	cmd := exec.Command("go install", path)
	return cmd.Run()
}
func Compile() {

}
