package algorithms

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	tests := []struct {
		desc   string
		input  []int
		output []int
	}{
		{"test00", []int{10, 8, 5, 2, 6, 2}, []int{2, 2, 5, 6, 8, 10}},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			MergeSort(test.input, len(test.input))
			want := test.output

			if !reflect.DeepEqual(test.input, want) {
				t.Errorf("MergeSort(%v, %d) = %v, want %v", test.input, len(test.input), test.input, want)
			}
		})
	}
}
