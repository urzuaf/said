package main

import (
	"os"
	"os/exec"
)

func stdoutReader(text string) {
	os.Stdout.WriteString(text + "\n")
}

func notStdoutReader(text string) {
	tempFile, err := os.CreateTemp("", "help-*.tmp")
	Check(err)

	defer os.Remove(tempFile.Name())

	_, err = tempFile.WriteString(text)
	Check(err)

	tempFile.Close()

	cmd := exec.Command("less", tempFile.Name())
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	Check(err)
}
