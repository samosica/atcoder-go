package main

import "fmt"

func ModPow(x, n, mod int64) int64 {
	p := int64(1)
	for n > 0 {
		if n&1 > 0 {
			p = p * x % mod
		}
		x = x * x % mod
		n >>= 1
	}
	return p
}

func main() {
	var N, M int64
	fmt.Scan(&N, &M)

	const MOD = 998_244_353
	A := make([]int64, N+1)
	A[0] = 1
	A[1] = M
	A[2] = M * (M - 1) % MOD
	if N >= 3 {
		A[3] = M * (M - 1) % MOD * (M - 2) % MOD
	}
	for n := int64(4); n <= N; n += 1 {
		// n - 1 人の数を決める
		// 残った1人の両隣の数が等しいかどうかで場合分けを行なう

		c := M * ModPow(M-1, n-2, MOD) % MOD
		c1 := A[n-2]
		c2 := c - c1
		if c2 < 0 {
			c2 += MOD
		}

		A[n] = (c1*(M-1) + c2*(M-2)) % MOD
	}

	fmt.Println(A[N])
}
