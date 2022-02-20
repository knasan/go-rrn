package main

import (
	"bufio"
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
	var ign bool
	c := len(r)
	for i := range r {
		cc := c - i - 1
		if verbose {
			fmt.Println("From: ", r[cc].from, "to:", r[cc].to)
		}

		if interactive {
			fmt.Println("rename? ")
			reader := bufio.NewReader(os.Stdin)
			t, err := reader.ReadString('\n')
			t = strings.Trim(t, "\n")
			if err != nil {
				panic(err)
			}
			switch t {
			case "y", "Y":
				ign = false
			default:
				ign = true
			}
		} else {
			ign = false
		}

		if !dry {
			if !ign {
				if err := os.Rename(r[cc].from, r[cc].to); err != nil {
					panic(err)
				}
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
			// splits path by "/" (Unix), length + 1 for the specified directory gives the depth
			depthcount = len(strings.Split(path, "/")) + 1

			if depthcount > depth {
				if depth != 0 {
					continue
				}
			}

			_, err := run(filepath.Join(path, d.Name()))
			if err != nil {
				panic(err)
			}

		} else {

			// fmt.Println("File:", d.Name(), "depth:", depth, "depthcount:", depthcount, "path:", path)
			from := filepath.Join(path, d.Name())
			to := filepath.Join(path, strings.ReplaceAll(d.Name(), searchChar, replaceChar))

			if from == to {
				continue
			}

			rpl := repl{from: from, to: to}
			replist = append(replist, rpl)
		}
	}

	return replist, err
}
