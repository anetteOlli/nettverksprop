/* DaytimeServer
 */
package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
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
	fmt.Println(buf)
	conn.Read(buf[0:])
	fmt.Println(buf)
	var bufRead = string(buf[0:])
	fmt.Println(bufRead)

	temp := strings.Split(bufRead,"\n")
	var bodyMid= ""
	for _, element := range temp {
		var el = strings.TrimSpace(element)
		if len(el) !=0 && el != "  " {
			bodyMid += "<li>" + el + "</li>"
		}
	}

	var body = bodyStart + bodyMid + bodyEnd

	var header  = "HTTP/1.0 200 OK\r\nContent-Type: text/html\r\ncharset=UTF-8\r\nContent-length: " + strconv.Itoa(len(body)) + "\r\n\r\n"
	fmt.Println(header, body)
	var total = header + body

	conn.Write([]byte(total)) // don't care about return value
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
