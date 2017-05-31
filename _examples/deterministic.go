// +build ignore

package main

import (
	"fmt"
	"github.com/f483/dejavu"
)

func main() {

	// always remembers last 1024 entries
	d := dejavu.NewDeterministic(1024)

	fmt.Println(d.Witness([]byte("foo"))) // entry added
	fmt.Println(d.Witness([]byte("foo"))) // remembers entry
}
