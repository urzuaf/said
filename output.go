package main

import (
	"fmt"
	"os"
	"os/exec"
	"text/tabwriter"
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

func List(list []string) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', 0)

	for i, item := range list {
		fmt.Fprintf(w, item+"\t")
		// we break lines every 3 items
		if (i+1)%3 == 0 {
			fmt.Fprintf(w, "\n")
		}
	}
	fmt.Fprintf(w, "\n")
	w.Flush()
}
