package e303

type NumArray struct {
	nums       []int
	prefixNums []int
}

func Constructor(nums []int) NumArray {
	prefixNums := make([]int, 0)
	prefixNums = append(prefixNums, 0)

	var sum int
	for num := range nums {
		sum += num
		prefixNums = append(prefixNums, sum)
	}

	return NumArray{
		nums:       nums,
		prefixNums: prefixNums,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.prefixNums[right+1] - this.prefixNums[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
