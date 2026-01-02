package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func p1(r *bufio.Reader, w *bufio.Writer) {
	
	var pos , ans int
	pos = 50
	ans = 0

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		dir := line[0]
		val, _ := strconv.Atoi(line[1:])

		if dir == 'L' {
			pos = (pos - val) % 100
		} else { // 'R'
			pos = (pos + val) % 100
		}

		if pos < 0 {
			pos += 100
		}

		if pos == 0 {
			ans++
		}
	}

	fmt.Fprintln(w, ans)
}

func p2(r *bufio.Reader, w *bufio.Writer) {
    var pos, ans int
    pos = 50
    ans = 0

    for {
        line, err := r.ReadString('\n')
        if err != nil {
            break
        }
        line = strings.TrimSpace(line)
        if line == "" {
            continue
        }

        dir := line[0]
        val, _ := strconv.Atoi(line[1:])

        if dir == 'L' {
            first := pos
            if pos == 0 {
                first = 100
            }
            if val >= first {
                hits := 1 + (val-first)/100
                ans += hits
            }
            pos = ((pos - val) % 100 + 100) % 100
        } else { // 'R'
            first := 100 - pos
            if pos == 0 {
                first = 100
            }
            if val >= first {
                hits := 1 + (val-first)/100
                ans += hits
            }
            pos = (pos + val) % 100
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

	p2(r, w);
}
