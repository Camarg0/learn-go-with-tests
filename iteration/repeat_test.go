package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

// Benchmarking: for measuring code perfomance (nanoseconds per operation)
func BenchmarkRepeat(b *testing.B) {
	// Loop returns true as long the benchmark is running
	for b.Loop() {
		// Code to measure
		Repeat("a")
	}
}
