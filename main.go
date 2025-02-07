package main

import (
	"fmt"
	"os"
)

func main() {
	//Read the configuration files
	configure()
	config := configurationObject.getConfiguration()

	// Check for optionals instructions like search and list
	processInstruction(config.instruction, config.arguments)

	if !isValidCommand(config.instruction) {
		fmt.Println("Comando no reconocido:", config.instruction)
		os.Exit(1)
	}

	filepath := config.dataFolder + config.language + "/" + config.instruction

	if !FileExists(filepath) {
		fmt.Println("No existe el archivo en la configuración solicitada")
		os.Exit(1)
	}

	var filedata string

	//know if the user wants only examples
	if config.quickExamples {
		filedata = readExamples(filepath)
	} else {
		filedata = readWholeFile(filepath)
	}

	//knows if the user wants to print in standard output
	if config.lessOutput {
		notStdoutReader(filedata)
	} else {
		stdoutReader(filedata)
	}
}
