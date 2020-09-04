package main

//This is the Client
import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	//"../message"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide host:port.")
		return
	}

	CONNECT := arguments[1]
	c, err := net.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Name>> ")
		text, _ := reader.ReadString('\n')

		fmt.Fprintf(c, text+"\n")

		received, _ := bufio.NewReader(c).ReadString('\n')

		fmt.Print("->: " + received)
		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client exiting...")
			return
		}

		fmt.Print("Title>> ")
		text2, _ := reader.ReadString('\n')
		fmt.Fprintf(c, text2+"\n")

		received, _ = bufio.NewReader(c).ReadString('\n')

		fmt.Print("->: " + received)
		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client exiting...")
			return
		}
	}
}
