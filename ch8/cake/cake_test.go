package cake_test

import (
	"testing"
	"time"

	"github.com/awmorgan/gobook/ch8/cake"
)

var defaults = cake.Shop{
	Verbose:      testing.Verbose(),
	Cakes:        20,
	BakeTime:     10 * time.Millisecond,
	NumIcers:     1,
	IceTime:      10 * time.Millisecond,
	InscribeTime: 10 * time.Millisecond,
}

func Benchmark(b *testing.B)  {
	// Baseline: one baker, one icer, one inscriber.
	// Each step takes exactly 10ms. No buffers.
	cakeshop := defaults
	cakeshop.Work(b.N)
}
