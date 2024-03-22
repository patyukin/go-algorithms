package intersection_of_sorted_arrays

func intersection(nums1 []int, nums2 []int) []int {
	var result []int
	var maxLength = max(len(nums1), len(nums2))
	var p1, p2 = 0, 0

	for p1 != len(nums1) && p2 != len(nums2) && p1 < maxLength && p2 < maxLength {
		if nums1[p1] < nums2[p2] {
			p1++
		} else if nums1[p1] > nums2[p2] {
			p2++
		} else {
			result = append(result, nums1[p1])
			p1++
			p2++
		}
	}

	return result
}
