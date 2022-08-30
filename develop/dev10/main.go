package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

type client struct {
	timeout time.Duration
	host    string
	port    string
	con     net.Conn
	exit    chan os.Signal
	wg      sync.WaitGroup
}

func newClient() *client {
	client := &client{}
	client.parseFlags()
	client.exit = make(chan os.Signal)
	signal.Notify(client.exit, syscall.SIGTERM, syscall.SIGINT)
	return client
}
func (c *client) connect() {
	var err error
	c.con, err = net.DialTimeout("tcp", c.host+c.port, c.timeout)
	if err != nil {
		log.Fatal("Can't connect!")
	}

}
func (c *client) parseFlags() {
	var timeStr string
	flag.StringVar(&timeStr, "timeout", "10", "")
	flag.StringVar(&c.host, "host", "localhost", "")
	flag.StringVar(&c.port, "port", ":555", "")
	flag.Parse()
	timeNum, err := strconv.Atoi(timeStr)

	if err != nil {
		log.Fatal("timout format is wrong!")
	}
	c.timeout = time.Duration(timeNum) * time.Second
}
func (c *client) listen() {

	for {
		message, err := bufio.NewReader(c.con).ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("Server disconnected")
				c.wg.Done()
				c.exit <- syscall.SIGTERM
			}
			return
		}
		fmt.Print("Received: " + message)
	}
}

func (c *client) send() {
	read := bufio.NewReader(os.Stdin)
	for {
		message, err := read.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				c.wg.Done()
				c.exit <- syscall.SIGTERM
			}
			return
		}
		_, err = c.con.Write([]byte(message))
		if err != nil {
			fmt.Println(err)
		}
	}
}

func main() {
	client := newClient()
	client.connect()
	client.wg = sync.WaitGroup{}
	client.wg.Add(2)

	go client.send()
	go client.listen()

	go gracefulShutdown(client)
	client.wg.Wait()
}

func gracefulShutdown(c *client) {
	<-c.exit
	fmt.Println("Bye!")
	c.con.Close()
	os.Exit(0)
}
