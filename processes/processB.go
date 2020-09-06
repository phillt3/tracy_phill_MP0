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
	-info includes (in order): recipient (To), Title , sender (From), date (Date), message (Content)
	*/

	arguments := os.Args //for collecting command line run inputs (port number)
	if len(arguments) == 1 {
		//checking if port number was provided
		fmt.Println("Please provide port number")
		return
	}

	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp", PORT)
	//opening channel under given port number, making processB the server
	if err != nil {
		//if error is returned by net.Listen, it is printed
		fmt.Println(err)
		return
	}
	defer l.Close()
	//the closing statement for the program, will not be executed until all other operations are complete

	c, err := l.Accept()
	//only after a successful call to Accept() that the TCP server can begin to interact with TCP client (linode.com)
	if err != nil {
		//if error is returned by l.Accept, it is printed
		fmt.Println(err)
		return
	}

	var messrec message.Message //initializing message struct to be constructed by input sent by client

	for x := 0; x < 5; x++ {
		netData, err := bufio.NewReader(c).ReadString('\n') //reading in input sent to the server
		if err != nil {
			//if err is given an error value, it is printed here
			fmt.Println(err)
			return
		}

		//based on iteration of the loop, order is determined for how data sent to the server is interpreted
		//order is: Recipient then title then sender then date, then message
		if x == 0 {
			messrec.To = string(netData) //storing recipient of message in appropriate struct field
		} else if x == 1 {
			messrec.Title = string(netData) //storing Title of message in appropriate struct field
		} else if x == 2 {
			messrec.From = string(netData) //storing sender of message in appropriate struct field
		} else if x == 3 {
			messrec.Date = string(netData) //storing date and time of message in appropriate struct field
		} else if x == 4 {
			messrec.Content = string(netData) //storing content of message in appropriate struct field
		}

		if x == 4 {
			//final step is to take constructed message and print it to user on the server end
			//The message has been printed in both a neat formatted version
			//It has also can be printed in its original struct format
			//return statement allows previous close() call to be executed
			fmt.Print("To: " + messrec.To)
			fmt.Print("Title: " + messrec.Title)
			fmt.Print("From: " + messrec.From)
			fmt.Print("Date: " + messrec.Date)
			fmt.Print("Content: " + messrec.Content)
			//fmt.Printf("Here is message printed in struct format:\n%v\n", messrec)
			t := time.Now()
			myTime := "Message received at " + t.Format(time.RFC3339) + "\n"
			c.Write([]byte(myTime))
			//Above is the explicit ACK that will be seen by the client and user in the terminal
			//It will provide the time and date the date was received and completely constructed
			fmt.Println("Exiting ProcessB TCP server!")
			return
		}
	}
}
