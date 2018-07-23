package main

// START 0 OMIT
import (
	"fmt"
	"log"
	"net"
	"net/rpc" // rpc standard lib package
	"reflect"

	"github.com/snassr/talk-0001-RPC/noderpc"
)

// END 0 OMIT

// START 1 OMIT

func main() {
	// create and register node object for RPC
	rpc.Register(&noderpc.Node{Name: "B"})

	// ...
	// END 1 OMIT

	// START 2 OMIT
	// create tcp address, bind to host on port 1200
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1200")
	if err != nil {
		log.Fatal(err)
	}
	// END 2 OMIT

	// START 3 OMIT
	// tcp network listener
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatal(err)
	}
	// END 3 OMIT

	fmt.Println("Server (Node B) started...")

	// START 4 OMIT
	for {
		// handle tcp client connections
		conn, err := listener.Accept()
		if err != nil {
			log.Println("listener accept error:", err)
		}
		// print connection info
		fmt.Println("received message", reflect.TypeOf(conn), &conn)

		// handle client connections via rpc
		rpc.ServeConn(conn)
	}
	// END 4 OMIT
}
