// Copyright (c) The Thanos Authors.
// Licensed under the Apache License 2.0.

package strutil

import (
	"sort"

	"github.com/bboreham/go-loser"
)

func NewStringListIter(s []string) *StringListIter {
	return &StringListIter{l: s, cur: -1}
}

type StringListIter struct {
	l   []string
	cur int
}

func (s *StringListIter) Next() bool {
	if s.cur+1 >= len(s.l) {
		return false
	}
	s.cur++
	return true
}
func (s *StringListIter) At() string { return s.l[s.cur] }

// MergeSlices merges a set of sorted string slices into a single ones
// while removing all duplicates.
func MergeSlices(a ...[]string) []string {
	its := make([]*StringListIter, 0, len(a))
	maxLength := 0

	for _, s := range a {
		if len(s) > maxLength {
			maxLength = len(s)
		}
		its = append(its, NewStringListIter(s))
	}
	r := make([]string, 0, maxLength)
	lt := loser.New(its, string([]byte{0xff}))
	if maxLength == 0 {
		return r
	}
	lt.Next()
	current := lt.At()
	r = append(r, current)
	for lt.Next() {
		if lt.At() != current {
			current = lt.At()
			r = append(r, current)
		}
	}
	return r
}

// MergeUnsortedSlices behaves like StringSlices but input slices are validated
// for sortedness and are sorted if they are not ordered yet.
func MergeUnsortedSlices(a ...[]string) []string {
	for _, s := range a {
		if !sort.StringsAreSorted(s) {
			sort.Strings(s)
		}
	}
	return MergeSlices(a...)
}

