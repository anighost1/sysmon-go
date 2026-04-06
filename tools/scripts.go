package main

import (
	"fmt"
	"os"
	"os/exec"
)

func runCmd(name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run scripts.go [run|build]")
		return
	}

	switch os.Args[1] {

	case "run":
		runCmd("go", "run", ".")

	case "build":
		runCmd("go", "build", "-o", "sysmon.exe")

	case "clean":
		os.Remove("sysmon.exe")
		fmt.Println("Cleaned")

	default:
		fmt.Println("Unknown command")
	}
}
