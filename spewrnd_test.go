package main

import (
	"io/ioutil"
	"testing"
)

func BenchmarkSpewRandom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		spewRandom(ioutil.Discard)
	}
}
