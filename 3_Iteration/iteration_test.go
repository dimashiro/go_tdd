package iteration

import (
	"fmt"
	"testing"
)

func TestIteration(t *testing.T) {
	t.Run("get right string", func(t *testing.T) {
		result := Repeat("a")
		expected := "aaaaa"
		if result != expected {
			t.Errorf("expected %q got %q", expected, result)
		}
	})
	t.Run("count symblols", func(t *testing.T) {
		result := len(Repeat("a"))
		if result != repeatCount {
			t.Errorf("expected %q got %q", repeatCount, result)
		}
	})
}
func ExampleIteration() {
	fmt.Println(Repeat("b"))
	//Output: bbbbb

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
