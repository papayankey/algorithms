package algorithms

func MergeSort(s []int, size int) {
	if size < 2 {
		return
	}

	mid := size >> 1
	left := make([]int, mid)
	right := make([]int, len(s)-mid)

	copy(left, s[:mid])
	copy(right, s[mid:])

	MergeSort(left, mid)
	MergeSort(right, len(s)-mid)

	merge(s, left, right, mid, len(s)-mid)
}

func merge(s []int, left []int, right []int, mid, r int) {
	i, j, k := 0, 0, 0
	for i < mid && j < r {
		if left[i] < right[j] {
			s[k] = left[i]
			i++
		} else {
			s[k] = right[j]
			j++
		}
		k++
	}

	for i < mid {
		s[k] = left[i]
		i++
		k++
	}

	for j < r {
		s[k] = right[j]
		j++
		k++
	}
}
