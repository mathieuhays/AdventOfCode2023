package algo

import (
	"cmp"
	"testing"
)

func TestSearch(t *testing.T) {
	type args[T cmp.Ordered] struct {
		collection  []T
		target      T
		initialLeft int
	}
	type testCase[T cmp.Ordered] struct {
		name string
		args args[T]
		want int
	}
	tests := []testCase[int]{
		{"test 1", args[int]{[]int{1, 2, 3, 4, 5}, 3, 0}, 2},
		{"test 2", args[int]{[]int{1, 2, 3, 4, 5}, 9, 0}, -1},
		{"test 3", args[int]{[]int{1, 2, 3, 4, 5}, 5, 0}, 4},
		{"test 4", args[int]{[]int{1, 2, 3, 4, 5}, 1, 0}, 0},
		{"test 5", args[int]{[]int{1, 2, 3, 4}, 4, 0}, 3},
		{"test 6", args[int]{[]int{12, 15, 19, 25, 32, 45, 56, 78, 100, 132, 156, 185, 200, 203}, 156, 0}, 10},
		{"test 7", args[int]{[]int{12, 15, 19, 25, 32, 45, 56, 78, 100, 132, 156, 185, 200, 203}, 12, 0}, 0},
		{"test 8", args[int]{[]int{12, 15, 19, 25, 32, 45, 56, 78, 100, 132, 156, 185, 200, 203}, 203, 0}, 13},
		{"test 9", args[int]{[]int{12, 15, 19, 25, 32, 45, 56, 78, 100, 132, 156, 185, 200, 203}, 400, 0}, -1},
		{"test 10", args[int]{[]int{12, 15, 19, 25, 32, 45, 56, 78, 100, 132, 156, 185, 200, 203}, 38, 0}, -1},
		{"test 11", args[int]{[]int{1}, 1, 0}, 0},
		{"test 12", args[int]{[]int{1}, 2, 0}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Search(tt.args.collection, tt.args.target, tt.args.initialLeft); got != tt.want {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
