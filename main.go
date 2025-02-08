package main

import (
	"fmt"
	"os"
)

const saidVersion = "v0.3.0"

func main() {
	//Read the configuration files
	configure()
	config := configurationObject.getConfiguration()

	// Check for optionals instructions like search and list
	command := processInstruction(config.instruction, config.arguments)

	if !isValidCommand(command) {
		fmt.Println("Comando no reconocido:", command)
		os.Exit(1)
	}

	// PROCESS FLAGS BEFORE EXECUTION

	if config.version {
		displayVersion()
		os.Exit(0)
	}

	filepath := config.dataFolder + config.language + "/" + command

	if !FileExists(filepath) {
		fmt.Println("No existe el archivo en la configuración solicitada")
		fmt.Println(config.dataFolder + config.language + "/" + command)
		os.Exit(1)
	}

	var filedata string

	// PROCESS OUTPUT FLAGS

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
