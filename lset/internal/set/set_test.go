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

package set

import (
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestStrings(t *testing.T) {
	for _, conf := range []struct {
		input []string
		want  Strings
	}{
		{
			input: nil,
			want:  Strings{},
		},
		{
			input: []string{"hello"},
			want: Strings{
				"hello": empty{},
			},
		},
		{
			input: []string{"foo", "bar"},
			want: Strings{
				"bar": empty{},
				"foo": empty{},
			},
		},
		{
			input: []string{"foo", "bar", "foo"},
			want: Strings{
				"bar": empty{},
				"foo": empty{},
			},
		},
	} {
		if diff := cmp.Diff(NewStrings(conf.input...), conf.want); diff != "" {
			t.Errorf("input %v: diff %s", conf.input, diff)
		}
	}
}

func TestStringsAdd(t *testing.T) {
	for _, conf := range []struct {
		s     Strings
		input string
		want  Strings
	}{
		{
			s:     NewStrings(),
			input: "foo",
			want: Strings{
				"foo": empty{},
			},
		},
		{
			s:     NewStrings("foo"),
			input: "bar",
			want: Strings{
				"bar": empty{},
				"foo": empty{},
			},
		},
		{
			s:     NewStrings("foo"),
			input: "foo",
			want: Strings{
				"foo": empty{},
			},
		},
	} {
		conf.s.Add(conf.input)
		if diff := cmp.Diff(conf.s, conf.want); diff != "" {
			t.Errorf("input %v: diff %s", conf.input, diff)
		}
	}
}

func TestStringsAddVarArgs(t *testing.T) {
	for _, conf := range []struct {
		s     Strings
		input []string
		want  Strings
	}{
		{
			s:     NewStrings(),
			input: []string{"bar", "foo"},
			want: Strings{
				"bar": empty{},
				"foo": empty{},
			},
		},
		{
			s:     NewStrings("foo"),
			input: []string{"bar"},
			want: Strings{
				"bar": empty{},
				"foo": empty{},
			},
		},
		{
			s:     NewStrings("bar", "foo"),
			input: []string{"bar"},
			want: Strings{
				"bar": empty{},
				"foo": empty{},
			},
		},
	} {
		conf.s.Add(conf.input...)
		if diff := cmp.Diff(conf.s, conf.want); diff != "" {
			t.Errorf("input %v: diff %s", conf.input, diff)
		}
	}
}

func TestStringsAddSet(t *testing.T) {
	for i, conf := range []struct {
		s1   Strings
		s2   Strings
		want Strings
	}{
		{
			s1: NewStrings(),
			s2: NewStrings("a", "b"),
			want: Strings{
				"a": empty{},
				"b": empty{},
			},
		},
		{
			s1: NewStrings("a", "b"),
			s2: NewStrings(),
			want: Strings{
				"a": empty{},
				"b": empty{},
			},
		},
		{
			s1: NewStrings("a"),
			s2: NewStrings("b"),
			want: Strings{
				"a": empty{},
				"b": empty{},
			},
		},
		{
			s1: NewStrings("a", "b"),
			s2: NewStrings("b", "c"),
			want: Strings{
				"a": empty{},
				"b": empty{},
				"c": empty{},
			},
		},
	} {
		conf.s1.AddSet(conf.s2)
		if diff := cmp.Diff(conf.s1, conf.want); diff != "" {
			t.Errorf("testcase %d\n%s", i, diff)
		}
	}
}

func TestStringsRemove(t *testing.T) {
	for _, conf := range []struct {
		s     Strings
		input string
		want  Strings
	}{
		{
			s:     NewStrings(),
			input: "foo",
			want:  Strings{},
		},
		{
			s:     NewStrings("foo"),
			input: "foo",
			want:  Strings{},
		},
		{
			s:     NewStrings("bar", "foo"),
			input: "bar",
			want: Strings{
				"foo": empty{},
			},
		},
	} {
		conf.s.Remove(conf.input)
		if diff := cmp.Diff(conf.s, conf.want); diff != "" {
			t.Errorf("input %v: diff %s", conf.input, diff)
		}
	}
}

func TestStringsRemoveVarArgs(t *testing.T) {
	for _, conf := range []struct {
		s     Strings
		input []string
		want  Strings
	}{
		{
			s:     NewStrings(),
			input: []string{"bar", "foo"},
			want:  Strings{},
		},
		{
			s:     NewStrings("bar", "foo", "qux"),
			input: []string{"bar", "foo"},
			want: Strings{
				"qux": empty{},
			},
		},
	} {
		conf.s.Remove(conf.input...)
		if diff := cmp.Diff(conf.s, conf.want); diff != "" {
			t.Errorf("input %v: diff %s", conf.input, diff)
		}
	}
}

func TestStringsRemoveSet(t *testing.T) {
	for i, conf := range []struct {
		s1   Strings
		s2   Strings
		want Strings
	}{
		{
			s1:   NewStrings(),
			s2:   NewStrings("a", "b"),
			want: Strings{},
		},
		{
			s1: NewStrings("a", "b"),
			s2: NewStrings(),
			want: Strings{
				"a": empty{},
				"b": empty{},
			},
		},
		{
			s1:   NewStrings("a", "b"),
			s2:   NewStrings("a", "b"),
			want: Strings{},
		},
		{
			s1: NewStrings("a"),
			s2: NewStrings("b"),
			want: Strings{
				"a": empty{},
			},
		},
		{
			s1: NewStrings("a", "b"),
			s2: NewStrings("b", "c"),
			want: Strings{
				"a": empty{},
			},
		},
	} {
		conf.s1.RemoveSet(conf.s2)
		if diff := cmp.Diff(conf.s1, conf.want); diff != "" {
			t.Errorf("testcase %d\n%s", i, diff)
		}
	}
}

func TestStringsEqual(t *testing.T) {
	for i, conf := range []struct {
		s1   Strings
		s2   Strings
		want bool
	}{
		{
			s1:   NewStrings(),
			s2:   NewStrings(),
			want: true,
		},
		{
			s1:   NewStrings("a"),
			s2:   NewStrings("a"),
			want: true,
		},
		{
			s1:   NewStrings("b", "a"),
			s2:   NewStrings("a", "b"),
			want: true,
		},
		{
			s1:   NewStrings("a", "b", "a"),
			s2:   NewStrings("a", "b"),
			want: true,
		},
		{
			s1:   NewStrings("a"),
			s2:   NewStrings(),
			want: false,
		},
		{
			s1:   NewStrings(),
			s2:   NewStrings("a", "b"),
			want: false,
		},
		{
			s1:   NewStrings("a", "b"),
			s2:   NewStrings("b", "c"),
			want: false,
		},
	} {
		if got := conf.s1.Equal(conf.s2); got != conf.want {
			t.Errorf("testcase %d: got %v, want %v", i, got, conf.want)
		}
	}
}

func TestStringsContains(t *testing.T) {
	for _, conf := range []struct {
		s     Strings
		input string
		want  bool
	}{
		{
			s:     NewStrings(),
			input: "bar",
			want:  false,
		},
		{
			s:     NewStrings("foo"),
			input: "foo",
			want:  true,
		},
		{
			s:     NewStrings("bar"),
			input: "qux",
			want:  false,
		},
		{
			s:     NewStrings("foo", "bar"),
			input: "foo",
			want:  true,
		},
	} {
		got := conf.s.Contains(conf.input)
		if got != conf.want {
			t.Errorf("input %q: got %v, want %v", conf.input, got, conf.want)
		}
	}
}

func TestStringsSlice(t *testing.T) {
	for i, want := range [][]string{
		{},
		{"a"},
		{"z", "a", "j"},
	} {
		s := NewStrings(want...)
		got := s.Slice()
		// Sort the returned slice and input before doing the comparison.
		sort.Strings(got)
		sort.Strings(want)
		if diff := cmp.Diff(got, want); diff != "" {
			t.Errorf("testcase %d\n%s", i, diff)
		}
	}
}

func TestStringsSortedSlice(t *testing.T) {
	for i, want := range [][]string{
		{},
		{"a"},
		{"z", "a", "j"},
	} {
		s := NewStrings(want...)
		got := s.SortedSlice()
		// Sort the input before comparison as the returned slice should already be sorted.
		sort.Strings(want)
		if diff := cmp.Diff(got, want); diff != "" {
			t.Errorf("testcase %d\n%s", i, diff)
		}
	}
}

func TestStringsIntersect(t *testing.T) {
	for i, conf := range []struct {
		s1   Strings
		s2   Strings
		want Strings
	}{
		{
			s1:   NewStrings(),
			s2:   NewStrings(),
			want: NewStrings(),
		},
		{
			s1:   NewStrings(),
			s2:   NewStrings("a"),
			want: NewStrings(),
		},
		{
			s1:   NewStrings("a", "b"),
			s2:   NewStrings(),
			want: NewStrings(),
		},
		{
			s1:   NewStrings("a", "b"),
			s2:   NewStrings("b", "c"),
			want: NewStrings("b"),
		},
		{
			s1:   NewStrings("a", "b", "c"),
			s2:   NewStrings("a", "b", "c", "d"),
			want: NewStrings("a", "b", "c"),
		},
	} {
		if diff := cmp.Diff(conf.s1.Intersect(conf.s2), conf.want); diff != "" {
			t.Errorf("testcase %d\n%s", i, diff)
		}
	}
}

func TestStringsDiff(t *testing.T) {
	for i, conf := range []struct {
		s1          Strings
		s2          Strings
		wantAdded   Strings
		wantRemoved Strings
	}{
		{
			s1:          NewStrings(),
			s2:          NewStrings(),
			wantAdded:   NewStrings(),
			wantRemoved: NewStrings(),
		},
		{
			s1:          NewStrings("a"),
			s2:          NewStrings("a"),
			wantAdded:   NewStrings(),
			wantRemoved: NewStrings(),
		},
		{
			s1:          NewStrings(),
			s2:          NewStrings("a"),
			wantAdded:   NewStrings("a"),
			wantRemoved: NewStrings(),
		},
		{
			s1:          NewStrings("a", "b"),
			s2:          NewStrings("b", "c"),
			wantAdded:   NewStrings("c"),
			wantRemoved: NewStrings("a"),
		},
		{
			s1:          NewStrings("a", "b", "c", "d"),
			s2:          NewStrings("a", "b", "c"),
			wantAdded:   NewStrings(),
			wantRemoved: NewStrings("d"),
		},
	} {
		added, removed := conf.s1.Diff(conf.s2)
		if diff := cmp.Diff(added, conf.wantAdded); diff != "" {
			t.Errorf("testcase %d\nadded: %s", i, diff)
		}
		if diff := cmp.Diff(removed, conf.wantRemoved); diff != "" {
			t.Errorf("testcase %d\nremoved: %s", i, diff)
		}
	}
}
