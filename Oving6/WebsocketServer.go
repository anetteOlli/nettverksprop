package main

import (
	"math/rand"
	"fmt"
	"time"
	b64 "encoding/base64"
)

func main() {

	//Create a random 16byte number
	s1 := rand.NewSource(time.Now().UnixNano())
	token := make([]byte, 16)
	rand.New(s1).Read(token)
	fmt.Println(token)
	//b64 encode
	sEnc := b64.StdEncoding.EncodeToString(token)
	fmt.Println(sEnc)
	//b64 decode
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(sDec)
}
