package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	me := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	op := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}

	pt1Total := 0

	pt2Total := 0

	br := bufio.NewScanner(f)
	for br.Scan() {
		line := br.Text()
		if line == "" {
			continue
		}

		vs := strings.Split(line, " ")

		// opponent hand
		opsc := op[vs[0]]

		// my hand
		msc := me[vs[1]]

		// part1
		if opsc == msc {
			pt1Total += (msc + 3)
		} else if msc-opsc == 1 || msc-opsc == -2 {
			pt1Total += (msc + 6)
		} else {
			pt1Total += msc
		}

		// part2
		if msc == 2 {
			pt2Total += opsc + 3
		} else if msc == 1 {
			if opsc == 1 {
				pt2Total += 3
			} else {
				pt2Total += opsc - 1
			}
		} else {
			if opsc == 3 {
				pt2Total += 1 + 6
			} else {
				pt2Total += opsc + 1 + 6
			}
		}
	}

	// part 1
	fmt.Println(pt1Total)

	// part 2
	fmt.Println(pt2Total)
}
