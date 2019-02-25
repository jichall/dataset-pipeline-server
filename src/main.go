package main

import (
	"flag"
	"log"
)

var (
	addr   string
	port   string
	deploy bool
	debug  bool
)

func init() {
	flag.StringVar(&addr, "a", "127.0.0.1", "Set the address of the server, if deploy variable is true this variable is ignored.")
	flag.StringVar(&port, "p", "4000", "Set the port of the server.")
	flag.BoolVar(&deploy, "d", false, "If the deploy variable is set, the ip address of the server will be set to 0.0.0.0 to make it available to everyone in the network")
	flag.BoolVar(&debug, "v", false, "Shows verbose output")
}

func main() {
	flag.Parse()

	if deploy {
		log.Print("[+] Server running on 0.0.0.0:" + port + "\n")
		InitServer("0.0.0.0", port)
	}

	log.Print("[+] Initializing database\n")
	InitDatabase()
	log.Print("[+] Initializing server on " + addr + ":" + port + "\n")
	InitServer(addr, port)

}
