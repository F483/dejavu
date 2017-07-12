[![Build Status](https://travis-ci.org/F483/dejavu.svg)](https://travis-ci.org/F483/dejavu)
[![Issues](https://img.shields.io/github/issues/f483/dejavu.svg)](https://github.com/f483/dejavu/issues)
[![Go Report Card](https://goreportcard.com/badge/github.com/f483/dejavu)](https://goreportcard.com/report/github.com/f483/dejavu)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/f483/dejavu/master/LICENSE)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/f483/dejavu)
[![Donate PayPal](https://img.shields.io/badge/Donate-PayPal-ff69b4.svg)](https://www.paypal.com/cgi-bin/webscr?cmd=_donations&business=fabian%2ebarkhau%40gmail%2ecom&lc=DE&item_name=https%3a%2f%2fgithub%2ecom%2fF483%2fdejavu&no_note=0&currency_code=EUR&bn=PP%2dDonationsBF%3abtn_donateCC_LG%2egif%3aNonHostedGuest)
[![Donate Bitcoin](https://img.shields.io/badge/Donate-Bitcoin-ff69b4.svg)](https://blockchain.info/address/13nAHLVo5GRdwVeLxEjbgEvyusrjdQogdD)
[![Avaivable For Hire](https://img.shields.io/badge/Available-For_Hire-ff69b4.svg)](https://f483.github.io)


# Déjà vu

Quickly detect already witnessed data, ideal for deduplication.

Limited memory of witnessed data, oldest are forgotten. Library is thread safe.
Offers deterministic and probabilistic (over an order of magnatude less memory
consuming) implementation. The probabilistic implementation uses bloom filters,
meaning false positives are possible but not false negatives.

# Command line usage

The command line interface/behaviour is similar to that of `cat` and is 
intended to be used as a drop in replacement for deduplication. It 
offers additional arguments to control removal/highlighing of duplicates.

TODO examples

# Library usage (golang)

## Probabilistic

```
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
```

## Deterministic

```
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
```

# Performance

## Linear memory usage: O(n)

### Probabilistic

0.000001 chance of false positive.

![Benchmark Memory](https://github.com/f483/dejavu/raw/master/_benchmark/probabilistic-memory.png)

### Deterministic

![Benchmark Memory](https://github.com/f483/dejavu/raw/master/_benchmark/deterministic-memory.png)


## Constant witness time: O(1)

### Probabilistic

0.000001 chance of false positive.

![Benchmark Time](https://github.com/f483/dejavu/raw/master/_benchmark/probabilistic-time.png)

### Deterministic

![Benchmark Time](https://github.com/f483/dejavu/raw/master/_benchmark/deterministic-time.png)
