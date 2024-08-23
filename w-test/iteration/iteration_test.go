package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}

const repeatCount = 5

func Repeat(s string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated = repeated + s
	}
	return repeated
}
