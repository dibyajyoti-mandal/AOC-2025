package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"log"
	"math/big"
)

func pow10(k int) *big.Int {
	res := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(k)), nil)
	return res
}

func divCeil(a, b *big.Int) *big.Int {
	q := new(big.Int).Div(a, b)
	tmp := new(big.Int).Mul(q, b)
	if tmp.Cmp(a) < 0 {
		q.Add(q, big.NewInt(1))
	}
	return q
}

func max(a, b *big.Int) *big.Int {
	if a.Cmp(b) >= 0 {
		return new(big.Int).Set(a)
	}
	return new(big.Int).Set(b)
}
func min(a, b *big.Int) *big.Int {
	if a.Cmp(b) <= 0 {
		return new(big.Int).Set(a)
	}
	return new(big.Int).Set(b)
}

func processRanges(line string) *big.Int {
	total := big.NewInt(0)
	line = strings.TrimSpace(line)
	if line == "" {
		return total
	}
	parts := strings.Split(line, ",")
	for _, pr := range parts {
		pr = strings.TrimSpace(pr)
		if pr == "" {
			continue
		}
		bnds := strings.Split(pr, "-")
		
		a := new(big.Int)
		b := new(big.Int)
		if _, ok := a.SetString(strings.TrimSpace(bnds[0]), 10); !ok {
			log.Fatalf("bad number: %s", bnds[0])
		}
		if _, ok := b.SetString(strings.TrimSpace(bnds[1]), 10); !ok {
			log.Fatalf("bad number: %s", bnds[1])
		}
		maxDigits := len(b.Text(10))
		maxK := maxDigits/2
		for k := 1; k <= maxK; k++ {
			p10k := pow10(k)
			mult := new(big.Int).Add(p10k, big.NewInt(1)) // 10^k + 1

			tmin := divCeil(a, mult)
			tmax := new(big.Int).Div(b, mult)

			lower := pow10(k - 1)
			upper := new(big.Int).Sub(p10k, big.NewInt(1))

			tmin = max(tmin, lower)
			tmax = min(tmax, upper)

			if tmin.Cmp(tmax) > 0 {
				continue
			}

			// count = tmax - tmin + 1
			count := new(big.Int).Sub(tmax, tmin)
			count.Add(count, big.NewInt(1))

			// sum_t = (tmin + tmax) * count / 2
			sum_t := new(big.Int).Add(tmin, tmax)
			sum_t.Mul(sum_t, count)
			sum_t.Div(sum_t, big.NewInt(2))

			// add mult * sum_t to total
			add := new(big.Int).Mul(mult, sum_t)
			total.Add(total, add)
		}
	}
	return total
}

func computeMultiplier(l, r int) *big.Int {
	m := big.NewInt(0)
	for i := 0; i < r; i++ {
		m.Add(m, pow10(i*l))
	}
	return m
}

func isPrimitive(s string) bool {
	L := len(s)
	for d := 1; d*2 <= L; d++ {
		if L%d != 0 {
			continue
		}
		sub := s[:d]
		repeated := strings.Repeat(sub, L/d)
		if repeated == s {
			return false
		}
	}
	return true
}

func processRangesLine(line string) *big.Int {
	line = strings.TrimSpace(line)
	if line == "" {
		return big.NewInt(0)
	}
	parts := strings.Split(line, ",")
	total := big.NewInt(0)
	seen := make(map[string]struct{}) // dedupe numbers (as decimal strings)

	maxDigits := 1
	rangeBounds := make([][2]*big.Int, 0, len(parts))
	for _, pr := range parts {
		pr = strings.TrimSpace(pr)
		if pr == "" {
			continue
		}
		bnds := strings.Split(pr, "-")
		if len(bnds) != 2 {
			log.Fatalf("bad range: %q", pr)
		}
		a := new(big.Int)
		b := new(big.Int)
		if _, ok := a.SetString(strings.TrimSpace(bnds[0]), 10); !ok {
			log.Fatalf("bad number: %s", bnds[0])
		}
		if _, ok := b.SetString(strings.TrimSpace(bnds[1]), 10); !ok {
			log.Fatalf("bad number: %s", bnds[1])
		}
		rangeBounds = append(rangeBounds, [2]*big.Int{a, b})
		if len(b.Text(10)) > maxDigits {
			maxDigits = len(b.Text(10))
		}
	}

	for l := 1; l <= maxDigits/2; l++ {
		maxR := maxDigits / l
		for r := 2; r <= maxR; r++ {
			mult := computeMultiplier(l, r) // big.Int
			tLower := pow10(l - 1)
			tUpper := new(big.Int).Sub(pow10(l), big.NewInt(1))

			for _, bnds := range rangeBounds {
				a := bnds[0]
				b := bnds[1]

				// tmin = ceil(a / mult), tmax = floor(b / mult)
				tmin := divCeil(a, mult)
				tmax := new(big.Int).Div(b, mult)

				tmin = max(tmin, tLower)
				tmax = min(tmax, tUpper)

				if tmin.Cmp(tmax) > 0 {
					continue
				}

				for t := new(big.Int).Set(tmin); t.Cmp(tmax) <= 0; t.Add(t, big.NewInt(1)) {
					ts := t.Text(10)
					
					if !isPrimitive(ts) {
						continue
					}
					// n = mult * t
					n := new(big.Int).Mul(mult, t)
					key := n.Text(10)
					if _, ok := seen[key]; ok {
						continue
					}
					seen[key] = struct{}{}
					total.Add(total, n)
				}
			}
		}
	}
	return total
}



func p1(r *bufio.Reader, w *bufio.Writer) {
	line, err := r.ReadString('\n')
	if err != nil && line == "" {
		rest, _ := r.ReadString(0)
		line += rest
	}
	line = strings.TrimSpace(line)
	sum := processRanges(line)
	fmt.Fprintln(w, sum.String())
}

func p2(r *bufio.Reader, w *bufio.Writer) {
	line, err := r.ReadString('\n')
	if err != nil && line == "" {
		rest, _ := r.ReadString(0)
		line += rest
	}
	line = strings.TrimSpace(line)
	sum := processRangesLine(line)
	fmt.Fprintln(w, sum.String())

}

func main() {
	in, _ := os.Open("input1.txt")
	defer in.Close()
	out, _ := os.Create("output.txt")
	defer out.Close()

	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	p2(r, w);
}