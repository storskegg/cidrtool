package strings

import (
	"math"
)

func SplitByN(str string, n int) (strs []string) {
	for i := 0; i < len(str); i += n {
		nextN := int(math.Min(float64(i+n), float64(len(str))))
		strs = append(strs, str[i:nextN])
	}
	return strs
}
