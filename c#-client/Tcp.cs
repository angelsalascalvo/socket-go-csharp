using System;
using System.Net;
using System.Net.Sockets;
using System.Text;
using NetworkConn;

/*
    COMUNICACION TCP SIN USO DE MECANISMO DE SEGURIDAD
    Envio de informacion en texto plano UTF8
*/
class Tcp{
    private static readonly string serverDirection = "127.0.0.1";
    private static readonly int serverPort = 2020;
    
    //------------------------------------------------------------>

    /**
    * Establecer la comunicacion con el servidor
    */
    public void startClient(){
        //Establecer la conexion TCP con el servidor
        IPEndPoint server = new IPEndPoint(IPAddress.Parse(serverDirection), serverPort);
        Socket socket = new Socket(AddressFamily.InterNetwork, SocketType.Stream, ProtocolType.Tcp);
        socket.Connect(server);
        Console.WriteLine("✅ Conectado al servidor");

        if(socket.Connected){
            //Enviar datos al servidor
            byte[] outputBuffer = new byte[512];
            outputBuffer = Encoding.UTF8.GetBytes("Digamelón 🍈");
            socket.Send(outputBuffer);

            //Recibir datos del servidor
            byte[] inputBuffer = new byte[512];
            socket.Receive(inputBuffer);
            Console.WriteLine("\n🟦 Recibido mensaje del servidor:");        
            Console.WriteLine(Encoding.UTF8.GetString(inputBuffer));
            NetworkConnMain.PrintPackageBytes(inputBuffer);
        
            //Cerrar la comunicacion con el servidor
            socket.Close();
        }
    }
    
}