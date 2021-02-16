package main

import "fmt"

func main() {
	res := findMedianSortedArrays([]int{1, 2}, []int{3, 4})
	fmt.Println(res)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	mergedArr := make([]int, 0)
	counter := 0
	counter2 := 0

	for counter < len(nums1) && counter2 < len(nums2) {
		if nums1[counter] < nums2[counter2] {
			mergedArr = append(mergedArr, nums1[counter])
			counter++
		} else {
			mergedArr = append(mergedArr, nums2[counter2])
			counter2++
		}
	}

	if counter >= len(nums1) {
		mergedArr = append(mergedArr, nums2[counter2:]...)
	} else if counter2 >= len(nums2) {
		mergedArr = append(mergedArr, nums1[counter:]...)
	}

	medianIndexes := []int{0, -1}
	if len(mergedArr)%2 == 0 {
		medianIndexes[0] = len(mergedArr)/2 - 1
		medianIndexes[1] = medianIndexes[0] + 1

		return float64(mergedArr[medianIndexes[0]]+mergedArr[medianIndexes[1]]) / 2
	} else {
		medianIndexes[0] = len(mergedArr) / 2
		return float64(mergedArr[medianIndexes[0]])
	}
}
