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

	br := bufio.NewScanner(f)

	ds := ""
	for br.Scan() {
		ds = br.Text()
	}

	// window size
	packetMarkerSize := 4
	messageMarkerSize := 14

	// part1
	for i := 0; i <= len(ds)-packetMarkerSize; i++ {
		marker := ds[i : i+packetMarkerSize]
		found := true
		for _, v := range marker {
			c := string(v)
			if strings.Count(marker, c) > 1 {
				found = false
				break
			}
		}
		if found {
			fmt.Println(i + packetMarkerSize)
			break
		}
	}

	// part2
	for i := 0; i <= len(ds)-messageMarkerSize; i++ {
		marker := ds[i : i+messageMarkerSize]
		found := true
		for _, v := range marker {
			c := string(v)
			if strings.Count(marker, c) > 1 {
				found = false
				break
			}
		}
		if found {
			fmt.Println(i + messageMarkerSize)
			break
		}
	}
}
