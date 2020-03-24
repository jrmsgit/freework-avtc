// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"testing"

	"github.com/jrmsgit/freework-avtc/record"
)

func BenchmarkParse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		record.Parse()
	}
}
