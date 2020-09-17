package iteration

import (
	"fmt"
	"testing"
)

func ExampleRepeat() {
	fmt.Println(Repeat("a"))
	// Output: aaaaa
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
