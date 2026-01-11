package cli

import (
	"fmt"
	"os"
	"os/exec"
)

func Preprocess(filename string, includeDir string) []byte {
	fmt.Println("[BBUSS] preprocess and compiler step...")

	cmd := exec.Command("gcc", "-E", "-P", "-x", "c", "-I"+includeDir, filename)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("[ERR] failed to preprocess with GCC")
		fmt.Println(string(out))
		os.Exit(1)
		return nil
	}
	fmt.Println(string(out))

	return out
}


