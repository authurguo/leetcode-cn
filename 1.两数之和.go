/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	for i, iv := range nums {
		for j, jv := range nums[i:] {
			if 0 != j && jv == target-iv {
				return []int{i, i + j}
			}
		}
	}
	return nil
}

// @lc code=end

