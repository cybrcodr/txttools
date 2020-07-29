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

// The ljoin command joins lines in the given file with a separator into a
// single line output.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var separator = flag.String("s", " ", "separator string between values")

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [-s <separator>] <file>\n\n", filepath.Base(os.Args[0]))
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}
	f, err := os.Open(flag.Arg(0))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	// Get first item.
	if scanner.Scan() {
		fmt.Printf("%s", scanner.Text())
	}
	for scanner.Scan() {
		fmt.Printf("%s%s", *separator, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
