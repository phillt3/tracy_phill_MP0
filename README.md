# tracy_phill_MP0
Send messages from one node to another utilizing Go packages working with TCP clients and servers.

**Citation Acknowledgment:**
* https://www.linode.com/docs/development/go/developing-udp-and-tcp-clients-and-servers-in-go/
    * Sourced directly for necessary components in developing connections between TCP client and server, specifically .listen, .dial, and channel creation functions which allow for the connection b/w processA and processB.
* https://dave.cheney.net/practical-go/presentations/qcon-china.html#_comments_on_variables_and_constants_should_describe_their_contents_not_their_purpose
    * Information used in order to reference correct commenting format and style
    
* https://www.geeksforgeeks.org/
    * Sourced indirectly for information on use and implementation of fmt.FPrintf(), which formats according to a specifier and then writes to given io.Writer
    
* https://github.com/Dariusrussellkish/Example-MP-Solution
    * sourced indirectly for information on how to develop and organize project (including package and variable names)
    
* golang.org
    * indirectly and directly sourced for playground testing and introductory guidance on certain golang features (how to run, compile, import packages etc)
  
# To Run
This project is run completely from the terminal, specifically two separate tabs, one for the server and one for the client.
First off make sure both terminals are working within the 'processes' directory within tracy_phill_MP0. Here is an example of what my working directory looks like from the command line.

```bash
/Users/philliptracy/Desktop/DistSystems/tracy_phill_MP0/processes
``` 
This must be applied to both terminals (I used the cd command). An example:

```bash
 cd /Users/philliptracy/Desktop/DistSystems/tracy_phill_MP0/processes
``` 
First run the server (processB) and provide port number (this case 1234):

```bash
go run processB.go 1234
``` 

In other terminal, connect the client (processA) to the server with this command:

```bash
go run processA.go 127.0.0.1:1234
``` 
Within the client terminal you will be prompted to input 4 items, pressing 'Enter' after each input:
* The recipient's email
* the title of the email
* the sender (you)
* the content

The date is being gathered and formatted automatically. Once the content of the message is provided and final 'Enter' pressed, the program will construct message on the server side and send an ACK back to the client, where both will terminate (specifcs in **Structure and Design**).

Sample Output:

* ProcessB initial startup, construction of message, and termination

```bash
Phillips-MacBook-Pro:processes philliptracy$ go run processB.go 1234
To: friend@bc.edu
Title: Meetup?
From: tracypb@bc.edu
Date: Sun Sep  6 12:19:58 2020
Content: Would you like to grab a bite soon?
Exiting ProcessB TCP server!
``` 

* ProcessA input and resulting ACK + termination

```bash
Phillips-MacBook-Pro:processes philliptracy$ go run processA.go 127.0.0.1:1234
>>To: friend@bc.edu
>>Title: Meetup?
>>From: tracypb@bc.edu
>>Content: Would you like to grab a bite soon?
->: Message received at 2020-09-06T12:20:08-04:00
ProcessA TCP client exiting...
``` 
# Structure and Design

### Message
Within `/message/message.go` there is a very simple struct and potential further implementation functions.
It contains the fields:

```bash
type Message struct {
	To      string
	From    string
	Title   string
	Content string
	Date    string
```
The date, although defined as a string, is gathered and formatted automatically by the program in RFC3339 and ANSIC format.

### Processes
Code for processB.go (the TCP server program) and processA.go (the TCP client program)
can both be found in `/processes/` directory.
The two communicate via functions and tools within go's 
"net" package. ProcessB opens the channel and 'listens' for a client to connect with it,
processA. ProcessA dials into the channel using the corresponding port number
where it then directs the user to start entering in specific message information which is then received and constructed using the message struct in ProcessB.

