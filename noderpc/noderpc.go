package noderpc

import (
	"fmt"
)

// START 1 OMIT

// Node represents an actions node.
type Node struct {
	Name string
}

// START 0 OMIT

// Say is an RPC method that says something from a node.
func (n *Node) Say(args *string, reply *string) error {
	fmt.Printf("Say - args: %v\n", *args)
	*reply = fmt.Sprintf("Thanks, %s to you too! from node %s", *args, n.Name)
	return nil
}

// END 0 OMIT

// Square is an RPC method that squares a number.
func (n *Node) Square(args *int, reply *string) error {
	fmt.Printf("Square - args: %v\n", *args)
	squared := (*args) * (*args)
	*reply = fmt.Sprintf("Square of %d is %d executed on node %s", *args, squared, n.Name)
	return nil
}

// END 1 OMIT
