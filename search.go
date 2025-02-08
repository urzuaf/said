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
	"cd":         {"Cambiar directorios de trabajo", "Changes the working directory"},
	"cp":         {"Copiar archivos", "Copy files "},
	"find":       {"Buscar archivos en carpetas, directorios", "Look for files in folders"},
	"ls":         {"Listar archivos y carpetas o directorios", "list folders and files "},
	"mkdir":      {"Crear carpetas o directorios", "Create folders"},
	"mv":         {"Mover o renombrar archivos", "Move or rename files"},
	"rm":         {"Eliminar, borrar archivos o carpetas, directorios", "Delete, remove files or folders"},
	"tree":       {"Mostrar la estructura de carpetas directorios", "Shows folder structure"},
	"grep":       {"Buscar, filtrar texto en archivos", "Look for text in files"},
	"touch":      {"Actualizar la fecha de modificación de un archivo, crear un archivo", "Update the modification date of a file, create a file"},
	"chmod":      {"Cambiar los permisos de archivos y carpetas o directorios", "Changes the permissions of files and folders"},
	"wc":         {"Contar lineas, palabras y bytes en archivos", "Counts lines, words and bytes in files"},
	"ssh":        {"Conectar a un servidor remoto", "Connect to a remote server"},
	"scp":        {"Copiar archivos de forma segura entre sistemas a través de SSH", "Copy files securely between systems through SSH"},
	"cat":        {"Mostrar, concatenar, unir el contenido de un archivo", "Show, join the content of a file"},
	"echo":       {"Imprimir, texto en la terminal", "Print text on the terminal"},
	"head":       {"Muestra, mostrar, las primeras lineas de un archivo", "Show the first lines of a file"},
	"tail":       {"Muestra, mostrar, las ultimas lineas de un archivo", "Show the last lines of a file"},
	"cmp":        {"Comparar archivos", "Compare files"},
	"sed":        {"Editor de texto en linea", "Inline editor for text"},
	"split":      {"Divide, dividir, separar un archivo en partes", "Split a file into parts"},
	"du":         {"Muestra, mostrar, el tamaño de archivos y carpetas o directorios en disco", "Show the size of files and folders"},
	"df":         {"Muestra, mostrar, el uso de disco", "Show the disk usage"},
	"diff":       {"Comparar dos archivos", "Compare two files"},
	"wget":       {"Descargar archivos", "Download files"},
	"curl":       {"Transfiere datos desde o hacia un servidor utilizando múltiples protocolos como HTTP, HTTPS, FTP", "Transfer data from or to a server using multiple protocols such as HTTP, HTTPS, FTP"},
	"top":        {"Monitoriza, monitorear procesos", "Monitor processes"},
	"service":    {"Inicia, detiene o consulta el estado de servicios en sistemas basados en SysVinit o systemd", "Start, stop or query the status of services on systems based on SysVinit or systemd"},
	"ps":         {"Muestra, mostrar los procesos activos", "Show active processes"},
	"kill":       {"Envía, enviar una señal a un proceso para finalizarlo, matarlo o interrumpirlo", "Send a signal to a process to kill, terminate it or interrupt it"},
	"killall":    {"Matar todos los procesos que coincidan con el nombre de un comando", "Kill all processes that match the name of a command"},
	"sort":       {"Ordenar, clasificar un archivo", "Sort a file"},
	"uniq":       {"unico Eliminar, duplicados de un archivo", "unique Remove duplicates from a file "},
	"nl":         {"Enumerar lineas de un archivo", "Number lines of a file"},
	"clear":      {"Limpiar la pantalla del terminal", "Clear the terminal screen"},
	"sudo":       {"Permite ejecutar un comando con privilegios de superusuario o como otro usuario", "Allows running a command with elevated privileges or as another user"},
	"whoami":     {"Muestra, mostrar el nombre de usuario actual que está ejecutando el terminal", "Displays, shows the current username that is executing the terminal"},
	"ifconfig":   {"Muestra, mostrar y configura interfaces de red en sistemas Unix y Linux", "Displays and configures network interfaces on Unix and Linux systems"},
	"netstat":    {"Muestra, mostrar estadísticas de red, conexiones de red y puertos abiertos", "Displays network statistics, network connections and open ports"},
	"pwd":        {"Muestra, mostrar el directorio de trabajo actual", "Displays the current working directory"},
	"chown":      {"Cambiar el propietario y el grupo de uno o más archivos o directorios", "Changes the owner and group of one or more files or directories"},
	"zip":        {"Comprimir archivos en un solo archivo ZIP", "Compress files into a ZIP archive"},
	"unzip":      {"Extraer archivos de un archivo ZIP", "Extract files from a ZIP archive"},
	"ln":         {"Crear enlaces simbólicos o duros", "Create symbolic or hard links"},
	"less":       {"Ver el contenido de un archivo página por página", "View file content page by page"},
	"uname":      {"Muestra información del sistema operativo", "Display system information"},
	"tar":        {"Crear o extraer archivos comprimidos en formato tar", "Create or extract tar archives"},
	"comm":       {"Comparar líneas comunes y únicas entre archivos", "Compare common and unique lines between files"},
	"export":     {"Definir o modificar variables de entorno", "Define or modify environment variables"},
	"mount":      {"Montar un sistema de archivos", "Mount a filesystem"},
	"traceroute": {"Mostrar la ruta que siguen los paquetes en la red hasta un destino", "Display the route packets take to a destination"},
	"ufw":        {"Configurar el firewall en sistemas basados en Ubuntu", "Configure the firewall on Ubuntu-based systems"},
	"iptables":   {"Administrar reglas de filtrado de paquetes en Linux", "Manage packet filtering rules in Linux"},
	"cal":        {"Muestra un calendario en la terminal", "Display a calendar in the terminal"},
	"alias":      {"Definir alias para comandos", "Define aliases for commands"},
	"dd":         {"Copiar y convertir datos a bajo nivel", "Copy and convert data at a low level"},
	"whereis":    {"Ubicar el binario, código fuente y manual de un comando", "Locate the binary, source, and manual of a command"},
	"whatis":     {"Muestra una breve descripción de un comando", "Display a brief description of a command"},
	"useradd":    {"Crear un nuevo usuario", "Create a new user"},
	"usermod":    {"Modificar un usuario existente", "Modify an existing user"},
	"passwd":     {"Cambiar la contraseña de un usuario", "Change a user's password"},
	"history":    {"Muestra el historial de comandos ejecutados en la terminal", "Show the history of commands executed in the terminal"},
	"sleep":      {"Pausa la ejecución de un programa o script por un período de tiempo especificado", "Pause the execution of a program or script for a specified period"},
	"basename":   {"Extrae el nombre de archivo de una ruta de archivo", "Extract the filename from a file path"},
	"dirname":    {"Extrae el nombre del directorio de una ruta de archivo", "Extract the directory name from a file path"},
	"stat":       {"Muestra información detallada sobre un archivo o directorio", "Show detailed information about a file or directory"},
	"hostname":   {"Muestra o configura el nombre del host del sistema", "Show or set the system's hostname"},
	"groups":     {"Muestra los grupos a los que pertenece un usuario", "Show the groups a user belongs to"},
	"id":         {"Muestra la información del usuario y grupo, como el ID de usuario (UID) y el ID de grupo (GID)", "Show user and group information, such as user ID (UID) and group ID (GID)"},
	"xargs":      {"Toma entrada estándar y la convierte en argumentos para ejecutar un comando", "Take standard input and convert it into arguments for a command"},
	"bzip2":      {"Herramienta de compresión para archivos, similar a gzip", "File compression tool, similar to gzip"},
	"gunzip":     {"Descomprime archivos .gz", "Decompress .gz files"},
	"gzip":       {"Comprime archivos", "Compress files"},
	"shred":      {"Sobrescribe un archivo para hacerlo irrecuperable", "Overwrite a file to make it unrecoverable"},
	"fg":         {"Trae un trabajo en segundo plano a primer plano", "Bring a background job to the foreground"},
	"bg":         {"Envía un trabajo a segundo plano", "Send a job to the background"},
	"jobs":       {"Muestra los trabajos actuales en segundo plano", "Show the current background jobs"},
	"trap":       {"Captura señales y las maneja dentro de un script", "Capture signals and handle them within a script"},
	"said":       {"ayuda para comandos de linux", "help for linux commands"},
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
