/*
 * @lc app=leetcode.cn id=52 lang=golang
 *
 * [52] N皇后 II
 */

// @lc code=start
type Holder struct {
	List [][]int
}

func deepCopy(s []int) []int {
	d := make([]int, len(s))
	for i, v := range s {
		d[i] = v
	}
	return d
}

func check(s []int, x int) bool {
	y := len(s)
	for i, v := range s {
		if v == x || v+i == x+y || v+y == x+i {
			return false
		}
	}
	return true
}

func track(s []int, n int, h *Holder) {
	if len(s) >= n {
		h.List = append(h.List, deepCopy(s))
		return
	}
	for i := 0; i < n; i++ {
		if check(s, i) {
			track(append(s, i), n, h)
		}
	}
}

func totalNQueens(n int) int {
	s := make([]int, 0, n)
	h := &Holder{
		List: make([][]int, 0, 128),
	}
	track(s, n, h)
	return len(h.List)
}

// @lc code=end

