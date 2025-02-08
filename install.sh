#!/bin/bash

ZIP_URL="https://github.com/urzuaf/said/releases/download/v0.3.0/said.zip"

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
