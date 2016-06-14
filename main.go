// Program pathsort sorts the lines from stdin. Each lines is sorted as file path, with subdirectories placed after normal files.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

func readLines(f io.Reader) (lines []string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fatal(err)
	}
	return lines
}

func main() {
	lines := readLines(os.Stdin)
	sort.Sort(pathCompare(lines))
	for _, l := range lines {
		fmt.Println(l)
	}
}

func fatal(a ...interface{}) {
	fmt.Fprintln(os.Stderr, a...)
	os.Exit(1)
}

type pathCompare []string

func (lst pathCompare) Len() int      { return len(lst) }
func (lst pathCompare) Swap(i, j int) { lst[i], lst[j] = lst[j], lst[i] }
func (lst pathCompare) Less(i, j int) bool {
	a, b := lst[i], lst[j]
	x, y := stripSharedPrefix(a, b)
	xdir := strings.Contains(x, "/")
	ydir := strings.Contains(y, "/")
	if !xdir && ydir {
		return true
	}
	if xdir && !ydir {
		return false
	}
	return a < b
}

func stripSharedPrefix(a, b string) (x, y string) {
	i := 0
	for {
		if i == len(a) || !strings.HasPrefix(b, a[:i+1]) {
			return a[i:], b[i:]
		}
		i++
	}
	return
}
