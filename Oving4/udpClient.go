package main
import (
	"bytes"
	"fmt"
	"net"
	"bufio"
	"os"
	"strings"
)

func main() {
	cont  := true
	for  cont{
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("hva er tall 1?")
		tall1, _ := reader.ReadString('\n')
		tall1 = strings.TrimSuffix(tall1, "\n") //fjerner newline på slutten av stringen
		fmt.Print(tall1)

		fmt.Print("hva er tall 2?")
		tall2, _ := reader.ReadString('\n')
		tall2 = strings.TrimSuffix(tall2, "\n") //fjerner newline på slutten av stringen
		fmt.Print(tall2)

		fmt.Print("skriv inn ADD eller MINUS")
		operatoren, _ := reader.ReadString('\n')
		operatoren = strings.TrimSuffix(operatoren, "\n")
		fmt.Print(operatoren)

		resultat := sendMessageToServer(tall1, tall2, operatoren)
		resultat = strings.TrimSuffix(resultat,"\n")
		fmt.Print(resultat)
		fmt.Print("fortsette? J/N")
		fortsette, _ := reader.ReadString('\n')
		fortsette = strings.TrimSuffix(fortsette, "\n")
		fmt.Print(fortsette)
		if fortsette == "J"{
			cont = true
			fmt.Print("\nfortsette var true")
		}else{
			cont = false
		}
	}


}

func sendMessageToServer(tall1 string, tall2 string, operatoren string) string{
	p :=  make([]byte, 2048)
	conn, err := net.Dial("udp", "127.0.0.1:1234")
	if err != nil {
		fmt.Printf("Some error %v", err)
		return "error"
	}
	_, err = conn.Write([]byte("tall1" + tall1))
	checkError(err)
	_, err = conn.Write([]byte("tall2" + tall2))
	checkError(err)
	_, err = conn.Write([]byte("operatoren" + operatoren))
	checkError(err)
	//fmt.Fprintf(conn, []byte("tall1" + tall1))
	//fmt.Fprint(conn, "tall2" +tall2)
	//fmt.Fprint(conn, "operatoren" + operatoren)
	_, err = bufio.NewReader(conn).Read(p)
	if err == nil {
		n := bytes.IndexByte(p, 0) //henter ut indexen til siste byte som ikke er null
		fmt.Printf("%s\n\n", p[0:n])
	} else {
		fmt.Printf("Some error %v\n", err)
	}
	conn.Close()
	return ""
}

func checkError(e error) {
	if(e !=nil){
		fmt.Fprintf(os.Stderr, "Fatal error %s", e.Error())
		os.Exit(1)
	}
}
