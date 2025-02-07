# Said (Shell-Aid) 🖥️

Said es una herramienta de línea de comandos diseñada para proporcionar ayuda rápida sobre comandos de Linux, con explicaciones claras, opciones comunes y ejemplos prácticos.

## 🚀 Instalación

*Pendiente*

## 🛠 Uso

Para obtener información sobre un comando, simplemente ejecuta:
```
said <comando>
```
Ejemplo:
```
said ls
```
Salida esperada:
```
📖 Descripción:

  Lista los archivos y directorios en una carpeta.

🛠️ Uso :

  ls [opciones] [ruta]

🔹 Ejemplos :

    ls -l   # Lista en formato detallado
    ls -a   # Muestra archivos ocultos
    ls -lh  # Tamaño en formato legible

📝 Opciones más usadas:

    -l → Formato largo con permisos, usuario, etc.
    -a → Muestra archivos ocultos (.archivos)
    -h → Tamaños en KB/MB en vez de bytes

ℹ️ Información adicional:
  Para más detalles: man ls
```

### Buscar
Si no conoces el nombre del comando puedes buscarlo a partir de palabras claves

```
said search usuario
```
salida:
```
Searching: usuario...
Coincidencias encontradas:
- usermod: Modificar un usuario existente
- useradd: Crear un nuevo usuario
- passwd: Cambiar la contraseña de un usuario
- whoami: Muestra, mostrar el nombre de usuario actual que está ejecutando el terminal
- sudo: Permite ejecutar un comando con privilegios de superusuario o como otro usuario
```
### Listar
Lista todos los comandos disponibles
```
said list
```
salida:
```
chown         cd          cp       
cmp           head        df       
traceroute    iptables    dd       
mkdir         ifconfig    comm     
whoami        said        wget     
top           ps          less     
useradd       scp         du       
diff          sed         sudo     
unzip         zip         tar      
touch         tail        split    
nl            whereis     echo     
whatis        passwd      rm       
chmod         pwd         kill     
mv            wc          ssh      
ln            ufw         find     
service       netstat     sort     
mount         alias       grep     
curl          killall     clear    
usermod       ls          cat      
uniq          cal         tree     
uname         export    
```

¡Esperamos que Said sea una herramienta útil en tu terminal! 🚀