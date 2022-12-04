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

	fullyOverlapPairs := 0
	anyOverlapPair := 0

	br := bufio.NewScanner(f)
	for br.Scan() {
		line := br.Text()
		secs := strings.Split(line, ",")
		sec1 := secs[0]
		sec2 := secs[1]

		sec1Min, _ := strconv.Atoi(sec1[:strings.Index(sec1, "-")])
		sec1Max, _ := strconv.Atoi(sec1[strings.Index(sec1, "-")+1:])

		sec2Min, _ := strconv.Atoi(sec2[:strings.Index(sec2, "-")])
		sec2Max, _ := strconv.Atoi(sec2[strings.Index(sec2, "-")+1:])

		if (sec2Min >= sec1Min && sec2Max <= sec1Max) || (sec1Min >= sec2Min && sec1Max <= sec2Max) {
			fullyOverlapPairs++
		}

		if (sec1Max >= sec2Min && sec1Min <= sec2Max) || (sec1Min >= sec2Max && sec1Max <= sec2Min) {
			anyOverlapPair++
		}

	}

	// part 1
	fmt.Println(fullyOverlapPairs)

	// part 2
	fmt.Println(anyOverlapPair)
}
