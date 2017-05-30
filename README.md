[![Build Status](https://travis-ci.org/F483/dejavu.svg)](https://travis-ci.org/F483/dejavu)
[![Go Report Card](https://goreportcard.com/badge/github.com/f483/dejavu)](https://goreportcard.com/report/github.com/f483/dejavu)
[![Issues](https://img.shields.io/github/issues/f483/dejavu.svg)](https://github.com/f483/dejavu/issues)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/f483/dejavu/master/LICENSE)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/f483/dejavu)


# Déjà vu

Quickly detect already witnessed data.

Memory of witnessed data entries is limited and oldest are forgotten.
It may give false negatives, but not false positives. Library is thread
safe.


# Example

```
package main

import (
	"fmt"
	"github.com/f483/dejavu"
)

func main() {

	// remembers last three entries
	d := dejavu.NewDejaVu(3)

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
```


# Performance

### Constant witness time: O(1)

![Benchmark Time](https://github.com/f483/dejavu/raw/master/benchmark-time.png)

### Linear memory usage: O(n)

![Benchmark Memory](https://github.com/f483/dejavu/raw/master/benchmark-memory.png)



# Support

TODO donation text and bitcoin/paypal links
