package main

// START 0 OMIT
import (
	"fmt"
	"log"
	"net/rpc" // rpc standard lib package
	"time"
)

// END 0 OMIT

// START 1 OMIT
func main() {
	// connect to remote server via rpc tcp
	client, err := rpc.Dial("tcp", ":1200")
	defer client.Close()
	if err != nil {
		log.Fatal(err)
	}
	// ...
	// END 1 OMIT

	fmt.Println("Client (Node A) running...")

	// action variables
	var serviceMethod string
	var argument interface{}

	// reply variable
	var reply string

	// --
	// RPC Call (1)
	// START 2 OMIT
	serviceMethod = "Node.Say"
	argument = "Hello World!"
	err = client.Call(serviceMethod, argument, &reply)
	if err != nil {
		log.Println(serviceMethod, err)
	}
	fmt.Printf("%v\n", reply)
	// END 2 OMIT

	// sleep for 2 seconds
	time.Sleep(2 * time.Second)

	// --
	// RPC Call (2)
	serviceMethod = "Node.Square"
	argument = 2
	err = client.Call(serviceMethod, argument, &reply)
	if err != nil {
		log.Println(serviceMethod, err)
	}
	fmt.Printf("%v\n", reply)

	fmt.Println("Client (Node A) done.")
}
