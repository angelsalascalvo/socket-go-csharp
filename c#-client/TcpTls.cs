using System;
using System.Net.Security;
using System.Net.Sockets;
using System.Security.Authentication;
using System.Security.Cryptography.X509Certificates;
using System.Text;
using NetworkConn;

/*
    COMUNICACION TCP SEGURA A TRAVES DEL PROTOCOLO TLS 1.2
    Envio de informacion en texto plano UTF8 cifrada usando el protocolo de seguridad TLS haciendo uso de 
    un certificado autofirmado.
*/
class TcpTls{
    private static readonly string serverDirection = "127.0.0.1";
    private static readonly int serverPort = 2021;
    
    //------------------------------------------------------------>

    /**
    * Establecer la comunicacion con el servidor
    */
    public void startClient(){
       
        //Establecer la conexion TCP con el servidor con el certificado digital
        TcpClient client = new TcpClient(serverDirection, serverPort);
        SslStream sslStream = new SslStream(client.GetStream(), false, ValidateServerCertificate); //Crear flujo comunicacion segura
        sslStream.AuthenticateAsClient(serverDirection); // Autenticarse con el servidor bajo el protocolo TLS1.2
        
        if(sslStream.IsAuthenticated){
            Console.WriteLine("✅ Conectado al servidor");
            
            //Enviar datos al servidor
            byte[] outputBuffer = new byte[512];
            outputBuffer = Encoding.UTF8.GetBytes("Hola caracola 🐌");
            sslStream.Write(outputBuffer);
            sslStream.Flush(); // Vaciar el buffer para hacer que los datos de el se envíen inmediatamente

            //Recibir datos del servidor
            byte[] inputBuffer = new byte[512];
            sslStream.Read(inputBuffer, 0, inputBuffer.Length);
            Console.WriteLine("\n🟦 Recibido mensaje del servidor:");        
            Console.WriteLine(Encoding.UTF8.GetString(inputBuffer));
            NetworkConnMain.PrintPackageBytes(inputBuffer);
        
            //Cerrar la comunicacion con el servidor
            client.Close();
        }
    }
    
    //------------------------------------------------------------>

    /*
    * Validacion de un certificado (Creado por una Autoridad de certificación - AC)
    */
     public static bool ValidateServerCertificate( object sender, X509Certificate certificate, X509Chain chain, SslPolicyErrors sslPolicyErrors) {
        
        return true; //Temp. Validacion fija al usar un certificado autogenerado (Para entorno de depuracion/pruebas)
        
        if (sslPolicyErrors == SslPolicyErrors.None)
            return true;

        Console.WriteLine("Certificate error: {0}"+ sslPolicyErrors);
        return false;
    }
}