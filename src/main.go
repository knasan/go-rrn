package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	version                  = "0.0.2"
	author                   = "smk (github@knasan.de)"
	license                  = "MIT"
	paths                    pathslice
	searchChar, replaceChar  string
	verbose, dry, sl, sa, sv bool
	depth                    int
)

// usage
var usage = func() {
	_, _ = fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])

	flag.PrintDefaults()
	os.Exit(0)
}

// define a type named pathslice as slice of strings
type pathslice []string

// String
// flag.Value interface - return string
func (p *pathslice) String() string {
	return fmt.Sprintf("%v", *p)
}

// Set
// flag.Value - append to type pathslice return error
func (p *pathslice) Set(v string) error {
	*p = append(*p, v)
	return nil
}

// display author
func sauthor() {
	fmt.Println("author:", author)
}

// display version
func sversion() {
	fmt.Println("version:", version)
}

// display license
func slicense() {
	fmt.Println("license:", license)
}

// initialize flags
func initialize() {
	// append path
	flag.Var(&paths, "d", "working directory (can be specified multiple times)")

	// find char
	flag.StringVar(&searchChar, "s", "", "char to search - default empty string")

	// replace char
	flag.StringVar(&replaceChar, "r", "", "char to replace - default empty string")

	// verbose
	flag.BoolVar(&verbose, "v", false, "verbose")

	// dry run
	flag.BoolVar(&dry, "D", false, "dry-run")

	// max-depth
	flag.IntVar(&depth, "depth", 0, "max depth (0 = disable)")

	// show
	// author
	flag.BoolVar(&sa, "author", false, "show author (and others defined) and exit")
	// licnese
	flag.BoolVar(&sl, "license", false, "show license (and others defined) and exit")
	// version
	flag.BoolVar(&sv, "version", false, "show version (and others defined) and exit")

	flag.Parse()

	// usage and exit when no options
	if flag.NFlag() == 0 {
		usage()
	}

	// use one or multiple show (version, author, license), set show to true and exit this application
	show := false
	if sa {
		sauthor()
		show = true
	}
	if sl {
		slicense()
		show = true
	}
	if sv {
		sversion()
		show = true
	}

	if show {
		os.Exit(0)
	}

	if replaceChar == searchChar {
		panic("search and replace is the same character")
	}

	if replaceChar == "" {
		replaceChar = " "
	}

	if searchChar == "" {
		searchChar = " "
	}

	if len(paths) == 0 {
		panic("required argument path (-d) is missing")
	}
}

// main
func main() {
	defer handler()
	initialize()

	replace()
}
