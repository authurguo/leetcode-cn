/*
 * @lc app=leetcode.cn id=13 lang=golang
 *
 * [13] 罗马数字转整数
 */

// @lc code=start
var (
	valueM = map[string]int{
		"I":  1,
		"IV": 4,
		"V":  5,
		"IX": 9,
		"X":  10,
		"XL": 40,
		"L":  50,
		"XC": 90,
		"C":  100,
		"CD": 400,
		"D":  500,
		"CM": 900,
		"M":  1000,
	}
)

func romanToInt(s string) int {
	l := len(s)
	v := 0
	for i := 0; i < l; i++ {
		vi := 0
		ok := false
		if vi, ok = valueM[string(s[i])]; !ok {
			return -1
		}

		if j := i + 1; j < l {
			vj := 0
			if vj, ok = valueM[string(s[j])]; !ok {
				return -2
			}
			if vi < vj {
				vk := 0
				if vk, ok = valueM[string(s[i])+string(s[j])]; !ok {
					return -3
				}
				v += vk
				i++
				continue
			}
		}
		v += vi
	}
	return v
}

// @lc code=end

