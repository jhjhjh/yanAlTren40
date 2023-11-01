package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"runtime"
)

func process(str string) map[rune]int {
	result := make(map[rune]int)
	for _, char := range str {
		if _, found := result[char]; !found {
			result[char] = 1
		} else {
			tmp := result[char]
			result[char] = tmp + 1
		}
	}
	return result

}

func DoInput(in io.Reader, out io.Writer) {
	var str1 string
	var str2 string
	fmt.Fscan(in, &str1)
	fmt.Fscan(in, &str2)
	map1 := process(str1)
	map2 := process(str2)
	if len(map1) != len(map2) {
		fmt.Fprintln(out, "NO")
		return
	} else {
		for key1, value1 := range map1 {
			if map2[key1] != value1 {
				fmt.Fprintln(out, "NO")
				return
			}
		}
	}

	fmt.Fprintln(out, "YES")

}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	DoInput(in, out)
}
