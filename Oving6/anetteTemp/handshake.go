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

	service := ":1200"
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


		conn.Close()// we're finished with this client
	}
}

func handleClient(conn net.Conn){
	var bodyStart  = "<!DOCTYPE html><HTML><body><H1> Hilsen. Du har koblet deg opp til min enkle web-tjener </h1>Header fra klient er: <ul>"
	var bodyEnd  = "</ul></body></HTML>\r\n\r\n"


	var buf [512]byte
	fmt.Println( "buf: ", buf)
	conn.Read(buf[0:])
	fmt.Println(buf)
	var bufRead = string(buf[0:])
	fmt.Println("bufferRead",bufRead)
	var secWebSocketKey ="blahblah placeholder"
	temp := strings.Split(bufRead,"\n")

	var bodyMid= ""
	for _, element := range temp {
		var el = strings.TrimSpace(element)
		if len(el) !=0 && el != "  " {
			bodyMid += "<li>" + el + "</li>"
		}
		if strings.Contains(el, "Sec-WebSocket-Key:"){
			fmt.Print("\nforhåpentligvis Sec-WebSocket-Key:" +el+"\n\n")
			secWebSocketKey = strings.Trim(el, "Sec-WebSocket-Key: ")
		}
	}

	var body = bodyStart + bodyMid + bodyEnd
	var secWebSocketAcceptString = secWebSocketKey +"258EAFA5-E914-47DA-95CA-C5AB0DC85B11"
	var h = sha1.New()
	h.Write([]byte(secWebSocketAcceptString))

	var secWebSocketAccept = base64.StdEncoding.EncodeToString(h.Sum(nil))
	var header  = "HTTP/1.1 101 Switching Protocols\r\nUpgrade: websocket\r\nConnection: Upgrade\r\nSec-WebSocket-Accept: " + secWebSocketAccept +"\r\n\r\n"
	fmt.Println(header, body)
	//var total = header + body

	conn.Write([]byte(header)) // don't care about return value

	beskjed := []byte{0x81, 0x83, 0xb4, 0xb5, 0x03, 0x2a, 0xdc, 0xd0, 0x6}




	conn.Write(beskjed)


}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
