/*
 * SERVIDOR TCP CON COMUNICACION SEGURA A TRAVES DEL PROTOCOLO TLS Y ENVIO JSON
 * Envio de informacion en formato json de forma segura (uso de tls con certificado autofirmado)
 */
package tcptlsjson

import (
	"crypto/rand"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

//StartServer Inicia el servidor TCP - TLS
func StartServer() {
	//Cargar el certificado y la clave privada de la comunicacion
	cert, err := tls.LoadX509KeyPair("../certificates/cert.pem", "../certificates/private-key.pem")
	checkError(err) //Comprobar errores

	//Crear configuracion del socket seguro
	config := tls.Config{Certificates: []tls.Certificate{cert}}
	now := time.Now()
	config.Time = func() time.Time { return now } //Establecer hora actual
	config.Rand = rand.Reader                     //Numero aleatorio utilizado por TLS para el Blinding de RSA

	//Iniciar el socket de servidor
	service := "127.0.0.1:2022"
	socket, err := tls.Listen("tcp", service, &config)
	checkError(err) // comprobar errores

	//Defer hace que el codigo se ejecute al final de la ejecucion del bloque de codigo
	defer socket.Close()

	//Escuchar nuevos clientes
	fmt.Println("Servidor TCP-TLS (json) esperando clientes...")
	for {
		conn, err := socket.Accept()
		//Si ocurre un error continuar con el bucle de escucha
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		fmt.Println("✅ Nuevo cliente conectado (" + conn.RemoteAddr().String() + ")")
		//Iniciar gorutina de comunicacion
		go handleClientCommunication(conn)
	}
}

//------------------------------------------------------------>

// Manejar la comunicacion con un cliente
func handleClientCommunication(conn net.Conn) {
	//Defer hace que el codigo se ejecute al final de la ejecucion del bloque de codigo
	defer conn.Close()

	inputBuffer := make([]byte, 512)
	for {
		// Leer mensajes
		n, err := conn.Read(inputBuffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
		}
		var user User
		json.Unmarshal(inputBuffer[:n], &user) //Convertir el slice de bytes (contenido json) en un objeto

		fmt.Print("\n---------------------------------------\n")
		fmt.Println("🟦 Recibido paquete del cliente " + conn.RemoteAddr().String() + ":")
		fmt.Println("Se ha recibido el objeto json de tipo usuario: ", user)
		fmt.Println(inputBuffer)
		fmt.Print("---------------------------------------\n")

		if user.Id == 0 {
			user.Admin = true //Como ejemplo si el id del usuario recibido es administrador lo marcamos
		}

		//Enviar los datos al cliente con el campo admin del usuario establecido
		outputBuffer, _ := json.Marshal(user)
		_, err2 := conn.Write(outputBuffer)
		if err2 != nil {
			fmt.Println(err2)
		}
	}

	println("🔌 Fin de la conexion del cliente " + conn.RemoteAddr().String())
}

//------------------------------------------------------------>

// User es la estructura empleada para parsear los datos json de la comunicacion por sockets
type User struct {
	Id    int
	Name  string
	Admin bool
}

//------------------------------------------------------------>

//Imprimir un error
func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
