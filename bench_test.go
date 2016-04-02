package paralyze

import "testing"

var (
	fasterFn = func() (interface{}, error) { return 55, nil }
)

func BenchmarkWithChannelAlloc(test *testing.B) {
	for i := 0; i < test.N; i++ {
		ParalyzeWithTimeout(0, []Paralyzable{
			fasterFn,
			fasterFn,
			fasterFn,
			fasterFn,
			fasterFn,
			fasterFn,
		}...)
	}
}

func BenchmarkWithoutChannelAlloc(test *testing.B) {
	for i := 0; i < test.N; i++ {
		Paralyze([]Paralyzable{
			fasterFn,
			fasterFn,
			fasterFn,
			fasterFn,
			fasterFn,
			fasterFn,
		}...)
	}
}