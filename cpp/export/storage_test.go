package main

import "testing"

func TestUUID(t *testing.T) {
	seen := make(map[uint64]struct{})

	for i := 0; i < 10000000; i++ {
		u := uuid()
		if u == 0 {
			t.Fatal("Got a 0 value uuid")
		}
		if _, exists := seen[u]; exists {
			t.Fatalf("Dup uuid: %v", u)
		}
		seen[u] = struct{}{}
	}
}

func BenchmarkUUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		uuid()
	}
}
