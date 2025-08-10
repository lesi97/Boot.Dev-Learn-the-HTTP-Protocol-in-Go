package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"strings"

	"boot.dev/httpfromtcp/internal/request"
)

const inputFilePath = "messages.txt"

func main() {

	r, err := request.RequestFromReader(strings.NewReader("GET / HTTP/1.1\r\nHost: localhost:42069\r\nUser-Agent: curl/7.81.0\r\nAccept: */*\r\n\r\n"))
	if err != nil {
		log.Fatalf("ERROR - RequestFromReader: %v\n", err.Error())
	}
	fmt.Printf("Response: %+v\n", r)

	
	r, err = request.RequestFromReader(strings.NewReader("GET /coffee HTTP/1.1\r\nHost: localhost:42069\r\nUser-Agent: curl/7.81.0\r\nAccept: */*\r\n\r\n"))
	if err != nil {
		log.Fatalf("ERROR - RequestFromReader: %v\n", err.Error())
	}
	fmt.Printf("Response: %+v\n", r)

	_, err = request.RequestFromReader(strings.NewReader("/coffee HTTP/1.1\r\nHost: localhost:42069\r\nUser-Agent: curl/7.81.0\r\nAccept: */*\r\n\r\n"))
	if err != nil {
		log.Fatalf("ERROR - RequestFromReader: %v\n", err.Error())
	}

	// listener, err := net.Listen("tcp", ":42069")
	// if err != nil {
	// 	log.Fatalf("failed to listen on port 42060")
	// }
	// defer listener.Close()

	// for {
	// 	conn, err := listener.Accept()
	// 	if err != nil {
	// 		log.Fatalf("failed to accept connection: %v\n", err)
	// 	}
	// 	// fmt.Printf("Connection success\n")
	// 	go func(c net.Conn) {
	// 		linesChan := getLinesChannel(io.ReadCloser(c))
	// 		for line := range linesChan {
	// 			fmt.Println(line)
	// 		}
	// 		c.Close()
	// 	}(conn)
	// }

	// f, err := os.Open(inputFilePath)
	// if err != nil {
	// 	log.Fatalf("could not open %s: %s\n", inputFilePath, err)
	// }

	// fmt.Printf("Reading data from %s\n", inputFilePath)
	// fmt.Println("=====================================")

	// linesChan := getLinesChannel(f)

	// for line := range linesChan {
	// 	fmt.Println("read:", line)
	// }
}

func getLinesChannel(f io.ReadCloser) <-chan string {
	lines := make(chan string)
	go func() {
		defer f.Close()
		defer close(lines)
		currentLineContents := ""
		for {
			b := make([]byte, 8, 8)
			n, err := f.Read(b)
			if err != nil {
				if currentLineContents != "" {
					lines <- currentLineContents
				}
				if errors.Is(err, io.EOF) {
					break
				}
				fmt.Printf("error: %s\n", err.Error())
				return
			}
			str := string(b[:n])
			parts := strings.Split(str, "\n")
			for i := 0; i < len(parts)-1; i++ {
				lines <- fmt.Sprintf("%s%s", currentLineContents, parts[i])
				currentLineContents = ""
			}
			currentLineContents += parts[len(parts)-1]
		}
	}()
	return lines
}