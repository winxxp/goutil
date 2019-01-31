package convey

import (
	"fmt"
	"github.com/winxxp/goutil/fileutil"
	"math"
)

func ShouldSliceAlmostEqual(actual interface{}, expected ...interface{}) string {
	var left, right []float32
	var ok bool

	if left, ok = actual.([]float32); !ok {
		return "actual not float slice"
	}
	if right, ok = expected[0].([]float32); !ok {
		return "expected not float slice"
	}

	if len(left) != len(right) {
		return fmt.Sprintf("len(actual) != len(expeced) %d != %d", len(left), len(right))
	}

	deltaFloat := 1e-3
	for i, v := range left {
		if math.Abs(float64(right[i])-float64(v)) <= deltaFloat {
			return ""
		} else {
			return fmt.Sprintf("Expected '%v' to almost equal '%v' (but it didn't)!", actual, expected)
		}
	}

	return ""
}

func ShouldFileEqual(actual interface{}, expected ...interface{}) string {
	f1, ok := actual.(string)
	if !ok {
		return "actual not filename"
	}
	f2, ok := expected[0].(string)
	if !ok {
		return "expected not filename"
	}

	if !fileutil.Equal(f1, f2) {
		return fmt.Sprintf("%s != %s", f1, f2)
	}

	return ""
}
