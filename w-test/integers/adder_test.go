package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(4, 2)
	expected := 6

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func Add(i int, i2 int) interface{} {
	return i + i2
}

// Will appear in `godoc` documentation
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
