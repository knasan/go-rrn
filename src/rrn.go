package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// type repl (replication type)
type repl struct {
	from string
	to   string
}

var depthcount int
var replist []repl

// handler for panic
func handler() {
	if r := recover(); r != nil {
		fmt.Println(r)
		// flag.PrintDefaults()
	}
}

// replace all files (recursive)
func replace() {
	for _, p := range paths {
		r, err := run(p)
		if err != nil {
			panic(err)
		}
		reverse(r)
	}
}

// reverse replace with depth limit go backwards.
func reverse(r []repl) {
	c := len(r)
	for i := range r {
		cc := c - i - 1
		if verbose {
			fmt.Println("From: ", r[cc].from, "to:", r[cc].to)
		}
		if !dry {
			if err := os.Rename(r[cc].from, r[cc].to); err != nil {
				panic(err)
			}
		}
	}
}

// run fill the repl type
func run(path string) ([]repl, error) {
	list, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, d := range list {
		if d.IsDir() {
			depthcount += 1

			if depth != 0 {
				if depthcount == depth+1 {
					break
				}
			}
			_, err := run(filepath.Join(path, d.Name()))
			if err != nil {
				panic(err)
			}

		} else {
			from := filepath.Join(path, d.Name())
			to := filepath.Join(path, strings.ReplaceAll(d.Name(), searchChar, replaceChar))

			if from == to {
				continue
			}

			// fmt.Println("add - ", from, to)

			rpl := repl{from: from, to: to}
			replist = append(replist, rpl)
		}
	}

	return replist, err
}
