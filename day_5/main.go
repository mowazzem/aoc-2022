package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	br := bufio.NewScanner(f)

	procedureStarted := false
	stacks := []string{}
	mp := make(map[int][]string)
	mp2 := make(map[int][]string)

	order := []int{}
	for br.Scan() {
		line := br.Text()
		if line == "" {
			procedureStarted = true

			mxln := len(stacks[len(stacks)-1])
			for i, v := range stacks {
				x := &stacks[i]
				for j := 0; j < mxln-len(v); j++ {
					*x += " "
				}
			}

			for _, nstr := range strings.Split(stacks[len(stacks)-1], " ") {
				n, _ := strconv.Atoi(nstr)
				order = append(order, n)
			}

			stacks = stacks[:len(stacks)-1]

			inc := 1
			for i := 0; i < mxln; i++ {
				if i%2 != 0 {
					continue
				}
				for j := len(stacks) - 1; j >= 0; j-- {
					c := string(stacks[j][i])
					if c == " " {
						continue
					}

					vs := mp[inc]
					vs = append(vs, c)
					mp[inc] = vs

					vss := mp2[inc]
					vss = append(vss, c)
					mp2[inc] = vss
				}
				inc++
			}

			continue
		}

		if !procedureStarted {
			parsed := ""
			for i, v := range line {
				x := string(v)
				if i%2 != 0 {
					parsed += string(x)
				}
			}
			if parsed != "" {
				stacks = append(stacks, parsed)
			}
			continue
		}

		proc := strings.Split(line, " ")
		c, _ := strconv.Atoi(proc[1])
		frm, _ := strconv.Atoi(proc[3])
		to, _ := strconv.Atoi(proc[5])

		v := mp[frm]
		v1 := v[:len(v)-c]
		mp[frm] = v1

		vv := mp2[frm]
		vv1 := vv[:len(vv)-c]
		mp2[frm] = vv1

		ex := v[len(v)-c:]

		ex2 := vv[len(vv)-c:]

		x := mp[to]
		for i := len(ex) - 1; i >= 0; i-- {
			x = append(x, ex[i])
		}
		mp[to] = x

		mp2[to] = append(mp2[to], ex2...)

	}

	res := ""
	res2 := ""
	for _, o := range order {
		v := mp[o]
		res += v[len(v)-1]

		vv := mp2[o]
		res2 += vv[len(vv)-1]
	}

	// part 1
	fmt.Println(res)

	// part2
	fmt.Println(res2)

}
