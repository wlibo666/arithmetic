package recursion

import (
	"testing"
)

func TestMulti(t *testing.T) {
	number := 6
	result := Multi(number)
	t.Logf("Multi(%d) = %d", number, result)
}
