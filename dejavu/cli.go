package main

import (
	"flag"
	"fmt"
	"github.com/f483/dejavu"
	"os"
)

const usageHeader string = `Usage: dejavu [OPTION]... [FILE]...

Concatenate FILE(s) and filter or output duplicate lines.

With no FILE, or when FILE is -, read standard input.

Options:
`

const usageFooter string = `
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
`

const helpDuplicates string = `output only duplicates instead of filtering`

const helpLimit string = `limit after which entries are forgotton`

const helpF string = `chance of false positive, between 0.0 and 1.0
	only for probabilistic mode`

const helpD string = `use deterministic mode instead of probabilistic
	WARNING requires order of magnitude more memory`

const helpVersion string = `output version information and exit`

const helpOutput string = `output file, defaults to stdout`

type options struct {
	limit         uint    // greater than 0
	fpRatio       float64 // between 0.0 and 1.0
	deterministic bool    // otherwise probabilistic
	duplicates    bool    // output duplicates instead of filtering
	version       bool    // show version and exit
	output        string  // output file, empty for stdout
}

func parseArgs() (options, []string) {
	var o options

	// set flags and default values
	flag.BoolVar(&o.duplicates, "d", false, helpDuplicates)
	flag.UintVar(&o.limit, "l", 1000000, helpLimit)
	flag.Float64Var(&o.fpRatio, "f", 0.000001, helpF)
	flag.BoolVar(&o.deterministic, "D", false, helpD)
	flag.BoolVar(&o.version, "v", false, helpVersion)
	flag.StringVar(&o.output, "o", "", helpOutput)

	// override default usage func
	flag.Usage = func() {
		fmt.Fprintf(os.Stdout, usageHeader)
		flag.PrintDefaults()
		fmt.Fprintf(os.Stdout, usageFooter)
	}

	// get flags and paths
	flag.Parse()
	paths := flag.Args()

	// read from stdin if no paths provided
	if len(paths) == 0 {
		paths = []string{"-"}
	}

	return o, paths
}

func main() {
	o, paths := parseArgs()

	// only print version
	if o.version {
		fmt.Println(dejavu.Version)
		return
	}

	// process data
	d := dejavu.New(!o.deterministic, uint32(o.limit), o.fpRatio)
	dejavu.ProcessPaths(d, !o.duplicates, o.output, paths...)
}
