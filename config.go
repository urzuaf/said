package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	quickExamples bool
	lessOutput    bool
	dataFolder    string
	language      string
	instruction   string
	arguments     string
}

var configurationObject Config

func (c Config) getConfiguration() Config {
	return c
}

// configuration of the program

func configure() {

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error al obtener el directorio home:", err)
		return
	}
	// Construir la ruta completa al archivo de configuraciones
	filePath := homeDir + "/.said/said.conf"

	flags := parseFlags()
	configfile := loadConfigFiles(filePath)
	instruction := getCommand()

	configurationObject = Config{
		quickExamples: flags["e"],
		lessOutput:    flags["l"],
		dataFolder:    homeDir + configfile["datafolder"],
		language:      configfile["language"],
		instruction:   instruction["instruction"],
		arguments:     instruction["arguments"],
	}
}

func getCommand() map[string]string {
	//default command its program itself
	instruction := "said"
	arguments := ""

	args := flag.Args()
	if len(args) > 0 {
		instruction = args[0]
	}
	if len(args) > 1 {
		arguments = strings.Join(args[1:], " ")
	}

	var command = map[string]string{
		"instruction": instruction,
		"arguments":   arguments,
	}

	return command
}

func parseFlags() map[string]bool {
	qEx := flag.Bool("e", false, "return only examples")
	lOut := flag.Bool("l", false, "Dont print in standard output")

	flag.Parse()

	var flags = map[string]bool{
		"e": *qEx,
		"l": *lOut,
	}
	return flags

}

func loadConfigFiles(filename string) map[string]string {
	config := make(map[string]string)

	file, err := os.Open(filename)
	Check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		//ignore empty lines or comments
		if strings.TrimSpace(line) == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			config[key] = value
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading config file:", err)
	}

	return config

}
