
let instalation = `ZIP_URL="https://github.com/urzuaf/said/releases/download/v0.3.0/said.zip"

INSTALL_DIR="$HOME/.said"

# Directorio de binarios 
BIN_DIR="/usr/bin"

# Verificar si el directorio de instalación existe, si no, crear
if [ ! -d "$INSTALL_DIR" ]; then
    echo "Creando directorio de instalación en $INSTALL_DIR..."
    mkdir -p "$INSTALL_DIR"
fi

echo "Descargando el archivo ..."
curl -L "$ZIP_URL" -o "$INSTALL_DIR/said.zip"

echo "Extrayendo el archivo ..."
unzip -q "$INSTALL_DIR/said.zip" -d "$INSTALL_DIR"

chmod +x "$INSTALL_DIR/said" 

echo "moviendo el binario a /usr/bin"
sudo mv $INSTALL_DIR/said $BIN_DIR/said

rm "$INSTALL_DIR/said.zip"

echo "¡Instalación completada!" 
`

let grepExample = `Comando: grep
        
📖 Descripción: 
    Buscar un patrón en archivos de texto. 
        
🛠️ Uso:  
    grep [opciones] "patrón" archivo  

🔹 Ejemplos:  
    grep "error" log.txt    # Busca "error" en el archivo log.txt  
    grep -i "usuario" *     # Búsqueda sin distinguir mayúsculas/minúsculas   
            
📝 Opciones más usadas:  
    -i → Ignora mayúsculas y minúsculas  
    -r → Busca recursivamente en carpetas  
    -n → Muestra el número de línea  
            
ℹ️ Información adicional:  
    Para más detalles: man grep 
`

let searchExample = `Coincidencias encontradas:
- usermod: Modificar un usuario existente
- useradd: Crear un nuevo usuario
- passwd:  Cambiar la contraseña de un usuario
- whoami: Muestra, mostrar el nombre de usuario actual que está ejecutando el terminal
- sudo:   Permite ejecutar un comando con privilegios de superusuario o como otro usuario
`

document.getElementById("inst").innerHTML = instalation;
document.getElementById("ge").innerHTML = grepExample;
document.getElementById("se").innerHTML = searchExample;