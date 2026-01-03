package main

import (
	"bufio"
	"fmt"
	"os"
	"io"
	"strings"
	"strconv"
)

func p1(r *bufio.Reader, w *bufio.Writer) {
	var ans int
	ans = 0
	for{
		line , err := r.ReadString('\n')
		if err != nil && err != io.EOF{
			break;
		}
		suf := make(map[byte]int)
		var curr int
		curr = 0
		n := len(line)

		for i:=n-1;i>=0; i--{
			suf[line[i]]++;
		}

		for i:=0; i<n; i++{
			suf[line[i]]--;
			var num int

			for d:=byte('9'); d>='0';d--{
				if(suf[d] > 0){
					num = int(line[i]-'0')*10 + int(d-'0')
					break;
				}
			}
			if(curr < num){
				curr = num
			}
		}
		// fmt.Fprintln(w, curr)

		ans += curr
		if err == io.EOF {
        	break
    	}
	}

	fmt.Fprintln(w , ans)
}

func p2(r *bufio.Reader, w *bufio.Writer) {
	const k = 12
	var ans int64 = 0

	for {
		line, err := r.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}
		s := strings.TrimSpace(line)
		if len(s) == 0 {
			if err == io.EOF {
				break
			}
			continue
		}

		n := len(s)
		var chosen string
		if n <= k {
			chosen = s
		} else {
			start := 0
			res := make([]byte, 0, k)
			for rem := k; rem > 0; rem-- {
				end := n - rem
				maxi := byte('0')
				maxPos := -1
				for j := start; j <= end; j++ {
					if s[j] > maxi {
						maxi = s[j]
						maxPos = j
						if maxi == '9' {
							// can't beat 9
							break
						}
					}
				}
				if maxPos == -1 {
					maxPos = start
					maxi = s[start]
				}
				res = append(res, maxi)
				start = maxPos + 1
			}
			chosen = string(res)
		}

		val, perr := strconv.ParseInt(chosen, 10, 64)
		if perr != nil {
			if err == io.EOF {
				break
			}
			continue
		}
		ans += val

		if err == io.EOF {
			break
		}
	}

	fmt.Fprintln(w, ans)
}

func main() {
	in, _ := os.Open("input1.txt")
	defer in.Close()
	out, _ := os.Create("output.txt")
	defer out.Close()

	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	p2(r, w)
}