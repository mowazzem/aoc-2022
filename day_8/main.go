package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	m, n := 0, 0

	input := [][]int{}

	br := bufio.NewScanner(f)
	for br.Scan() {
		line := br.Text()
		row := []int{}
		for _, c := range strings.Split(line, "") {
			digit, _ := strconv.Atoi(c)
			row = append(row, digit)
		}
		m = len(row)
		input = append(input, row)
	}

	n = len(input)

	// part1
	outerVisibleTrees := 2 * (m + n - 2)
	_ = outerVisibleTrees
	innerVisibleTrees := 0
	scenicScores := []int{}
	for i := 1; i < n-1; i++ {
		for j := 1; j < m-1; j++ {
			tree := input[i][j]
			leftvisible := true
			rightvisible := true
			topvisible := true
			bottomvisible := true

			lscene := 0
			rscene := 0
			tscene := 0
			bscene := 0

			for k := j - 1; k >= 0; k-- {
				left := input[i][k]
				if left >= tree {
					leftvisible = false
					lscene++
					break
				}
				lscene++
			}

			for k := j + 1; k < m; k++ {
				right := input[i][k]
				if right >= tree {
					rightvisible = false
					rscene++
					break
				}
				rscene++
			}

			for k := i - 1; k >= 0; k-- {
				top := input[k][j]
				if top >= tree {
					topvisible = false
					tscene++
					break
				}
				tscene++
			}

			for k := i + 1; k < n; k++ {
				bottom := input[k][j]
				if bottom >= tree {
					bottomvisible = false
					bscene++
					break
				}
				bscene++
			}

			scenicScores = append(scenicScores, lscene*rscene*tscene*bscene)

			if leftvisible || rightvisible || topvisible || bottomvisible {
				innerVisibleTrees++
			}
		}
	}

	// part1
	fmt.Println(outerVisibleTrees + innerVisibleTrees)

	// part2
	sort.Ints(scenicScores)
	fmt.Println(scenicScores[len(scenicScores)-1])
}
