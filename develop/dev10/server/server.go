package main

import (
	tel "github.com/reiver/go-telnet"
	"time"
)

func main() {

	var handler tel.Handler = tel.EchoHandler

	err := tel.ListenAndServe(":555", handler)
	if nil != err {
		panic(err)
	}

}

func send() {
	//conn := tel.
	for {
		time.Sleep(time.Second)

	}
}
