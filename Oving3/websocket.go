package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"

)


func main() {
	http.HandleFunc("/ws", wsHandler)
	http.HandleFunc("/", rootHandler)

	panic(http.ListenAndServe(":8080", nil))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {

	/**
	w setter automatisk status til 200 med mindre man spesifiserer på egenhånd
	w benytter text/html som default
	*/

	fmt.Fprintf(w, "%s", "<html><body><h1>Hilsen. Du har koblet deg opp til min enkle web-tjener</h1><ul>")


	for name, headers :=range r.Header {
		var data = ""
		for _, h := range headers{
			data +=h
		}
		fmt.Fprintf(w, "%s", "<li>" + name +" = " + data + "</li>")

	}
	fmt.Fprintf(w, "%s", "</ul></body></html>")

}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Origin") != "http://"+r.Host {
		http.Error(w, "Origin not allowed", 403)
		return
	}

	conn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
	}
	print(conn)
}

