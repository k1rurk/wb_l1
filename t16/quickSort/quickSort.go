package quickSort

func QuickSort(arr []int, l, r int) {
	if l < r {

		q := partition(arr, l, r)

		QuickSort(arr, l, q)

		QuickSort(arr, q+1, r)

	}
}

func partition(arr []int, l, r int) int {
	i, j := l, r
	p := arr[(l+r)>>1]

	for i <= j {
		for p > arr[i] {
			i++
		}
		for p < arr[j] {
			j--
		}
		if i >= j {
			break
		}
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	return j
}
