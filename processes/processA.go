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
		fmt.Print("Title>> ")
		text, _ := reader.ReadString('\n')

		fmt.Fprintf(c, text+"\n")

		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client exiting...")
			return
		}

		received, _ := bufio.NewReader(c).ReadString('\n')

		fmt.Print(">>To ")
		text2, _ := reader.ReadString('\n')
		fmt.Fprintf(c, text2+"\n")

		received, _ = bufio.NewReader(c).ReadString('\n')

		fmt.Print(">>From ")
		text3, _ := reader.ReadString('\n')
		fmt.Fprintf(c, text3+"\n")

		received, _ = bufio.NewReader(c).ReadString('\n')

		fmt.Print(">>Date ")
		text4, _ := reader.ReadString('\n')
		fmt.Fprintf(c, text4+"\n")

		received, _ = bufio.NewReader(c).ReadString('\n')

		fmt.Print(">>Content ")
		text5, _ := reader.ReadString('\n')
		fmt.Fprintf(c, text5+"\n")

		received, _ = bufio.NewReader(c).ReadString('\n')

		fmt.Print("->: " + received)

		/*fmt.Print("->: " + received)
		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client exiting...")
			return
		}*/
	}
}
