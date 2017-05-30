[![Build Status](https://travis-ci.org/f483/dejavu.svg)](https://travis-ci.org/f483/dejavu)
[![Coverage](https://coveralls.io/repos/f483/dejavu/badge.svg)](https://coveralls.io/r/f483/dejavu)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/f483/dejavu/master/LICENSE)
[![Issues](https://img.shields.io/github/issues/f483/dejavu.svg)](https://github.com/f483/dejavu/issues)
[![Go Report Card](https://goreportcard.com/badge/github.com/f483/dejavu)](https://goreportcard.com/report/github.com/f483/dejavu)


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

## Constant witness time: O(1)

Time to witness 1000000 data entries:

![Benchmark Time](https://github.com/f483/dejavu/raw/master/benchmark-time.png)


## Linear memory usage per stored entrie: O(n)

Memory usage for 
![Benchmark Memory](https://github.com/f483/dejavu/raw/master/benchmark-memory.png)

# Support

TODO donation text and bitcoin/paypal links
