package main

import (
	tel "github.com/reiver/go-telnet"
	"log"
)

func main() {

	var handler tel.Handler = tel.EchoHandler

	err := tel.ListenAndServe(":555", handler)
	if nil != err {
		log.Fatal(err)
	}

}
