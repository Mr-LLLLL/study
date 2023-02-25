package probles

type NumArray struct {
	sum []int
}

func NewNumArray(nums []int) NumArray {
	obj := new(NumArray)
	obj.sum = make([]int, len(nums)+1)

	for i := 0; i < len(nums); i++ {
		obj.sum[i+1] += obj.sum[i] + nums[i]
	}

	return *obj
}

func (this *NumArray) SumRange(left, right int) int {
	return this.sum[right+1] - this.sum[left]
}
