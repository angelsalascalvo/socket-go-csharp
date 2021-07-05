using System;
using System.Net.Security;
using System.Net.Sockets;
using System.Security.Authentication;
using System.Security.Cryptography.X509Certificates;
using System.Text;
using Newtonsoft.Json;
using NetworkConn;

/*
    COMUNICACION TCP SEGURA A TRAVES DEL PROTOCOLO TLS 1.2 + ENVIO JSON
    Envio de informacion en formato JSON cifrada usando el protocolo de seguridad TLS haciendo uso de 
    un certificado autofirmado.
*/
class TcpTlsJson{
    private static readonly string serverDirection = "127.0.0.1";
    private static readonly int serverPort = 2022;
    
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
            User user = new User();
            user.id = 4;
            user.name = "Tobias";
            byte[] outputBuffer = Encoding.UTF8.GetBytes(JsonConvert.SerializeObject(user)); //Convertir el objeto en formato json y a su vez en array de bytes para el envio
            sslStream.Write(outputBuffer);
            sslStream.Flush(); // Vaciar el buffer para hacer que los datos de el se envíen inmediatamente

            //Recibir datos del servidor
            byte[] inputBuffer = new byte[512];
            sslStream.Read(inputBuffer, 0, inputBuffer.Length);
            user = JsonConvert.DeserializeObject<User>(Encoding.UTF8.GetString(inputBuffer));

            Console.WriteLine("\n🟦 Recibida respuesta del servidor:");        
            Console.WriteLine("Segun el servidor "+user.name+ " "+ (user.admin?"es administrador.":"no es administrador"));
            NetworkConnMain.PrintPackageBytes(inputBuffer);
        
            //Cerrar la comunicacion con el servidor
            client.Close();
        }
    }
    
    //------------------------------------------------------------>

    public class User {
        public int id;
        public string name;
        public bool admin;
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