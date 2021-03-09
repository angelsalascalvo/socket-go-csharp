/*
 * SERVIDOR TCP SIN PROTOCOLO DE SEGURIDAD IMPLEMENTADO
 * Envio de informacion en texto plano UTF8
 */
package tcp

import (
	"fmt"
	"io"
	"net"
	"os"
)

//StartServer Inicia el servidor TCP
func StartServer() {
	//Iniciar el socket de servidor
	service := "127.0.0.1:2020"
	socket, err := net.Listen("tcp", service)
	checkError(err) // comprobar errores

	//Defer hace que el codigo se ejecute al final de la ejecucion del bloque de codigo
	defer socket.Close()

	//Escuchar nuevos clientes
	fmt.Println("Servidor TCP esperando clientes...")
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
