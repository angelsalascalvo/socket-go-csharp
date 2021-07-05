# Socket Go - C#

En este repositorio encontrarás ejemplos básicos de una comunicación a través de sockets TCP entre un cliente creado en C# y un servidor en Go Lang.
<br>En el proyecto se incluyen 3 ejemplos diferentes:

- **Comunicación TCP**: Transferencia básica de información en texto plano entre cliente (C#) y servidor (Go) a través de un socket bajo el protocolo TCP.
- **Comunicación TCP Segura**: Transferencia básica de información en texto plano entre cliente (C#) y servidor (Go) a través de un socket TCP cifrado mediante del protocolo de seguridad TLS 1.2. Para el correcto funcionamiento de este, se hace uso del certificado autofirmado y la clave privada incluidas en el directorio "certificates" con fines de prueba/testeo.
- **Comunicación TCP Segura + JSON**: Transferencia básica de información en formato JSON entre cliente (C#) y servidor (Go) a través de un socket TCP + TLS 1.2. <br>La transferencia en formato JSON permite la serialización / deserialización de la información transmitida en objetos/estructuras del lenguaje, con las que poder operar fácilmente.
<br>

## Estructura de directorios

En el proyecto encontramos 3 directorios diferentes:

- **C#-Client**: Proyecto C Sharp con los diferentes ejemplos sobre la funcionalidad de cliente.
- **Go-Server**: Proyecto Go Lang con los diferentes ejemplos sobre la funcionalidad de servidor.
- **Certificates**: Contiene la clave privada, certificado (autofirmado) utilizados a modo de prueba para la implementación de TLS en los ejemplos.

>  *(Para el uso del protocolo de seguridad TLS dentro de un entorno real es necesario emplear un certificado firmado por una Autoridad de Certificación AC)*
<br>

## Ejecución
Para el correcto funcionamiento debemos asegurarnos de ejecutar tanto en el cliente como en el servidor el codigo correspondiente al mismo ejemplo. La ejecucion de un ejemplo u otro se establecerá desde la función principal u origen de estos (`main.go` `Program.cs`)
<br><br>
Para lanzar la propia ejecución se utilizarán los siguientes comandos de consola:
- Go `go run .\main.go`
- C# `dotnet run`
