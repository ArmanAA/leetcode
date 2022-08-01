package main

import (
	"fmt"
	"sort"
)

func main() {
	// fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	// fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
	board := [][]string{{".", ".", ".", ".", "5", ".", ".", "1", "."}, {".", "4", ".", "3", ".", ".", ".", ".", "."}, {".", ".", ".", ".", ".", "3", ".", ".", "1"}, {"8", ".", ".", ".", ".", ".", ".", "2", "."}, {".", ".", "2", ".", "7", ".", ".", ".", "."}, {".", "1", "5", ".", ".", ".", ".", ".", "."}, {".", ".", ".", ".", ".", "2", ".", ".", "."}, {".", "2", ".", "9", ".", ".", ".", ".", "."}, {".", ".", "4", ".", ".", ".", ".", ".", "."}}
	fmt.Println(isValidSudoku(board))
}

// 1- Contains Duplicate: https://leetcode.com/problems/contains-duplicate/submissions/

func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)
	for _, v := range nums {
		if _, ok := m[v]; ok {
			return true
		}
		m[v] = true
	}
	return false
}

// 2- Valid Anagram: https://leetcode.com/problems/valid-anagram/

// can be solved with comparing 2 maps as well
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	chars := [26]int{}
	for i := 0; i < len(s); i++ {
		chars[int(s[i]-'a')]++
		chars[int(t[i]-'a')]--
	}
	for i := 0; i < 26; i++ {
		if chars[i] != 0 {
			return false
		}
	}

	return true
}

// 3- Two Sum: https://leetcode.com/problems/two-sum/
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, n := range nums {
		if v, ok := m[n]; ok {
			return []int{i, v}
		}
		m[target-n] = i
	}
	return []int{}
}

// 4- Group Anagrams: https://leetcode.com/problems/group-anagrams/
func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]int][]string)
	for _, word := range strs {
		var arr [26]int
		for _, c := range word {
			arr[c-97]++
		}
		m[arr] = append(m[arr], word)
	}
	var result [][]string
	for _, v := range m {
		result = append(result, v)
	}
	return result
}

// 5- Product of Array Except Self: https://leetcode.com/problems/product-of-array-except-self/

func productExceptSelf(nums []int) []int {
	// currIndex := 0
	pre := make(map[int]int)
	post := make(map[int]int)
	var result []int
	j := len(nums) - 1
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			pre[i] = 1
		} else {
			pre[i] = nums[i-1] * pre[i-1]
		}
		if j == len(nums)-1 {
			post[j] = 1
		} else {
			post[j] = post[j+1] * nums[j+1]
		}
		j--
	}
	for i := 0; i < len(nums); i++ {
		result = append(result, pre[i]*post[i])
	}
	return result
}

// 6- Top K Frequent Elements: https://leetcode.com/problems/top-k-frequent-elements/
func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)

	for _, v := range nums {
		if _, ok := m[v]; !ok {
			m[v] = 1
		} else {
			m[v]++
		}
	}
	var temp PairList
	for k, v := range m {
		p := Pair{Key: k, Value: v}
		temp = append(temp, p)
	}
	sort.Sort(temp)

	var result []int
	for i := 0; i < k; i++ {
		result = append(result, temp[i].Key)
	}
	return result
}

type Pair struct {
	Key   int
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value > p[j].Value }

// 7- Valid Sudoku: https://leetcode.com/problems/valid-sudoku/
func isValidSudoku(board [][]string) bool {
	// check rows
	for _, row := range board {
		m := make(map[string]bool)
		for _, v := range row {
			if string(v) == "." {
				continue
			}
			if _, ok := m[v]; !ok {
				m[v] = true
			} else {
				return false
			}
		}
	}
	// check cols
	for i := 0; i <= 8; i++ {
		m := make(map[string]bool)
		for j := 0; j <= 8; j++ {

			item := board[j][i]
			if string(item) == "." {
				continue
			}
			if _, ok := m[item]; !ok {
				m[item] = true
			} else {
				return false
			}
		}
	}
	// check 3*3 squares
	m := make(map[[2]int]string)
	// for n := 0; n <= 2; n++ {
	for i := 0; i <= 8; i++ {
		for j := 0; j <= 8; j++ {

			item := board[i][j]

			if string(item) == "." {
				continue
			}
			arr := [2]int{i / 3, j / 3}
			fmt.Println("case: ", arr, "i: ", i, "j: ", j, "item: ", item)

			if _, ok := m[arr]; !ok {
				m[arr] = item
			} else {
				return false
			}
			// fmt.Println("case: ", arr)
			// switch arr {
			// case [2]int{0, 0}:
			// 	fmt.Println("case: ", arr, "i: ", i, "j: ", j, "item: ", item)

			// 	if _, ok := m0[item]; !ok {
			// 		m0[item] = true
			// 	} else {
			// 		return false
			// 	}
			// case [2]int{1, 1}:
			// 	fmt.Println("case: ", arr, "i: ", i, "j: ", j, "item: ", item)

			// 	if _, ok := m1[item]; !ok {
			// 		m1[item] = true
			// 	} else {
			// 		return false
			// 	}
			// case [2]int{2, 2}:
			// 	fmt.Println("case: ", arr, "i: ", i, "j: ", j, "item: ", item)

			// 	if _, ok := m2[item]; !ok {
			// 		m2[item] = true
			// 	} else {
			// 		return false
			// 	}
			// }
		}
	}
	// }

	return true
}

// [["5","3",".",".","7",".",".",".","."],
//  ["8",".",".","1","9","5",".",".","."],
//  [".","9","8",".",".",".",".","6","."],
//  ["8",".",".",".","6",".",".",".","3"],
//  ["4",".",".","8",".","3",".",".","1"],
//  ["7",".",".",".","2",".",".",".","6"],
//  [".","6",".",".",".",".","2","8","."],
//  [".",".",".","4","1","9",".",".","5"],
//  [".",".",".",".","8",".",".","7","9"]]
