// 動的計画法
// 1本の金属棒からいくつかの金属棒を切り出して売る.
// 長さ i インチの金属棒の価格 Pi は以下の価格表により決まる.
// また一回の切り出しには定数コスト c がかかるとする.
// 売値は切り出した金属棒の価格 Pi の合計から切り出した回数 * c のコストを差し引いたものである.
// 長さ i  | 1 | 2 | 3 | 4 | 5  | 6  | 7  | 8  | 9  | 10 |
// 価格 Pi | 1 | 5 | 8 | 9 | 10 | 17 | 17 | 20 | 24 | 30 |
// 長さ r の金属棒が与えられたとき売値を最大化する切り出し方を求めよ.

package main

import (
	"fmt"
	"os"
)

// ２つの整数のうち大きい方を返す関数
func max(x, y int) int {
	if x <= y {
		return y
	}
	return x
}

// n: 切り出す前の金属棒

func cutRod(n, c int, p, m []int) int {
	// 答えが履歴に格納されていた場合その値を返す
	fmt.Println(n)
	if m[n] >= 0 {
		return m[n]
	}

	var q int
	// n = 0 になったら 0 を返す
	if n == 0 {
		q = 0
	} else {
		q = -10000
		for i := 1; i <= n; i++ {
			q = max(q, p[i]-c+cutRod(n-i, c, p, m))
		}
	}

	m[n] = q
	return q
}

func main() {
	// n: 切り出す前の金属棒
	// c: コスト
	var n, c int
	fmt.Scanf("%d %d", &n, &c)

	if n == 0 {
		fmt.Println(0)
		os.Exit(0)
	}
	// 価格を格納した配列長さ i の金属棒の価格 p[i]
	var p = []int{0, 1, 5, 8, 9, 10, 17, 17, 20, 24, 30}

	// 部分問題の答えを格納する配列
	m := make([]int, n+1)
	for i := 0; i <= n; i++ {
		fmt.Println(i)
		m[i] = -10000
	}
	fmt.Println(cutRod(n, c, p, m))
}
