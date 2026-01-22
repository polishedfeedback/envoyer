package main

import (
	"fmt"
	"log"
	"os"

	"github.com/polishedfeedback/envoyer/internal/transfer"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	command := os.Args[1]

	switch command {
	case "receive":
		port := ":8080"
		if (len(os.Args)) > 2 {
			port = os.Args[2]
		}

		err := transfer.StartReceiver(port)
		if err != nil {
			log.Fatal(err)
		}
	case "send":
		if len(os.Args) < 3 {
			fmt.Println("Error: filepath required")
			printUsage()
			return
		}

		filePath := os.Args[2]

		address := "localhost:8080"
		if len(os.Args) > 3 {
			address = os.Args[3]
		}

		err := transfer.SendFile(filePath, address)
		if err != nil {
			log.Fatal(err)
		}
	default:
		printUsage()
	}
}

func printUsage() {
	fmt.Println("LocalSend - File Transfer Tool")
	fmt.Println("\nUsage:")
	fmt.Println("  localsend receive [port]")
	fmt.Println("  localsend send <filepath> [address]")
	fmt.Println("\nExamples:")
	fmt.Println("  localsend receive")
	fmt.Println("  localsend receive :9090")
	fmt.Println("  localsend send file.txt")
	fmt.Println("  localsend send photo.jpg 192.168.1.5:8080")
}
