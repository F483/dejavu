[![Build Status](https://travis-ci.org/F483/dejavu.svg)](https://travis-ci.org/f483/dejavu)
[![Coverage](https://coveralls.io/repos/F483/dejavu/badge.svg)](https://coveralls.io/r/F483/dejavu)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/F483/dejavu/master/LICENSE)
[![Issues](https://img.shields.io/github/issues/F483/dejavu.svg)](https://github.com/F483/dejavu/issues)

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

TODO benchmark/plot performance and memory usage


# Support

TODO donation text and bitcoin/paypal links
