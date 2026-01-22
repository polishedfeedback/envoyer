package transfer

import (
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
)

func SendFile(filePath, address string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return fmt.Errorf("error getting file info: %v", err)
	}

	fileName := filepath.Base(filePath)
	fileSize := fileInfo.Size()

	conn, err := net.Dial("tcp", address)
	if err != nil {
		return fmt.Errorf("error connecting to sender: %v", err)
	}
	defer conn.Close()

	fmt.Fprintf(conn, "%s\n", fileName)
	fmt.Fprintf(conn, "%d\n", fileSize)

	_, err = io.Copy(conn, file)
	if err != nil {
		return fmt.Errorf("error sending file: %v", err)
	}

	fmt.Printf("✓ Sent: %s (%d bytes)\n", fileName, fileSize)
	return nil

}
