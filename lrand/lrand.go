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

// The lrand command selects up to given n number of line items from a given
// list of line items either in a file or from stdin.
//
// TODO: Current output does not preserve the order of the input. It may be nice
// to have that.
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s <count> <file>\n", filepath.Base(os.Args[0]))
	fmt.Fprintln(os.Stderr)
	fmt.Fprintln(os.Stderr, "Output to stdout up to <count> number of line items from given <file>.")
	fmt.Fprintln(os.Stderr, "If file is '-', it reads from stdin.")
	fmt.Fprintln(os.Stderr)
}

func main() {
	if len(os.Args) != 3 {
		usage()
		os.Exit(1)
	}

	count, err := strconv.ParseUint(os.Args[1], 10, 64)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	items, err := readFile(os.Args[2])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	size := len(items)
	if count >= uint64(size) {
		fmt.Fprintln(os.Stderr, "count greater or equal to number of line items")
		os.Exit(1)
	}

	rand.Seed(time.Now().UnixNano())

	for i := uint64(0); i < count; i++ {
		idx := rand.Intn(size)
		fmt.Println(items[idx])
		items[idx] = items[size-1]
		size--
	}
}

func readFile(filename string) ([]string, error) {
	f := os.Stdin
	if filename != "-" {
		var err error
		if f, err = os.Open(filename); err != nil {
			return nil, err
		}
		defer f.Close()
	}
	scanner := bufio.NewScanner(f)

	items := []string{}
	for scanner.Scan() {
		items = append(items, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
