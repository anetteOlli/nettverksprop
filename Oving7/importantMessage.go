package main

import (
	"encoding/hex"
	"fmt"
	"crypto/rsa"
	"crypto/rand"
	"encoding/base64"
	"crypto/sha512"
	"log"
	"encoding/pem"
	"crypto/x509"
	"encoding/asn1"

)

func main()  {
	//MESSAGE :="important message!"
	const SIGNATURE_text ="9c3e8d77333fcee3885747250fd48c8a6a5a8e62c24f8ef5f578c752469880409f69fa94a70dae0f71acc7a3988cc81e66881cbc75d5096dedfeeb3d17fb88fd27abe5d32f3b705a11045a91b5b5986f34948009e9b35e8026f986ae871e986392ae37e0458223d62b05fbb50935f63fa920590454d7851d35bf7b3d4cf0752c4683666bcb0398843d141113f32442f8d38f7910a43102da331a6e56fd2a3b3dbe49abf15b4e93c5a81341ed9f87e6bd972536e185e2cde096105db51de519f980901585b2c312b8a097853434bf144a3f14182f2d1b971169280b15061b781a21b8954c626aa4d9417275c1b1812eb0b9770b8320db2f1093f6e775105d39d5"
	const PUBLIC_KEY_text ="MIIBCgKCAQEA639u2haGdEoEQ5wf7lfTHEvDW2FuLBNmZgailV3N9L2JCI9NKtk1QOlEW2t6jweRfzjNf7Qs9XZkk6v6hveW2AZAYuhbNxQFT1FOk+Ez2RFVLLNZfIc+sXD0VURkORY7m+CFHfT+pf6hlLrvZONEWdJ1ZmxDtMOH6hTESCOooxdJ8m2+WsA5GuzOvaagZD/P4Gf9uoVjk/+G4jsB3YyaGAu+hs/Xx/ti9xPwFtCiUloJlUxhsDz9my67QMmPype4vv1w2Hhaj3UabCQi5qj4JgSctNayRy73Wk0iXtos1s2S38CUsUuSL7oZWDeIi2pZS0NT7e8cZllAHgSuX8MW+wIDAQAB"
	//må sjekke lengden på hashen og velg korrekt hash funksjon etterpå.
	var PUBLIC_KEY_hele = "-----BEGIN RSA PUBLIC KEY-----\n"+
						"MIIBCgKCAQEA639u2haGdEoEQ5wf7lfTHEvDW2FuLBNmZgailV3N9L2JCI9NKtk1"+
						"QOlEW2t6jweRfzjNf7Qs9XZkk6v6hveW2AZAYuhbNxQFT1FOk+Ez2RFVLLNZfIc+"+
						"sXD0VURkORY7m+CFHfT+pf6hlLrvZONEWdJ1ZmxDtMOH6hTESCOooxdJ8m2+WsA5"+
						"GuzOvaagZD/P4Gf9uoVjk/+G4jsB3YyaGAu+hs/Xx/ti9xPwFtCiUloJlUxhsDz9"+
						"my67QMmPype4vv1w2Hhaj3UabCQi5qj4JgSctNayRy73Wk0iXtos1s2S38CUsUuS"+
						"L7oZWDeIi2pZS0NT7e8cZllAHgSuX8MW+wIDAQAB"+
						"\n-----END RSA PUBLIC KEY-----"


	signature_as_byte,_ := hex.DecodeString(SIGNATURE_text)
	fmt.Print(signature_as_byte[0:], "\n")

	public_key_as_bytes,_ := base64.StdEncoding.DecodeString(PUBLIC_KEY_text)
	fmt.Print(public_key_as_bytes[0:], "\n")

	//Lager publickey objekt ut fra den gitte key-strengen
	block, _ := pem.Decode([]byte(PUBLIC_KEY_hele))
	var pk rsa.PublicKey
	asn1.Unmarshal(block.Bytes, &pk)
	fmt.Println(pk)

	fmt.Println("Exponent : ", pk.E)
	fmt.Println("Modulus : ", pk.N)


	//trinnn 1: hent hashen ut av signaturen


	//trinn 2: hash message med korrekt sha algoritme basert på lengden på svaret i trinn 1

	//sammenligne med hashen fra trinn 1 med hashen i trinn 2



	//oppgave 2:
	//er ikke så mye styr.
	//skal lage skjølsignert sertrifikat. Det vil jo ikke nettleseren godta men det er ok.
	//øvelsen er å sette opp med egne sertrifikater etc.
}

// BytesToPrivateKey bytes to private key
func BytesToPrivateKey(public_key_as_bytes []byte) *rsa.PrivateKey {
	block, _ := pem.Decode(public_key_as_bytes)
	enc := x509.IsEncryptedPEMBlock(block)
	b := block.Bytes
	var err error
	if enc {
		log.Println("is encrypted pem block")
		b, err = x509.DecryptPEMBlock(block, nil)
		if err != nil {
			fmt.Println(err)
		}
	}
	key, err := x509.ParsePKCS1PrivateKey(b)
	if err != nil {
		fmt.Println(err)
	}
	return key
}

// DecryptWithPrivateKey decrypts data with private key
func DecryptWithPrivateKey(signature_as_byte []byte, priv *rsa.PrivateKey) []byte {
	hash := sha512.New()
	plaintext, err := rsa.DecryptOAEP(hash, rand.Reader, priv, signature_as_byte, nil)
	if err != nil {
		fmt.Println(err)
	}
	return plaintext
}


