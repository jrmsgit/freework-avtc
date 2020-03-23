// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"testing"

	"github.com/jrmsgit/freework-avtc/record"
)

func BenchmarkGen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		record.Gen()
	}
}
