package locks

import (
	"testing"
	"testing/quick"
	"math/rand"
	"reflect"
)

func TestBeat(t *testing.T) {
	tests := []struct{doors, iterations, unlocked int } {
		{3,1,2},
		{100,100,50},
	}
	for _, test := range  tests {
		unlocked := smart(test.doors, test.iterations) 
		if unlocked != test.unlocked {
			t.Errorf("smart(%d,%d) = %d; want %d", test.doors, test.iterations, unlocked, test.unlocked)
		}
	}
}

func TestQuick(t *testing.T) {
	var config quick.Config
	config.Values = func(v []reflect.Value, r *rand.Rand) {
		v[0] = reflect.ValueOf(r.Intn(2000))
		v[1] = reflect.ValueOf(r.Intn(300))
	}
	if err := quick.CheckEqual(smart, dumb, &config); err != nil {
		t.Error(err)
	}
}

func dumb(doors, iterations int) int {
	locked := make([]bool, doors)
	for i := 0; i < iterations-1; i++ {
		for j := 1; j < doors; j += 2 {
			locked[j] = true
		}
		for j := 2; j < doors; j += 3 {
			locked[j] = !locked[j]
		}
	}
	if iterations > 0 {
		locked[doors-1] = !locked[doors-1]
	}
	count := 0
	for _, d := range locked {
		if !d {
			count++
		}
	}
	return count
}

