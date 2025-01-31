package main

import (
	"os"
	"strings"
)

func readWholeFile(filepath string) string {
	filebytes, err := os.ReadFile(filepath)
	Check(err)
	filedata := string(filebytes)
	return filedata
}

func readExamples(filepath string) string {
	var capturing bool
	var result []string

	lines := strings.Split(readWholeFile(filepath), "\n")
	for _, line := range lines {

		if strings.HasPrefix(line, "🛠️ Uso") || strings.HasPrefix(line, "🔹 Ejemplos") {
			capturing = true
		} else if strings.HasPrefix(line, "📝 Opciones") || strings.HasPrefix(line, "ℹ️ Información") {
			capturing = false
		}

		if capturing {
			result = append(result, line)
		}
	}

	return strings.Join(result, "\n")
}
