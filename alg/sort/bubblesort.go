package sort

type bubbleSort struct{}

func (b *bubbleSort) Sort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			var tmp int
			if arr[j] > arr[j+1] {
				tmp = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
			}
		}
	}

	return arr
}
