package main

func countSplitInv(a, b []int) (int, []int) {
	n := len(a) + len(b)
	c := make([]int, 0, n)

	i, j, count := 0, 0, 0
	for k := 0; k < n; k++ {
		if j >= len(b) {
			c = append(c, a[i])
			i++
		} else if i >= len(a) {
			c = append(c, b[j])
			j++
		} else if a[i] <= b[j] {
			c = append(c, a[i])
			i++
		} else {
			c = append(c, b[j])
			count += len(a) - i
			j++
		}
	}

	return count, c
}

// CountInv ...
func CountInv(a []int) (int, []int) {
	if len(a) <= 1 {
		return 0, a
	}

	n := len(a) / 2
	leftCount, c := CountInv(a[:n])
	rightCount, d := CountInv(a[n:])
	splitCount, b := countSplitInv(c, d)

	return leftCount + rightCount + splitCount, b
}
