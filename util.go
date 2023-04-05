package algorithms

// GetMax returns the maximum element in the slice
func GetMax(s []int, n int) int {
	max := s[0]
	for i := 1; i < n; i++ {
		if s[i] > max {
			max = s[i]
		}
	}
	return max
}

// Swap exchanges values of given positions
func Swap(src []int, i, j int) {
	src[i], src[j] = src[j], src[i]
}
