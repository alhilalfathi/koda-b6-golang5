package clearscreen

import (
	"os"
	"os/exec"
	"runtime"
)

func ClearScreen() {
	if runtime.GOOS == "windows" {
		// Windows uses the "cls" command, executed via "cmd /c".
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		// Linux and macOS use the "clear" command.
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
