const copySVG = `<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="size-6">
  <path stroke-linecap="round" stroke-linejoin="round" d="M8.25 7.5V6.108c0-1.135.845-2.098 1.976-2.192.373-.03.748-.057 1.123-.08M15.75 18H18a2.25 2.25 0 0 0 2.25-2.25V6.108c0-1.135-.845-2.098-1.976-2.192a48.424 48.424 0 0 0-1.123-.08M15.75 18.75v-1.875a3.375 3.375 0 0 0-3.375-3.375h-1.5a1.125 1.125 0 0 1-1.125-1.125v-1.5A3.375 3.375 0 0 0 6.375 7.5H5.25m11.9-3.664A2.251 2.251 0 0 0 15 2.25h-1.5a2.251 2.251 0 0 0-2.15 1.586m5.8 0c.065.21.1.433.1.664v.75h-6V4.5c0-.231.035-.454.1-.664M6.75 7.5H4.875c-.621 0-1.125.504-1.125 1.125v12c0 .621.504 1.125 1.125 1.125h9.75c.621 0 1.125-.504 1.125-1.125V16.5a9 9 0 0 0-9-9Z" />
</svg>`

const doneVG = `<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="size-6">
  <path stroke-linecap="round" stroke-linejoin="round" d="M11.35 3.836c-.065.21-.1.433-.1.664 0 .414.336.75.75.75h4.5a.75.75 0 0 0 .75-.75 2.25 2.25 0 0 0-.1-.664m-5.8 0A2.251 2.251 0 0 1 13.5 2.25H15c1.012 0 1.867.668 2.15 1.586m-5.8 0c-.376.023-.75.05-1.124.08C9.095 4.01 8.25 4.973 8.25 6.108V8.25m8.9-4.414c.376.023.75.05 1.124.08 1.131.094 1.976 1.057 1.976 2.192V16.5A2.25 2.25 0 0 1 18 18.75h-2.25m-7.5-10.5H4.875c-.621 0-1.125.504-1.125 1.125v11.25c0 .621.504 1.125 1.125 1.125h9.75c.621 0 1.125-.504 1.125-1.125V18.75m-7.5-10.5h6.375c.621 0 1.125.504 1.125 1.125v9.375m-8.25-3 1.5 1.5 3-3.75" />
</svg>
`

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

let uninstall = `#Delete the said binary
sudo rm /usr/bin/said

#Delete the said data
sudo rm -r ~/.said/*`

document.getElementById("inst").innerHTML = instalation;
document.getElementById("ge").innerHTML = grepExample;
document.getElementById("se").innerHTML = searchExample;
document.getElementById("un").innerHTML = uninstall;

function addCopyButton(codeElement) {
    const button = document.createElement("button");
    button.innerHTML = copySVG
    button.style.position = "absolute";
    button.style.top = "8px";
    button.style.right = "8px";
    button.style.cursor = "pointer";
    button.style.fontSize = "16px";
    button.style.borderRadius = "8px";
    button.style.padding = "4px"
    button.style.color = "#969696"
    button.style.transition = "color 0.1s ease";
    button.onmouseenter = () => {
        button.style.color = "#dddddd"
    }
    button.onmouseleave = () => {
        button.style.color = "#969696"
    }

    const pre = codeElement.parentElement;
    pre.style.position = "relative";
    
    button.addEventListener("click", () => {
        navigator.clipboard.writeText(codeElement.innerText).then(() => {
            button.innerHTML = doneVG;
            setTimeout(() => button.innerHTML=copySVG, 1500);
        });
    });

    pre.appendChild(button);
}


document.querySelectorAll("code").forEach(addCopyButton);
