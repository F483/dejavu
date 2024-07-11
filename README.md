[![Build Status](https://travis-ci.org/F483/dejavu.svg)](https://travis-ci.org/F483/dejavu)
[![Issues](https://img.shields.io/github/issues/f483/dejavu.svg)](https://github.com/f483/dejavu/issues)
[![Go Report Card](https://goreportcard.com/badge/github.com/f483/dejavu)](https://goreportcard.com/report/github.com/f483/dejavu)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/f483/dejavu/master/LICENSE)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/f483/dejavu)


# Déjà vu

Quickly detect already witnessed data, ideal for deduplication.

Limited memory of witnessed data, oldest are forgotten. Library is thread safe.
Offers deterministic and probabilistic (over an order of magnitude less memory
consuming) implementation. The probabilistic implementation uses bloom filters,
meaning false positives are possible but not false negatives.


# Installation

## Download binary release

Compiled binaries for many platforms are available and can be downloaded for
the [latest release](https://github.com/F483/dejavu/releases/latest).

Extract the binary for your platform and add it to your system path.

## Compile from source

Requires golang [environment/workspace](https://golang.org/doc/code.html).

```
# compile and install library
go get github.com/f483/dejavu

# compile and install binary
go install github.com/f483/dejavu/dejavu
```

# Command line usage

```
$ dejavu -h
Usage: dejavu [OPTION]... [FILE]...

Concatenate FILE(s) and filter or output duplicate lines.

With no FILE, or when FILE is -, read standard input.

Options:
  -D	use deterministic mode instead of probabilistic
	WARNING requires order of magnitude more memory
  -d	output only duplicates instead of filtering
  -f float
    	chance of false positive, between 0.0 and 1.0
	only for probabilistic mode (default 1e-06)
  -l uint
    	limit after which entries are forgotton (default 1000000)
  -o string
    	output file, defaults to stdout
  -v	output version information and exit

Examples:
  dejavu
	default probabilistic deduplication from stdin to std out with
	1mil entry limit and 1/1mil chance of false positive (~8M mem usage)
  dejavu -o s f - g
	deduplicat f, then stdin, then g, to output s
  dejavu -l 10000000 -fp 0.000000001
	probabilistic deduplication with 10mil entry limit
	and 1/1bil chance of false positive (~70M mem usage)
  dejavu -d -D -l 65536
	output duplicates and avoid false positives with deterministic mode
	lower entry limit to avoid excessive memory usage

Implementation:
  Efficient probabilistic and deterministic duplicate detection with O(1) 
  detection time and O(n) memory usage in relation to entry limit. Default
  probabilistic implementation uses bloom filters, meaning false
  positives are possible but not false negatives.

Author: Fabian Barkhau <f483@protonmail.com>
Project: https://github.com/f483/dejavu
License: MIT https://raw.githubusercontent.com/f483/dejavu/master/LICENSE
```

# Library usage (golang)

## Probabilistic example

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

## Deterministic example

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
