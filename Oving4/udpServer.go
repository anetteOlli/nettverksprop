package main

import (
	"bytes"
	"fmt"
"net"
	"os"
	"strconv"
	"strings"
)



func CheckError(err error)  {
	if(err!=nil){
		fmt.Printf("Fatal error %s", err.Error())
		os.Exit(1)
	}
	
}


func main() {
	service := ":1234"
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	CheckError(err)
	conn, err := net.ListenUDP("udp", udpAddr)
	CheckError(err)
	for {
		handleClient(conn)
	}
}

func handleClient(conn *net.UDPConn){
	buf := make([]byte, 512)
	tall1 := 0
	tall2 := 0
	tall3 := 0


	//deadline := time.Now().Add(10 * time.Second) //får 10 sekunder deadline på seg på å sende over alle 3 pakkene



	//henter ut tall1
	_, addr, err :=conn.ReadFromUDP(buf[0:])
	CheckError(err)
	n := bytes.IndexByte(buf, 0) //henter ut indexen til siste byte som ikke er null
	lest1 := string(buf[0:n])    //converterer bytearray fra index 0 til index n til en string
	if strings.Contains(lest1, "tallEN") {
		//fjerner ordet tall 1 og skal plukke ut selve tallet:
		fmt.Print(lest1 + " ")
		lest1 = strings.Trim(lest1, "tallEN") //fjerner "tall1" fra stringen
		fmt.Print(lest1 + " ")
		tall, err := strconv.Atoi(lest1)     //er det samme som Integer.parseInt() i java, gikk ikke an å skrive tall1 direkte, så dermed ny tall variabel her
		fmt.Print(strconv.Itoa(tall) + "\n")
		CheckError(err)
		tall1 = tall
	}

	//henter ut tall 2:
	_,addr, err = conn.ReadFromUDP(buf[0:])
	CheckError(err)
	n = bytes.IndexByte(buf, 0)
	lest1 = string(buf[0:n])
	if strings.Contains(lest1, "tallTO"){
		fmt.Print(lest1 + " ")
		lest1 = strings.Trim(lest1, "tallTO")
		fmt.Print(lest1 + " ")
		tall, err := strconv.Atoi(lest1)
		CheckError(err)
		fmt.Print(strconv.Itoa(tall) + "\n")
		tall2 = tall
	}

	//henter ut operatoren:
	_,addr, err = conn.ReadFromUDP(buf[0:])
	CheckError(err)
	n = bytes.IndexByte(buf, 0)
	lest1 = string(buf[0:n])
	if strings.Contains(lest1, "operatoren"){
		fmt.Print(lest1 + "\n")
		lest1 = strings.Trim(lest1, "operatoren")
		if strings.Contains(lest1, "ADD"){
			tall3 = tall1 + tall2
			fmt.Print(strconv.Itoa(tall3) + "\n")
			conn.WriteToUDP([]byte(strconv.Itoa(tall3)), addr)
			return
		}else if strings.Contains(lest1, "MINUS"){
			tall3 = tall1 - tall2
			conn.WriteToUDP([]byte(strconv.Itoa(tall3)), addr)
			return
		}
	}

}