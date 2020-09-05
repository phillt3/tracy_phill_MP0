package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

//Welcome to the client

func main() {
	/*main function responsible for gathering and writing message input to be sent to server (processB)
	-Ensures connection to server by given host port in the command line: example go run processA.go 1234
	-Does this by checking if port was provided and proceeds to attempt to connect with that port (net.Dial)
	^^^This is portion is mainly implemented from Linode.com exercise "Create a TCP and UDP Client and Server using Go"
	-loop does most of the work, reading in specific information one at a time and writing/sending it to the server
	-info includes (in order): Title of message, recipient (To), sender (From), date (Date), message (Content)
	*/
	arguments := os.Args //For collecting command-line input
	if len(arguments) == 1 {
		//checks case where host port not provided at end of run statement
		fmt.Println("Please provide host:port.")
		return
	}

	CONNECT := arguments[1]
	c, err := net.Dial("tcp", CONNECT)
	//connection made using given host port to server
	if err != nil {
		//if error is returned by net.Dial, is printed here
		fmt.Println(err)
		return
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Title: ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(c, text+"\n")

		//received, _ := bufio.NewReader(c).ReadString('\n')

		fmt.Print("To: ")
		text2, _ := reader.ReadString('\n')
		fmt.Fprintf(c, text2+"\n")

		//received, _ = bufio.NewReader(c).ReadString('\n')

		fmt.Print("From: ")
		text3, _ := reader.ReadString('\n')
		fmt.Fprintf(c, text3+"\n")

		//received, _ = bufio.NewReader(c).ReadString('\n')

		t := time.Now()
		fmt.Fprintf(c, t.Format(time.ANSIC)+"\n")

		//received, _ = bufio.NewReader(c).ReadString('\n')

		fmt.Print("Content: ")
		text5, _ := reader.ReadString('\n')
		fmt.Fprintf(c, text5+"\n")

		received, _ := bufio.NewReader(c).ReadString('\n')

		fmt.Print("->: " + received)

		fmt.Println("TCP client exiting...")
		return
	}
}
