package integers

import (
	"fmt"
	"testing"
	"testing/quick"
)

func TestAdder(t *testing.T) {
	commutative := func(x, y int) bool {
		if Add(x, y) != Add(y, x) {
			return false
		}
		return true
	}
	if err := quick.Check(commutative, nil); err != nil {
		t.Error(err)
	}

	sum := func(x, y int) bool {
		if Add(x, y) != x+y {
			return false
		}
		return true
	}
	if err := quick.Check(sum, nil); err != nil {
		t.Error(err)
	}
}

func EampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
