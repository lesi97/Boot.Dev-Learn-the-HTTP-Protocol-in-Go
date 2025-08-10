package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func upsender() {
	u, err := net.ResolveUDPAddr("udp", "localhost:42069")
	if err != nil {
		log.Fatalf("ERROR - ResolveUDPAddr: %v\n", err.Error())
	}

	conn, err := net.DialUDP("udp", nil, u)
	if err != nil {
		log.Fatalf("ERROR - DialUDP: %v\n", err.Error())
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		
		str, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("ERROR - ReadLine: %v\n", err.Error())
		}
		conn.Write([]byte(str))
	}
}