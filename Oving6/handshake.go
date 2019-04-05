/* DaytimeServer
 */
package main

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	service := ":1300"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Print("hva faen?")
			continue
		}
		handleClient(conn)
		sendMessage(conn)

		conn.Close()// we're finished with this client
	}
}

func handleClient(conn net.Conn){
	//var bodyStart  = "<!DOCTYPE html><HTML><body><H1> Hilsen. Du har koblet deg opp til min enkle web-tjener </h1>Header fra klient er: <ul>"
	//var bodyEnd  = "</ul></body></HTML>\r\n\r\n"


	var buf [512]byte
	fmt.Println( "buf: ", buf)
	conn.Read(buf[0:])
	fmt.Println(buf)
	var bufRead = string(buf[0:])
	fmt.Println("bufferRead",bufRead)
	var secWebSocketKey string
	temp := strings.Split(bufRead,"\n")


	for _, element := range temp {
		var el = strings.TrimSpace(element)

		if strings.Contains(el, "Sec-WebSocket-Key:"){
			fmt.Print("\nforhåpentligvis Sec-WebSocket-Key:" +el+"\n\n")
			secWebSocketKey = strings.Trim(el, "Sec-WebSocket-Key: ")
		}
	}


	var secWebSocketAcceptString = secWebSocketKey +"258EAFA5-E914-47DA-95CA-C5AB0DC85B11" //concater sec-WebSocket-Key med den nøkkel-greia fra RFc-6455
	var h = sha1.New()  //sha1 krypterer den slik den forlanger at det skal gjøres
	h.Write([]byte(secWebSocketAcceptString)) //skriver den ut som en streng

	var secWebSocketAccept = base64.StdEncoding.EncodeToString(h.Sum(nil)) //enkoder den til base64 som RFC forlanger
	var header  = "HTTP/1.1 101 Switching Protocols\r\nUpgrade: websocket\r\nConnection: Upgrade\r\nSec-WebSocket-Accept: " + secWebSocketAccept +"\r\n\r\n"


	conn.Write([]byte(header)) // don't care about return value




}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
func sendMessage(conn net.Conn){
	startBeskjed := []byte{0x81} //bitmønseter som MÅ sendes over i starten er 1000 0001 -> tilsvarer hex 81.
	//bitmønsteret er valgt: første bit: 1 betyr FIN, dette er siste del av melding.
	//etterfulgt av 3 0'er --> obligatoriske nuller
	//4 neste bits: op-kode til klient: velger op-kode 0001 = tekst

	beskjed := []byte("hei") //beskjeden vi ønsker å sende over på byteformat
	beskjedLengde := len(beskjed) //vi skal ikke maskere beskjeden, og vi er rimelig sikker på at beskjeden er under 127 bytes. dermed blir de syv bitene bare lengden på beskjeden

	beskjedIByte := byte(beskjedLengde)


	beskjedSendes := append(append(startBeskjed, beskjedIByte), beskjed... ) //setter sammen framen
	fmt.Print("utskrift av beskjed som sendes over: start av beskjed: ", startBeskjed[0:], "lengde på beskjed: ", beskjedLengde, ", selve beskjeden: ", beskjed[0:], "\n")
	fmt.Print("totalbeskjed: ", beskjedSendes[0:])


	conn.Write(beskjedSendes)
}