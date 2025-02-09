# SAID - Guía de Comandos de Linux 🐧  

SAID es una herramienta que proporciona información rápida y útil sobre comandos de Linux. Con SAID, puedes obtener descripciones, ejemplos de uso y opciones avanzadas para cada comando.  

## 🚀 Funcionalidades  

### 📌 Consultar un comando  
Muestra información detallada sobre un comando específico.  


```bash
said grep
```
---

```
Comando: grep  

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
Para más detalles: `man grep`
```

### 🔍 Buscar comandos 
Encuentra comandos relacionados con una palabra clave en su descripción.


```bash
said search archivos
```
--- 

```bash
Coincidencias encontradas:
- usermod: Modificar un usuario existente
- useradd: Crear un nuevo usuario
- passwd: Cambiar la contraseña de un usuario
- whoami: Muestra, mostrar el nombre de usuario actual que está ejecutando el terminal
- sudo: Permite ejecutar un comando con privilegios de superusuario o como otro usuario
```
### 🌐 Más información

🐙 GitHub: https://github.com/urzuaf/said

---

📢 SAID: Aprende y domina los comandos de Linux de manera rápida y sencilla! 🚀