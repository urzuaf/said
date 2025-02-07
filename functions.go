package main

import (
	"fmt"
	"os"
)

func processInstruction(instruction string, args string) {
	switch instruction {
	case "search":
		searchCommands(args)
	case "list":
		listCommands()
	default:
		return
	}
}

// Search for keywords in commands
func searchCommands(args string) {
	if len(args) <= 0 {
		fmt.Printf("No hay nada para buscar")
		return
	}

	results := smartSearch(args)
	if len(results) > 0 {
		fmt.Println("Coincidencias encontradas:")
		for _, result := range results {
			fmt.Println(result)
		}
	} else {
		fmt.Println("No se encontraron comandos relacionados con:", args)
	}

	os.Exit(0)
}

// List all commands
func listCommands() {
	//Get all the keys from the commandDescriptions
	keys := make([]string, 0, len(commandDescriptions))
	for key := range commandDescriptions {
		keys = append(keys, key)
	}

	List(keys)
	os.Exit(0)
}
