package algorithm

func merge(data []int, start int, mid int, end int) {
	aux := make([]int, len(data))
	copy(aux, data)

	i, j := start, mid+1
	for idx := start; idx <= end; idx++ {
		if i > mid {
			data[idx] = aux[j]
			j++
		} else if j > end {
			data[idx] = aux[i]
			i++
		} else if aux[i] > aux[j] {
			data[idx] = aux[j]
			j++
		} else {
			data[idx] = aux[i]
			i++
		}
	}
}

func sort(data []int, start int, end int) {
	if start >= end {
		return
	}
	mid := start + (end-start)/2
	sort(data, start, mid)
	sort(data, mid+1, end)
	merge(data, start, mid, end)
}

func MergeSort(data []int) {
	sort(data, 0, len(data)-1)
}
