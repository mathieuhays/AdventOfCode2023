package algo

import (
	"cmp"
	"math"
)

func Search[T cmp.Ordered](collection []T, target T, initialLeft int) int {
	left := initialLeft
	right := len(collection) - 1

	for {
		cursor := left + int(math.Ceil(float64(((right - left) / 2))))
		value := collection[cursor]

		if value == target {
			return cursor
		}

		if value > target {
			right = cursor - 1
		} else {
			left = cursor + 1
		}

		if right < left {
			break
		}
	}

	return -1
}
