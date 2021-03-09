using System;

namespace NetworkConn{
    class NetworkConnMain{

        static void Main(string[] args){
            Console.OutputEncoding = System.Text.Encoding.UTF8;

            //Ejemplo 1
            Tcp tcp = new Tcp();
            //tcp.startClient();

            //Ejemplo 2
            TcpTls tcpTls = new TcpTls();
            //tcpTls.startClient();
            
            //Ejemplo 3
            TcpTlsJson tcpTlsJson = new TcpTlsJson();
            tcpTlsJson.startClient();
        }

        //------------------------------------------------------------>

        /*
        *Imprimir contenido de un paquete de datos (array de bytes)
        */
        public static void PrintPackageBytes(byte[] data) {
            string result = "{ ";
            for (int i = 0; i < data.Length; i++){
                result += data[i].ToString() + (i==data.Length-1? " " : ", ");
            }
            result += "}";
            Console.WriteLine(result +" size: ("+data.Length+")");
        }
    }
}
//dotnet run