package transfer

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

func StartReceiver(port string) error {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		return fmt.Errorf("error listening on port: %v", err)
	}

	fmt.Printf("Listening on %s\n", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Errorf("error accepting connection: %v", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	fileName, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("error reading file name: %v", err)
	}
	fileName = strings.TrimSpace(fileName)

	sizeStr, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading filename: %v\n", err)
		return
	}
	sizeStr = strings.TrimSpace(sizeStr)
	fileSize, err := strconv.ParseInt(sizeStr, 10, 64)

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close()

	_, err = io.CopyN(file, reader, fileSize)
	if err != nil {
		fmt.Printf("Error receiving file: %v\n", err)
		return
	}

	fmt.Printf("✓ Received: %s (%d bytes)\n", fileName, fileSize)
}
