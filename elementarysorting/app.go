package elementarysorting

// 挿入ソート
func insertionSort(a []int) {
	for c := 1; c < len(a); c++ {
		evalTarget := a[c]
		i := c - 1
		for i >= 0 && a[i] > evalTarget {
			a[i+1] = a[i]
			i--
		}
		a[i+1] = evalTarget
	}
}
