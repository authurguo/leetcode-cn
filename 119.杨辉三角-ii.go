/*
 * @lc app=leetcode.cn id=119 lang=golang
 *
 * [119] 杨辉三角 II
 */

// @lc code=start
// func getColumn(rowIndex int, colIndex int) int {
// 	if rowIndex == 0 {
// 		return 1
// 	} else if rowIndex == 1 {
// 		return 1
// 	} else if rowIndex == 2 {
// 		switch colIndex {
// 		case 1:
// 			return 2
// 		default:
// 			return 1
// 		}
// 	} else {
// 		if colIndex == 0 || colIndex >= rowIndex {
// 			return 1
// 		}
// 		k := getColumn(rowIndex-1, colIndex) + getColumn(rowIndex-1, colIndex-1)
// 		return k
// 	}
// }

// func getRow(rowIndex int) []int {
// 	r := make([]int, 0, rowIndex)
// 	m := (rowIndex+1)/2 + (rowIndex+1)%2
// 	for i := 0; i < m; i++ {
// 		r = append(r, getColumn(rowIndex, i))
// 	}
// 	for i := m; i <= rowIndex; i++ {
// 		r = append(r, r[rowIndex-i])
// 	}
// 	return r
// }

func getRow(rowIndex int) []int {
	r := make([]int, 0, rowIndex)
	for i := 0; i <= rowIndex; i++ {
		r = append(r, 1)
		for j := i - 1; j > 0; j-- {
			r[j] += r[j-1]
		}
	}
	return r
}

// @lc code=end

