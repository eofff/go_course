package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

type Record struct {
	A, B, C, D, E, F string
}

var (
	rec = Record{"a", "b", "c", "d", "e", "f"}
)

func BenchmarkConcatFmt(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%v_%v_%v_%v_%v_%v_%v",
			rec.A,
			rec.B,
			rec.C,
			rec.D,
			rec.E,
			rec.F)
	}
}

func BenchmarkConcatBytes(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var buff bytes.Buffer
		buff.WriteString(rec.A)
		buff.WriteString("_")
		buff.WriteString(rec.B)
		buff.WriteString("_")
		buff.WriteString(rec.C)
		buff.WriteString("_")
		buff.WriteString(rec.D)
		buff.WriteString("_")
		buff.WriteString(rec.E)
		buff.WriteString("_")
		buff.WriteString(rec.F)
		buff.String()
	}
}

func BenchmarkConcatStrings(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		strings.Join([]string{rec.A,
			rec.B,
			rec.C,
			rec.D,
			rec.E,
			rec.F},
			"_")
	}
}

// BenchmarkConcatFmt-4       	 1000000	      2026 ns/op	     128 B/op	       7 allocs/op
// BenchmarkConcatBytes-4     	 3000000	       497 ns/op	     112 B/op	       1 allocs/op
// BenchmarkConcatStrings-4   	 5000000	       329 ns/op	      32 B/op	       2 allocs/op
