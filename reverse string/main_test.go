package main

import "testing"

func BenchmarkReverseString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		str := "Hello World"
		reversed := ""

		for j := range str {
			reversed += string(str[(len(str)-j)-1])
		}
	}
}
