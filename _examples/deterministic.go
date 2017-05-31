// +build ignore

package main

import (
	"fmt"
	"github.com/f483/dejavu"
)

func main() {

	// remembers last three entries
	d := dejavu.NewDejaVuDeterministic(3)

	// add entries
	fmt.Println(d.Witness([]byte("foo")))
	fmt.Println(d.Witness([]byte("bar")))

	// remembers entry
	fmt.Println(d.Witness([]byte("bar")))

	// remembers oldest entry before overwriting
	fmt.Println(d.Witness([]byte("foo")))

	// add entries
	fmt.Println(d.Witness([]byte("bam")))
	fmt.Println(d.Witness([]byte("baz")))

	// forgot oldest
	fmt.Println(d.Witness([]byte("bar")))
}
