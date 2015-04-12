package goherokuname

import (
	"fmt"
	"testing"
)

func ExampleHaikunate() {
	seed = 99
	fmt.Println(Haikunate())
	// Output: dawn-pond-0223
}

func ExampleHaikunateHex() {
	seed = 99
	fmt.Println(HaikunateHex())
	// Output: dawn-pond-2a6b
}

func ExampleHaikunateCustom() {
	seed = 99
	fmt.Println(HaikunateCustom("+", 5, "abcd"))
	// Output: dawn+pond+cccdb
}

func BenchmarkHaikunateDeterministic(b *testing.B) {
	seed = 99
	for i := 0; i < b.N; i++ {
		Haikunate()
	}
}

func BenchmarkHaikunateHexDeterministic(b *testing.B) {
	seed = 99
	for i := 0; i < b.N; i++ {
		HaikunateHex()
	}
}

func BenchmarkHaikunateCustomDeterministic(b *testing.B) {
	seed = 99
	for i := 0; i < b.N; i++ {
		HaikunateCustom("+", 5, "abcd")
	}
}

func BenchmarkHaikunateRandom(b *testing.B) {
	seed = 0
	for i := 0; i < b.N; i++ {
		Haikunate()
	}
}

func BenchmarkHaikunateHexRandom(b *testing.B) {
	seed = 0
	for i := 0; i < b.N; i++ {
		HaikunateHex()
	}
}

func BenchmarkHaikunateCustomRandom(b *testing.B) {
	seed = 0
	for i := 0; i < b.N; i++ {
		HaikunateCustom("+", 5, "abcd")
	}
}
