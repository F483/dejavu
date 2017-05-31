[![Build Status](https://travis-ci.org/F483/dejavu.svg)](https://travis-ci.org/F483/dejavu)
[![Go Report Card](https://goreportcard.com/badge/github.com/f483/dejavu)](https://goreportcard.com/report/github.com/f483/dejavu)
[![Issues](https://img.shields.io/github/issues/f483/dejavu.svg)](https://github.com/f483/dejavu/issues)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/f483/dejavu/master/LICENSE)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/f483/dejavu)
[![Donate PayPal](https://img.shields.io/badge/Donate-PayPal-orange.svg)](https://www.paypal.com/cgi-bin/webscr?cmd=_donations&business=fabian%2ebarkhau%40gmail%2ecom&lc=DE&item_name=https%3a%2f%2fgithub%2ecom%2fF483%2fdejavu&no_note=0&currency_code=EUR&bn=PP%2dDonationsBF%3abtn_donateCC_LG%2egif%3aNonHostedGuest)
[![Donate Bitcoin](https://img.shields.io/badge/Donate-Bitcoin-orange.svg)](https://blockchain.info/address/13nAHLVo5GRdwVeLxEjbgEvyusrjdQogdD)


# Déjà vu

Quickly detect already witnessed data.

Limited memory of witnessed data, oldest are forgotten. Library is thread safe.


# Example

```
package main

import (
	"fmt"
	"github.com/f483/dejavu"
)

func main() {

	// always remembers last three entries
	d := dejavu.NewDeterministic(3)

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

![Benchmark Time](https://github.com/f483/dejavu/raw/master/_benchmark/time.png)

### Linear memory usage: O(n)

![Benchmark Memory](https://github.com/f483/dejavu/raw/master/_benchmark/memory.png)
