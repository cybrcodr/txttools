// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// The csvflat command line tool removes new lines in each column.
package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		usage()
		os.Exit(1)
	}

	filename := os.Args[1]
	f := os.Stdin
	if filename != "-" {
		var err error
		f, err = os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading file %s: %v", filename, err)
			os.Exit(1)
		}
		defer f.Close()
	}

	if err := process(f); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s <file>\n", filepath.Base(os.Args[0]))
	fmt.Fprintln(os.Stderr)
	fmt.Fprintln(os.Stderr, "Removes new line character in CSV columns")
	fmt.Fprintln(os.Stderr, "If file is '-', it reads from stdin.")
	fmt.Fprintln(os.Stderr)
}

func process(f *os.File) error {
	w := csv.NewWriter(os.Stdout)
	r := csv.NewReader(f)
	r.ReuseRecord = true
	r.FieldsPerRecord = -1

	for {
		cols, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		recs := make([]string, len(cols))
		for i, col := range cols {
			recs[i] = removeNewLines(col)
		}
		if err := w.Write(recs); err != nil {
			return err
		}
		w.Flush()
	}

	return w.Error()
}

func removeNewLines(s string) string {
	list := strings.Split(s, "\n")
	return strings.Join(list, " ")
}
