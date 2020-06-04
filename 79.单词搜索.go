/*
 * @lc app=leetcode.cn id=79 lang=golang
 *
 * [79] 单词搜索
 */

// @lc code=start
type P struct {
	X int
	Y int
}

func checkHistory(list []*P, x, y int) bool {
	for _, p := range list {
		if p.X == x && p.Y == y {
			return true
		}
	}
	return false
}

func track(board [][]byte, history []*P, w, h, x, y int, word string) bool {
	if len(word) == 0 {
		return true
	}
	if y < 0 || y >= h || x < 0 || x >= w {
		return false
	}
	c := word[0]
	if checkHistory(history, x, y) || board[y][x] != c {
		return false
	}
	history = append(history, &P{
		X: x,
		Y: y,
	})

	return track(board, history, w, h, x, y-1, word[1:]) ||
		track(board, history, w, h, x+1, y, word[1:]) ||
		track(board, history, w, h, x, y+1, word[1:]) ||
		track(board, history, w, h, x-1, y, word[1:])

}

func exist(board [][]byte, word string) bool {
	w := 0
	h := len(board)
	if h > 0 {
		w = len(board[0])
	}
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			history := make([]*P, 0, w*h)
			if track(board, history, w, h, x, y, word) {
				return true
			}
		}
	}
	return false
}

// @lc code=end

