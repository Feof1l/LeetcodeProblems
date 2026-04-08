package main

import (
	"log"
	"math"
)

func main() {
	log.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		tmp := nums2
		nums2 = nums1
		nums1 = tmp
	}

	left := 0
	right := len(nums1)
	for left <= right {
		partitionA := (left + right) / 2
		partitionB := (len(nums1)+len(nums2)+1)/2 - partitionA

		maxALeft := math.MinInt64
		if partitionA != 0 {
			maxALeft = nums1[partitionA-1]
		}
		maxBLeft := math.MinInt64
		if partitionB != 0 {
			maxBLeft = nums2[partitionB-1]
		}
		minARight := math.MaxInt64
		if partitionA != len(nums1) {
			minARight = nums1[partitionA]
		}
		minBRight := math.MaxInt64
		if partitionB != len(nums2) {
			minBRight = nums2[partitionB]
		}

		if maxALeft <= minBRight && maxBLeft <= minARight {
			if (len(nums1)+len(nums2))%2 == 0 {
				maxLeft := math.Max(float64(maxALeft), float64(maxBLeft))
				minRight := math.Min(float64(minARight), float64(minBRight))
				return (maxLeft + minRight) / 2.0
			} else {
				return math.Max(float64(maxALeft), float64(maxBLeft))
			}
		} else if maxBLeft > minARight {
			left = partitionA + 1
		} else {
			right = partitionA - 1
		}
	}

	return 0.0
}
