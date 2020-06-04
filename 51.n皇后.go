/*
 * @lc app=leetcode.cn id=51 lang=golang
 *
 * [51] N皇后
 */

// @lc code=start
type Holder struct {
	Graph  map[int]string
	Result [][]string
}

func NewHolder(n int) *Holder {
	drawF := func(n int, i int) string {
		slice := make([]byte, n)
		x := 0
		for ; x < i; x++ {
			slice[x] = '.'
		}
		slice[x] = 'Q'
		for x += 1; x < n; x++ {
			slice[x] = '.'
		}
		return string(slice)
	}
	graph := make(map[int]string)
	for i := 0; i < n; i++ {
		graph[i] = drawF(n, i)
	}

	return &Holder{
		Graph:  graph,
		Result: make([][]string, 0, 128),
	}
}

func (h *Holder) Add(s []int) {
	d := make([]string, len(s))
	for i, v := range s {
		d[i] = h.Graph[v]
	}
	h.Result = append(h.Result, d)
}

func (h *Holder) Get() [][]string {
	return h.Result
}

func check(stack []int, x int) bool {
	y := len(stack)
	for i, v := range stack {
		if v == x || v+y == x+i || x+y == v+i {
			return false
		}
	}
	return true
}

func track(stack []int, n int, holder *Holder) {
	if h := len(stack); h >= n {
		holder.Add(stack)
		return
	}
	for i := 0; i < n; i++ {
		if check(stack, i) {
			track(append(stack, i), n, holder)
		}
	}
}

func solveNQueens(n int) [][]string {
	s := make([]int, 0, n)
	h := NewHolder(n)
	track(s, n, h)
	return h.Get()
}

// @lc code=end

