package main

//This is the Server
import (
	"../message"
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
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

		if x == 4 {
			fmt.Print("Title: " + messrec.Title)
			fmt.Print("To: " + messrec.To)
			fmt.Print("From: " + messrec.From)
			fmt.Print("Date: " + messrec.Date)
			fmt.Print("Content: " + messrec.Content)
			fmt.Println("Exiting TCP server!")
			return
		}
	}
}
