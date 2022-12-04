package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	br := bufio.NewScanner(f)

	calories := []int{}

	total := 0

	for br.Scan() {
		line := br.Text()
		if line == "" {
			calories = append(calories, total)
			total = 0
			continue
		}
		calorie, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		total += calorie
	}

	if len(calories) == 0 {
		return
	}

	sort.Ints(calories)

	topThreeCals := 0
	for _, v := range calories[len(calories)-3:] {
		topThreeCals += v
	}

	//part 1
	fmt.Println(calories[len(calories)-1])

	//part 2
	fmt.Println(topThreeCals)

}
