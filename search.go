package main

import (
	"fmt"
	"strings"
)

func isValidCommand(key string) bool {
	_, exists := commandDescriptions[key]
	return exists
}

var commandDescriptions = map[string][]string{
	"cd":       {"Cambiar directorios de trabajo", "Changes the working directory"},
	"cp":       {"Copiar archivos", "Copy files "},
	"find":     {"Buscar archivos en carpetas, directorios", "Look for files in folders"},
	"ls":       {"Listar archivos y carpetas o directorios", "list folders and files "},
	"mkdir":    {"Crear carpetas o directorios", "Create folders"},
	"mv":       {"Mover o renombrar archivos", "Move or rename files"},
	"rm":       {"Eliminar, borrar archivos o carpetas, directorios", "Delete, remove files or folders"},
	"tree":     {"Mostrar la estructura de carpetas directorios", "Shows folder structure"},
	"grep":     {"Buscar, filtrar texto en archivos", "Look for text in files"},
	"touch":    {"Actualizar la fecha de modificación de un archivo, crear un archivo", "Update the modification date of a file, create a file"},
	"chmod":    {"Cambiar los permisos de archivos y carpetas o directorios", "Changes the permissions of files and folders"},
	"wc":       {"Contar lineas, palabras y bytes en archivos", "Counts lines, words and bytes in files"},
	"ssh":      {"Conectar a un servidor remoto", "Connect to a remote server"},
	"scp":      {"Copiar archivos de forma segura entre sistemas a través de SSH", "Copy files securely between systems through SSH"},
	"cat":      {"Mostrar, concatenar, unir el contenido de un archivo", "Show, join the content of a file"},
	"echo":     {"Imprimir, texto en la terminal", "Print text on the terminal"},
	"head":     {"Muestra, mostrar, las primeras lineas de un archivo", "Show the first lines of a file"},
	"tail":     {"Muestra, mostrar, las ultimas lineas de un archivo", "Show the last lines of a file"},
	"cmp":      {"Comparar archivos", "Compare files"},
	"sed":      {"Editor de texto en linea", "Inline editor for text"},
	"split":    {"Divide, dividir, separar un archivo en partes", "Split a file into parts"},
	"du":       {"Muestra, mostrar, el tamaño de archivos y carpetas o directorios en disco", "Show the size of files and folders"},
	"df":       {"Muestra, mostrar, el uso de disco", "Show the disk usage"},
	"diff":     {"Comparar dos archivos", "Compare two files"},
	"wget":     {"Descargar archivos", "Download files"},
	"curl":     {"Transfiere datos desde o hacia un servidor utilizando múltiples protocolos como HTTP, HTTPS, FTP", "Transfer data from or to a server using multiple protocols such as HTTP, HTTPS, FTP"},
	"top":      {"Monitoriza, monitorear procesos", "Monitor processes"},
	"service":  {"Inicia, detiene o consulta el estado de servicios en sistemas basados en SysVinit o systemd", "Start, stop or query the status of services on systems based on SysVinit or systemd"},
	"ps":       {"Muestra, mostrar los procesos activos", "Show active processes"},
	"kill":     {"Envía, enviar una señal a un proceso para finalizarlo, matarlo o interrumpirlo", "Send a signal to a process to kill, terminate it or interrupt it"},
	"killall":  {"Matar todos los procesos que coincidan con el nombre de un comando", "Kill all processes that match the name of a command"},
	"sort":     {"Ordenar, clasificar un archivo", "Sort a file"},
	"uniq":     {"unico Eliminar, duplicados de un archivo", "unique Remove duplicates from a file "},
	"nl":       {"Enumerar lineas de un archivo", "Number lines of a file"},
	"clear":    {"Limpiar la pantalla del terminal", "Clear the terminal screen"},
	"sudo":     {"Permite ejecutar un comando con privilegios de superusuario o como otro usuario", "Allows running a command with elevated privileges or as another user"},
	"whoami":   {"Muestra, mostrar el nombre de usuario actual que está ejecutando el terminal", "Displays, shows the current username that is executing the terminal"},
	"ifconfig": {"Muestra, mostrar y configura interfaces de red en sistemas Unix y Linux", "Displays and configures network interfaces on Unix and Linux systems"},
	"netstat":  {"Muestra, mostrar estadísticas de red, conexiones de red y puertos abiertos", "Displays network statistics, network connections and open ports"},
	"pwd":      {"Muestra, mostrar el directorio de trabajo actual", "Displays the current working directory"},
	"chown":    {"Cambiar el propietario y el grupo de uno o más archivos o directorios", "Changes the owner and group of one or more files or directories"},
	"said":     {"ayuda para comandos de linux", "help for linux commands"},
}

func smartSearch(query string) []string {
	var results []string
	query = strings.ToLower(query)
	fmt.Println("Searching: " + query + "...")

	// search descriptions
	for cmd, descs := range commandDescriptions {
		for _, desc := range descs {
			if strings.Contains(strings.ToLower(desc), query) {
				results = append(results, fmt.Sprintf("- %s: %s", cmd, desc))
				break
			}
		}
	}

	return results
}
