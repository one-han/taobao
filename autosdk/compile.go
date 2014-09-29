package autosdk

import (
	"log"
	"os/exec"
)

func Format(path string) error {
	cmd := exec.Command("gofmt", "-w", path)
	return cmd.Run()
}

func Install(paths ...string) error {
	for _, path := range paths {
		log.Println(path)
		cmd := exec.Command("go", "install", path)
		if err := cmd.Run(); err != nil {
			return err
		}
	}
	return nil
}
func Compile() {

}
