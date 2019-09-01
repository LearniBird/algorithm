package leetcode


func TwoSum(nums []int, target int) []int {
	numLength := len(nums)
	i,j := 0,0
	for ;i < numLength; i++ {
		for j = i+1;j < numLength; j++ {
			if (nums[i]+nums[j]) == target {
				goto End
			}
		}
	}
	End:
	return []int{i,j}
}