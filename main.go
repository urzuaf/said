package main

import (
	"fmt"
	"os"
)

func main() {
	//Read the configuration files
	config := configurationObject.getConfiguration()

	// if user inputs "search <keyword>" look for related commands
	if config.instruction == "search" && len(config.arguments) > 1 {
		results := smartSearch(config.arguments)
		if len(results) > 0 {
			fmt.Println("Coincidencias encontradas:")
			for _, result := range results {
				fmt.Println(result)
			}
		} else {
			fmt.Println("No se encontraron comandos relacionados con:", config.arguments)
		}
		return
	}

	if !isValidCommand(config.instruction) {
		fmt.Println("Comando no reconocido: ", config.instruction)
		os.Exit(1)
	}

	filepath := config.dataFolder + config.language + "/" + config.instruction

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
