package goxoshiro256starstar_test

import (
	"github.com/epes/goxoshiro256starstar"
	"math"
	"testing"
)

func TestNext(t *testing.T) {
	r := goxoshiro256starstar.New(1234567)

	tests := []uint64{
		3504822795582309479,
		1819558768956484042,
		1250851346055027673,
		16940231675099994102,
		11585879347611423030,
		8134400763355999650,
		16522854393704305783,
		6681395768013188110,
		6428666302753294433,
		3497713684171130211,
	}

	for _, expected := range tests {
		got := r.Next()

		if expected != got {
			t.Fatalf("expected: %v, got: %v", expected, got)
		}
	}
}

func TestSpread(t *testing.T) {
	r := goxoshiro256starstar.New(987654321)
	results := make([]uint64, 5)

	for i := 0; i < 100000; i++ {
		asRatio := float64(r.Next()) / (1 << 64)
		floor := int(math.Floor(asRatio * 5))
		results[floor]++
	}

	tests := []uint64{
		20072,
		20195,
		19854,
		19932,
		19947,
	}

	for idx, expected := range tests {
		got := results[idx]

		if expected != got {
			t.Fatalf("expected: %v, got: %v", expected, got)
		}
	}
}
