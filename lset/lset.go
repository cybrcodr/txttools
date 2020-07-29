// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// The lset command line tool provides set operations between 2 files, where
// each input file is a set of line items. If a file has duplicated line items,
// those will be treated as the same value. Resulting values are printed to
// stdout where each value is a line.
package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/cybrcodr/txttools/lset/internal/set"
)

func main() {
	if len(os.Args) != 4 {
		usage()
		os.Exit(1)
	}
	var cmd func(set.Strings, set.Strings)
	switch os.Args[1] {
	case "diff":
		cmd = diff
	case "cross":
		cmd = intersect
	case "minus":
		cmd = subtract
	default:
		fmt.Fprintf(os.Stderr, "Invalid command %q\n", os.Args[1])
		usage()
		os.Exit(1)
	}

	set1, err := readFile(os.Args[2])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	set2, err := readFile(os.Args[3])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	cmd(set1, set2)
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s <command> <file1> <file2>\n", filepath.Base(os.Args[0]))
	fmt.Fprintln(os.Stderr)
	fmt.Fprintln(os.Stderr, "where <command> is one of:")
	fmt.Fprintln(os.Stderr, "\tdiff  - shows unique lines that are different between file1 (-) and file2 (+)")
	fmt.Fprintln(os.Stderr, "\tminus - shows unique lines in file1 that are not in file2")
	fmt.Fprintln(os.Stderr, "\tcross - shows unique lines that are common between file1 and file2")
	fmt.Fprintln(os.Stderr)
}

func readFile(filename string) (set.Strings, error) {
	s := set.Strings{}
	f, err := os.Open(filename)
	if err != nil {
		return set.Strings{}, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s.Add(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return set.Strings{}, err
	}
	return s, nil
}

// diff prints the difference between s1 (-) and s2 (+).
func diff(s1, s2 set.Strings) {
	adds, subs := s1.Diff(s2)
	for _, val := range subs.SortedSlice() {
		fmt.Printf("-%s\n", val)
	}
	for _, val := range adds.SortedSlice() {
		fmt.Printf("+%s\n", val)
	}
}

// subtract prints out values in s1 that are not in s2.
func subtract(s1, s2 set.Strings) {
	res := s1.Minus(s2)
	for _, val := range res.SortedSlice() {
		fmt.Println(val)
	}
}

// intersect prints out values that are in both s1 and s2.
func intersect(s1, s2 set.Strings) {
	for _, val := range s1.Intersect(s2).SortedSlice() {
		fmt.Println(val)
	}
}
