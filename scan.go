package main

import (
	"log"
	"net"
	"os"
	"strconv"
	"sync"
)

var wg sync.WaitGroup

func testTCPHandShake(host string, port int) {
	defer wg.Done()
	conn, err := net.Dial("tcp", host+":"+strconv.Itoa(port))
	if err != nil {
		return
	}
	log.Printf("- %d is open", port)
	conn.Close()
}

func main() {
	if len(os.Args) == 1 {
		log.Println("No args provided")
		os.Exit(1)
	}

	log.Println("Scanning host:")
	wg.Add(65535)
	for port := 1; port < 65536; port++ {
		go testTCPHandShake(os.Args[1], port)
	}
	wg.Wait()
	log.Println("done!")
}
