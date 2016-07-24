package sort

func BubbleSort(srcdata []int) {
	count := len(srcdata)
	for i := 0; i < count; i++ {
		for j := i + 1; j < count; j++ {
			if srcdata[j] < srcdata[i] {
				tmp := srcdata[j]
				srcdata[j] = srcdata[i]
				srcdata[i] = tmp
			}
		}
	}
}
