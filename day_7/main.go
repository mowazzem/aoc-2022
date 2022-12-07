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

type statement string

func (stmnt statement) IsCommand() bool {
	if string(stmnt[0]) == "$" {
		return true
	}

	return false
}

func (stmnt statement) isDir() bool {
	ss := strings.Split(string(stmnt), " ")
	if len(ss) == 0 {
		return false
	}

	return ss[0] == "dir"
}

func (stmnt statement) dirName() string {
	return strings.Split(string(stmnt), " ")[1]
}

func (stmnt statement) fileSize() int {
	s := strings.Split(string(stmnt), " ")
	size, _ := strconv.Atoi(s[0])

	return size
}

type commandType string

const (
	tCd commandType = "cd"
	tLs commandType = "ls"

	limit      int = 100000
	filesystem int = 70000000
	required   int = 30000000
)

type command struct {
	tYpe commandType
	arg  string
}

func newCommand(s statement) command {
	cmd := command{}
	if s[2:4] == "ls" {
		cmd.tYpe = tLs
		return cmd
	}

	cmd.tYpe = tCd
	cmd.arg = string(s[5:])
	return cmd
}

type directory struct {
	name   string
	size   int
	parent *directory
	childs directories
}

type directories []*directory

func (dir *directory) Walk(total *int, sizeList *[]int) {
	for _, chld := range dir.childs {
		chld.Walk(total, sizeList)
		dir.size += chld.size
	}
	if dir.size <= limit {
		*total += dir.size
	}
	if dir.parent != nil {
		*sizeList = append(*sizeList, dir.size)
	}
}

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	br := bufio.NewScanner(f)

	inputs := []string{}
	for br.Scan() {
		inputs = append(inputs, br.Text())
	}

	var currentDirectory *directory
	var root *directory

	for _, line := range inputs {
		stmnt := statement(line)
		if stmnt.IsCommand() {
			cmd := newCommand(stmnt)
			switch cmd.tYpe {
			case tLs:
				continue
			case tCd:
				if cmd.arg == ".." {
					currentDirectory = currentDirectory.parent
				} else if cmd.arg == "/" {
					root = &directory{
						name:   cmd.arg,
						parent: nil,
					}
					currentDirectory = root
				} else {
					for _, child := range currentDirectory.childs {
						if child.name == cmd.arg {
							currentDirectory = child
							break
						}
					}
				}
			}

			continue
		}

		if stmnt.isDir() {
			dir := &directory{
				name:   stmnt.dirName(),
				parent: currentDirectory,
			}
			currentDirectory.childs = append(currentDirectory.childs, dir)
			continue
		}

		currentDirectory.size += stmnt.fileSize()
	}

	total := 0
	sizeList := []int{}
	root.Walk(&total, &sizeList)

	// part1
	fmt.Println(total)

	// part2
	sl := []int{}
	diff := filesystem - root.size
	for _, size := range sizeList {
		if size+diff >= required {
			sl = append(sl, size)
		}
	}
	sort.Ints(sl)
	fmt.Println(sl[0])

}
