package array

type NumArray struct {
	data []int
}

func Constructor(nums []int) NumArray {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	return NumArray{data: nums}
}

func (na *NumArray) SumRange(left, right int) int {
	if left == 0 {
		return na.data[right]
	}
	return na.data[right] - na.data[left-1]
}
