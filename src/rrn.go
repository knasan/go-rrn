package main

import (
	"flag"
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

// handler for panic
func handler() {
	if r := recover(); r != nil {
		fmt.Println(r)
		flag.PrintDefaults()
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
	var limit int
	for i := range r {
		cc := c - i - 1
		fi, err := os.Stat(r[cc].from)
		if err == nil {
			if fi.IsDir() {
				limit += 1
			}
		}
		if limit == depth {
			break
		}
		if verbose {
			fmt.Println("From: ", r[cc].from, "to:", r[cc].to)
		}
		if err := os.Rename(r[cc].from, r[cc].to); err != nil {
			panic(err)
		}
	}
}

// run fill the repl type
func run(path string) (replist []repl, err error) {
	err = filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			first := strings.TrimRight(filepath.Clean(path), info.Name())
			last := path[strings.LastIndex(path, "/")+1:]
			last = strings.ReplaceAll(last, searchChar, replaceChar)
			to := filepath.Join(first, last)
			if path != to {
				r := repl{from: path, to: to}
				replist = append(replist, r)
			}
			return nil
		})
	return replist, err
}
