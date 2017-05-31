// +build ignore

package main

import (
	"fmt"
	"github.com/f483/dejavu"
)

func main() {

	// probably remembers last 65536 with 0.000001 chance of false positive
	p := dejavu.NewProbabilistic(65536, 0.000001)

	fmt.Println(p.Witness([]byte("bar"))) // entry added
	fmt.Println(p.Witness([]byte("bar"))) // probably remembers entry
}
