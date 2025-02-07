#!/bin/bash

ZIP_URL="https://github.com/urzuaf/said/releases/download/0.1.0/said.zip"

INSTALL_DIR="$HOME/.said"

# Directorio de binarios accesibles
BIN_DIR="/usr/bin"

# Verificar si el directorio de instalación existe, si no, crear
if [ ! -d "$INSTALL_DIR" ]; then
    echo "Creando directorio de instalación en $INSTALL_DIR..."
    mkdir -p "$INSTALL_DIR"
fi

# Descargar el archivo ZIP
echo "Descargando el archivo zip..."
curl -L "$ZIP_URL" -o "$INSTALL_DIR/said.zip"

# Extraer el archivo ZIP
echo "Extrayendo el archivo zip..."
unzip -q "$INSTALL_DIR/said.zip" -d "$INSTALL_DIR"


chmod +x "$INSTALL_DIR/said" 

# Verificar si hay archivos de configuración adicionales (por ejemplo said.conf)
if [ -f "$INSTALL_DIR/said.conf" ]; then
    echo "Configuración encontrada en $INSTALL_DIR/said.conf"
    echo "Por favor edita el archivo de configuración si es necesario."
fi

# Mover el ejecutable a la carpeta de binarios
mv $INSTALL_DIR/said $BIN_DIR/said

# Limpieza del archivo ZIP
echo "Eliminando archivo zip descargado..."
rm "$INSTALL_DIR/said.zip"

echo "¡Instalación completada!"
