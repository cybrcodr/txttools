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

// Package set contains set data structure for strings.
package set

import "sort"

// Strings is a set of strings.
type Strings map[string]empty

type empty struct{}

// NewStrings constructs a Strings object.
func NewStrings(strs ...string) Strings {
	s := make(Strings, len(strs))
	s.Add(strs...)
	return s
}

// Add adds given string(s).
func (s Strings) Add(strs ...string) {
	for _, value := range strs {
		s[value] = empty{}
	}
}

// AddSet adds values from another Strings object.
func (s Strings) AddSet(strs Strings) {
	for str := range strs {
		s.Add(str)
	}
}

// Remove removes string from set.
func (s Strings) Remove(strs ...string) {
	for _, value := range strs {
		delete(s, value)
	}
}

// RemoveSet removes values from another Strings object.
func (s Strings) RemoveSet(strs Strings) {
	for str := range strs {
		s.Remove(str)
	}
}

// Equal returns true if s and o contain the same elements.
func (s Strings) Equal(o Strings) bool {
	if len(s) != len(o) {
		return false
	}
	for str := range s {
		if !o.Contains(str) {
			return false
		}
	}
	return true
}

// Contains returns true if string is in set, else false.
func (s Strings) Contains(str string) bool {
	_, ok := s[str]
	return ok
}

// Slice returns strings within set.
func (s Strings) Slice() []string {
	ret := make([]string, 0, len(s))
	for str := range s {
		ret = append(ret, str)
	}
	return ret
}

// SortedSlice returns strings in sorted order.
func (s Strings) SortedSlice() []string {
	strs := s.Slice()
	sort.Strings(strs)
	return strs
}

// Intersect returns the intersection of s and o.
func (s Strings) Intersect(o Strings) Strings {
	ret := Strings{}
	for str := range s {
		if o.Contains(str) {
			ret.Add(str)
		}
	}
	return ret
}

// Diff returns the difference between the current and other sets.
// The first returned value are strings in the other set not in the current.
// The second returned value are strings in the current set not in the other.
func (s Strings) Diff(other Strings) (Strings, Strings) {
	return other.Minus(s), s.Minus(other)
}

// Minus returns strings in other set that are not in current set.
func (s Strings) Minus(other Strings) Strings {
	leftover := Strings{}
	for i := range s {
		if _, ok := other[i]; !ok {
			leftover.Add(i)
		}
	}
	return leftover
}
