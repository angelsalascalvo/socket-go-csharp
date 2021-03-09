/*
 * SERVIDOR TCP CON COMUNICACION SEGURA A TRAVES DEL PROTOCOLO TLS
 * Envio de informacion en texto plano UTF8 de forma segura (uso de tls con certificado autofirmado)
 */
package tcptls

import (
	"crypto/rand"
	"crypto/tls"
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
	service := "127.0.0.1:2021"
	socket, err := tls.Listen("tcp", service, &config)
	checkError(err) // comprobar errores

	//Defer hace que el codigo se ejecute al final de la ejecucion del bloque de codigo
	defer socket.Close()

	//Escuchar nuevos clientes
	fmt.Println("Servidor TCP-TLS esperando clientes...")
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
		txtInputBuffer := string(inputBuffer[:n])

		fmt.Print("\n---------------------------------------\n")
		fmt.Println("🟦 Recibido paquete del cliente " + conn.RemoteAddr().String() + ":")
		fmt.Println(txtInputBuffer)
		fmt.Println(inputBuffer)
		fmt.Print("---------------------------------------\n")

		// Enviar mensaje
		_, err2 := conn.Write([]byte("Mensaje ( " + txtInputBuffer + " ) recibido en servidor."))
		if err2 != nil {
			fmt.Println(err2)
		}
	}

	println("🔌 Fin de la conexion del cliente " + conn.RemoteAddr().String())
}

//------------------------------------------------------------>

//Imprimir un error
func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
