// Copyright (c) The Thanos Authors.
// Licensed under the Apache License 2.0.

package strutil

import (
	"math/rand"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var randomChar = "abcderfgijklmnopq"

func BenchmarkMergeSlices(b *testing.B) {
	randomStrings := []string{}
	input := make([][]string, 500)

	for i := 0; i < 5000000; i++ {
		sb := strings.Builder{}
		for j := 0; j < 15; j++ {
			sb.WriteByte(randomChar[rand.Int()%len(randomChar)])
		}
		randomStrings = append(randomStrings, sb.String())
	}

	mapContains := map[string]struct{}{}

	for i := 0; i < 800000; i++ {
		s := randomStrings[rand.Int()%len(randomStrings)]
		for j := 0; j < 3; j++ {
			index := rand.Int() % len(input)
			if _, ok := mapContains[s]; !ok {
				input[index] = append(input[index], s)
				mapContains[s] = struct{}{}
			}
		}
	}

	for i := 0; i < len(input); i++ {
		sort.Strings(input[i])
	}
	require.Equal(b, dedupUsingMap(input...), MergeSlices(input...))

	b.Run("Run", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			r := MergeSlices(input...)
			require.NotEmpty(b, r)
		}
	})
}

func dedupUsingMap(slices ...[]string) []string {
	rSet := map[string]struct{}{}
	for _, s := range slices {
		for _, v := range s {
			rSet[v] = struct{}{}
		}
	}
	values := make([]string, 0, len(rSet))
	for v := range rSet {
		values = append(values, v)
	}

	sort.Strings(values)
	return values
}

