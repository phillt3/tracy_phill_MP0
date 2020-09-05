package main

import (
	"../message"
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

//Welcome to the server

func main() {
	/*main function responsible for receiving and constructing a message sent by the client(processA)
	-Ensures connection to client by defining the port number to be used by client: example go run processB.go 1234
	-Checks if port was given and proceeds to attempt to "listen" for connection to that port by processA
	^^^This is portion is mainly implemented from Linode.com exercise "Create a TCP and UDP Client and Server using Go"
	-loop does most of the work, reading in data in specific order, using data to build instance of struct message
	-parts of struct message are then printed separately but just teh same can be printed all together
	-info includes (in order): Title of message, recipient (To), sender (From), date (Date), message (Content)
	*/

	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide port number")
		return
	}

	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	c, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	var messrec message.Message

	for x := 0; x < 5; x++ {
		netData, err := bufio.NewReader(c).ReadString('\n')

		if err != nil {
			fmt.Println(err)
			return
		}

		if x == 0 {
			messrec.Title = string(netData)
			//fmt.Print(messrec.Title)
		} else if x == 1 {
			messrec.To = string(netData)
			//fmt.Print(messrec.To)
		} else if x == 2 {
			messrec.From = string(netData)
			//fmt.Print(messrec.From)
		} else if x == 3 {
			messrec.Date = string(netData)
			//fmt.Print(messrec.Date)
		} else if x == 4 {
			messrec.Content = string(netData)
			//fmt.Print(messrec.Content)
		}

		t := time.Now()
		myTime := "Message received at " + t.Format(time.RFC3339) + "\n"
		c.Write([]byte(myTime))
		c.Write([]byte("!"))

		if x == 4 {
			fmt.Print("Title: " + messrec.Title)
			fmt.Print("To: " + messrec.To)
			fmt.Print("From: " + messrec.From)
			fmt.Print("Date: " + messrec.Date)
			fmt.Print("Content: " + messrec.Content)
			fmt.Printf("Here is message printed in struct format:\n%v\n", messrec)
			fmt.Println("Exiting ProcessB TCP server!")
			return
		}
	}
}
