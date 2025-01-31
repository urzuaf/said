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
	"cd":       {"Cambia, cambiar el directorio, directorios de trabajo", "Changes, change the working directory"},
	"cp":       {"Copia, copiar archivos, archivo", "Copy files, file"},
	"find":     {"Busca, buscar archivos,archivo en carpetas, carpeta", "Look for file, files in folders"},
	"ls":       {"Lista, listar archivos y carpetas", "list folders, folder and files, file"},
	"mkdir":    {"Crea, crear carpetas", "Create folders, folder"},
	"mv":       {"Mueve, mover o renombra,renombrar archivos", "Move or rename files"},
	"rm":       {"Elimina, borra, eliminar, borrar archivos o carpetas", "Delete, remove files or folders"},
	"tree":     {"Muestra, mostrar estructura de carpetas", "Shows folder structure"},
	"grep":     {"Busca, buscar, filtra, filtrar texto en archivos", "Look for text in files"},
	"touch":    {"Actualiza, actualizar la fecha de modificación de un archivo", "Update the modification date of a file", "crear, crea un archivo", "Create, make a file"},
	"chmod":    {"Cambia, cambiar los permisos de archivos y carpetas", "Changes, change the permissions of files and folders"},
	"wc":       {"Cuenta, contar, contar lineas, palabras y bytes en archivos", "Counts, count lines, words and bytes in files"},
	"ssh":      {"Conecta, conectar a un servidor remoto", "Connect to a remote server"},
	"scp":      {"Copia, copiar archivos de forma segura entre sistemas a través de SSH", "Copy files securely between systems through SSH"},
	"cat":      {"Muestra, mostrar, concatena, concatenar, unir, une el contenido de un archivo", "Show, join the content of a file"},
	"echo":     {"Imprime, imprimir, imprimir texto en la terminal", "Print, print text on the terminal"},
	"head":     {"Muestra, mostrar, las primeras lineas de un archivo", "Show, show the first lines of a file"},
	"tail":     {"Muestra, mostrar, las ultimas lineas de un archivo", "Show, show the last lines of a file"},
	"cmp":      {"Compara, comparar dos archivos", "Compare two files"},
	"sed":      {"Editor de texto enILINE, editor de texto enILINE", "Inline editor for text"},
	"split":    {"Divide, dividir, separa, separar un archivo en partes", "Split a file into parts"},
	"du":       {"Muestra, mostrar, el tamaño de archivos y carpetas", "Show, show the size of files and folders"},
	"df":       {"Muestra, mostrar, el uso de disco", "Show, show the disk usage"},
	"diff":     {"Compara, comparar dos archivos", "Compare two files"},
	"wget":     {"Descarga, descargar archivos", "Download files"},
	"curl":     {"Transfiere datos desde o hacia un servidor utilizando múltiples protocolos como HTTP, HTTPS, FTP.", "Transfer data from or to a server using multiple protocols such as HTTP, HTTPS, FTP."},
	"top":      {"Monitoriza, monitorear, monitorear procesos", "Monitor processes"},
	"service":  {"Inicia, detiene o consulta el estado de servicios en sistemas basados en SysVinit o systemd.", "Start, stop or query the status of services on systems based on SysVinit or systemd."},
	"ps":       {"Muestra, mostrar, los procesos activos", "Show, show active processes"},
	"kill":     {"Envía, enviar una señal a un proceso para finalizarlo o interrumpirlo", "Send, send a signal to a process to terminate it or interrupt it"},
	"killall":  {"Mata todos los procesos que coincidan con el nombre de un comando", "Kill all processes that match the name of a command"},
	"sort":     {"Ordena, ordenar, clasificar, clasificar un archivo", "Sort, sort a file"},
	"uniq":     {"unico Elimina, eliminar, duplicados de un archivo", "unique Remove duplicates from a file unique"},
	"nl":       {"Numerar, numerar, numerar lineas de un archivo", "Number lines of a file"},
	"clear":    {"Limpia, limpiar la pantalla del terminal", "Clear the terminal screen"},
	"sudo":     {"Permite ejecutar un comando con privilegios de superusuario o como otro usuario", "Allows running a command with elevated privileges or as another user"},
	"whoami":   {"Muestra el nombre de usuario actual que está ejecutando el terminal", "Displays the current username that is executing the terminal"},
	"ifconfig": {"Muestra y configura interfaces de red en sistemas Unix y Linux", "Displays and configures network interfaces on Unix and Linux systems"},
	"netstat":  {"Muestra estadísticas de red, conexiones de red y puertos abiertos", "Displays network statistics, network connections and open ports"},
	"pwd":      {"Muestra el directorio de trabajo actual (directorio en el que se encuentra el usuario)", "Displays the current working directory (directory where the user is located)"},
	"chown":    {"Cambia el propietario y el grupo de uno o más archivos o directorios", "Changes the owner and group of one or more files or directories"},
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
