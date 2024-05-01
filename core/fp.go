package core

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func handleClient(client net.Conn) {
	defer client.Close()

	// Read the initial request line
	bufferedReader := bufio.NewReader(client)
	requestLine, err := bufferedReader.ReadString('\n')
	if err != nil {
		log.Println("Error reading request line:", err)
		return
	}
	if strings.HasPrefix(requestLine, "CONNECT ") {
		log.Println("Handling CONNECT")
		// Handle CONNECT for HTTPS
		handleConnect(client, bufferedReader, requestLine)
		return
	}
	log.Printf("Request line is %s", requestLine)
	// Here, instead of copying to remote, you'd typically continue the TLS handshake or forward as necessary
	// This is a placeholder for whatever operation you need with the SNI
}

func handleConnect(client net.Conn, reader *bufio.Reader, requestLine string) {
	// Extract the host and port from the request line
	hostPort := strings.Split(strings.TrimSpace(strings.TrimPrefix(requestLine, "CONNECT ")), " ")[0]
	destConn, err := net.Dial("tcp", hostPort)
	log.Printf("Connecting to %s", hostPort)
	if err != nil {
		log.Println("Error connecting to destination:", err)
		return
	}
	log.Printf("Connected to %s", hostPort)
	defer destConn.Close()

	log.Println("Sending back 200 OK to the client")
	// Send back a 200 OK to the client
	_, err = client.Write([]byte("HTTP/1.1 200 Connection established\r\n\r\n"))
	if err != nil {
		log.Println("Error sending connection confirmation:", err)
		return
	}
	log.Println("Sent 200 OK to the client")

	// Start tunnel - bidirectional copy
	// Start bidirectional copy
	done := make(chan error, 2)
	go func() {
		_, err := io.Copy(destConn, client)
		done <- err
	}()

	go func() {
		_, err := io.Copy(client, destConn)
		done <- err
	}()

	// Wait for both copies to complete
	for i := 0; i < 2; i++ {
		if err := <-done; err != nil {
			log.Printf("Error during tunneling: %v", err)
		}
	}
	log.Println("Tunneling completed for both directions")
}
func readClientHello(reader io.Reader) (*tls.ClientHelloInfo, error) {
	var hello *tls.ClientHelloInfo
	err := tls.Server(readOnlyConn{reader: reader}, &tls.Config{
		GetConfigForClient: func(argHello *tls.ClientHelloInfo) (*tls.Config, error) {
			hello = new(tls.ClientHelloInfo)
			*hello = *argHello
			return nil, nil
		},
	}).Handshake()

	if hello == nil {
		return nil, err
	}

	return hello, nil
}

func StartForwardProxy(hostname, port string) {
	address := fmt.Sprintf("%s:%s", hostname, port)
	log.Printf("Forward proxy address: %s", address)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen on %s: %v", address, err)
	}
	defer listener.Close()
	log.Printf("Proxy listening on %s", address)

	for {
		client, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %s\n", err)
			continue
		}

		go handleClient(client)
	}
}
