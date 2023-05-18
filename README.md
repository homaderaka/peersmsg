# peersmsg

peersmsg is a Go package currently under active development, aimed at providing simple and flexible message parsing and validation for peer-to-peer communication. It introduces a set of interfaces and structures for message handling and parsing.

:construction: Project Status: In progress :construction:

The development of this project is ongoing, and the public API may change as the library matures. Please stay tuned for future updates.

## Installation
```
go get github.com/homaderaka/peersmsg
```

## Example Usage

Here is an example of how peersmsg could be used in a simple peer-to-peer (P2P) chat system:

```golang
// server.go
package main

import (
	"bufio"
	"fmt"
	"github.com/homaderaka/peersmsg"
	"net"
)

func main() {
	listener, _ := net.Listen("tcp", "localhost:8080")
	defer listener.Close()

	p := peersmsg.NewParser('\n', peersmsg.WithValidator(func(m peersmsg.Message) error {
		if len(m.Bytes()) > 512 {
			return fmt.Errorf("message too long")
		}
		return nil
	}))

	for {
		conn, _ := listener.Accept()
		go handleRequest(conn, p)
	}
}

func handleRequest(conn net.Conn, p peersmsg.Parser) {
	defer conn.Close()

	for {
		msg, err := p.NextMessage(conn)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Received:", msg.String())
	}
}

```

```golang
// client.go
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, _ := net.Dial("tcp", "localhost:8080")

	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			message := scanner.Text()
			fmt.Fprintf(conn, message+"\n")
		}
	}()

	buf := make([]byte, 512)
	for {
		n, _ := conn.Read(buf)
		fmt.Println("Received:", string(buf[:n]))
	}
}

```

In these examples:

server.go represents a basic peer-to-peer (P2P) node that listens for incoming TCP connections and processes messages using the peersmsg package. It prints each valid message to the console and discards messages that are too long.

client.go is a simple client that connects to the server and allows for messages to be sent from the console. It also listens for incoming messages from the server and prints them to the console.
