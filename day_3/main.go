package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	smll := make(map[string]int, 26)
	for i, j := 65, 27; i <= 90; i, j = i+1, j+1 {
		smll[string(i)] = j
	}

	cptl := make(map[string]int, 26)
	for i, j := 97, 1; i <= 122; i, j = i+1, j+1 {
		cptl[string(i)] = j
	}

	f, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	totalPriority := 0

	groupOfThrees := []string{}
	groupCount := 0

	listOfGroups := [][]string{}

	br := bufio.NewScanner(f)

	for br.Scan() {
		line := br.Text()
		ln := len(line)

		done := map[string]struct{}{}

		if groupCount == 3 {
			listOfGroups = append(listOfGroups, groupOfThrees)
			groupOfThrees = []string{}
			groupCount = 0
		}

		groupOfThrees = append(groupOfThrees, line)
		groupCount++

		// first half
		fh := line[:ln/2]

		// second half
		sh := line[ln/2:]

		for _, v := range fh {
			c := string(v)
			if _, ok := done[c]; ok {
				continue
			}
			done[c] = struct{}{}

			if strings.Contains(sh, c) {
				if p, ok := smll[c]; ok {
					totalPriority += p
					continue
				}

				if p, ok := cptl[c]; ok {
					totalPriority += p
				}
			}
		}

	}

	if len(groupOfThrees) > 0 {
		listOfGroups = append(listOfGroups, groupOfThrees)
	}

	// part 1
	fmt.Println(totalPriority)

	totalGroupPoint := 0
	for _, g := range listOfGroups {
		gt := 0
		done := map[string]struct{}{}
		for _, v := range g[0] {
			c := string(v)
			if _, ok := done[c]; ok {
				continue
			}
			done[c] = struct{}{}

			if strings.Contains(g[1], c) && strings.Contains(g[2], c) {
				if p, ok := smll[c]; ok {
					gt += p
					continue
				}

				if p, ok := cptl[c]; ok {
					gt += p
				}
			}
		}

		totalGroupPoint += gt
	}

	// part 2
	fmt.Println(totalGroupPoint)
}
