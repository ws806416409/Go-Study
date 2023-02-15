/**
 * Copyright: ws-test.All rights reserved.
 *
 * @Description
 * @Date 2023/2/15 15:43
 * @Author DizzyLime
 */
package main

import "fmt"

func main() {
	fmt.Println("WS-Go-TEST")
	ans := partition("abc")
	fmt.Print(ans)
}

func isPalindrome(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func partition(s string) (ans [][]string) {
	path := []string{}
	n := len(s)
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, append([]string(nil), path...)) // 固定答案
			return
		}
		for j := i; j < n; j++ { // 枚举子串的结束位置
			if isPalindrome(s, i, j) {
				path = append(path, s[i:j+1])
				dfs(j + 1)
				path = path[:len(path)-1] // 恢复现场
			}
		}
	}
	dfs(0)
	return
}
