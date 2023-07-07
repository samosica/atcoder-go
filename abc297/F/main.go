package main

import (
	"fmt"
	"math/bits"
)

func Extgcd(a, b int) (s, t, g int) {
	s, t = 1, 0
	s2, t2 := 0, 1
	for b != 0 {
		a, b, s, t, s2, t2 = b, a%b, s2, t2, s-(a/b)*s2, t-(a/b)*t2
	}
	g = a
	return
}

func Modinv(n, mod int) int {
	inv_n, _, _ := Extgcd(n, mod)
	if inv_n < 0 {
		inv_n += mod
	}
	return inv_n
}

func ComputeFactorial(n, mod int) (fact, invFact []int) {
	fact = make([]int, n+1)
	invFact = make([]int, n+1)

	fact[0] = 1
	for i := 1; i <= n; i++ {
		fact[i] = fact[i-1] * i % mod
	}

	invFact[n] = Modinv(fact[n], mod)
	for i := n; i > 0; i-- {
		invFact[i-1] = invFact[i] * i % mod
	}

	return
}

func main() {
	var h, w, k int
	fmt.Scan(&h, &w, &k)

	const mod = 998244353
	fact, invFact := ComputeFactorial(h*w, mod)

	comb := func(n, k int) int {
		if n < 0 || k < 0 || n < k {
			return 0
		}
		return fact[n] * invFact[k] % mod * invFact[n-k] % mod
	}

	table := make([][]int, h+1)

	for i := range table {
		table[i] = make([]int, w+1)
		for j := range table[i] {
			table[i][j] = comb(i*j, k)
		}
	}

	e := 0
	invTotal := Modinv(table[h][w], mod)
	for i := 1; i <= h; i++ {
		for j := 1; j <= w; j++ {
			p := 0
			for l := uint(0); l < 16; l++ {
				i2 := i - bits.OnesCount(l&0b1100)
				j2 := j - bits.OnesCount(l&0b0011)
				if i2 >= 0 && j2 >= 0 {
					if bits.OnesCount(l)%2 == 0 {
						p += table[i2][j2]
						if p >= mod {
							p -= mod
						}
					} else {
						p -= table[i2][j2]
						if p < 0 {
							p += mod
						}
					}
				}
			}
			p = p * (h - i + 1) * (w - j + 1) % mod
			p = p * invTotal % mod
			e += p * i * j % mod
			if e >= mod {
				e -= mod
			}
		}
	}

	fmt.Println(e)
}
