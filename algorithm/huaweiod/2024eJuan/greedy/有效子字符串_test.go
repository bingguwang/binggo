package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str1 := scanner.Text()
	scanner.Scan()
	str2 := scanner.Text()

	l1, l2 := 0, 0
	for l1 < len(str1) && l2 < len(str2) {
		if str1[l1] == str2[l2] {
			l1++
			l2++
		} else {
			l2++
		}
	}
	if l1 == len(str1) {
		fmt.Println(str2[:l2-1])
		fmt.Println(l2 - 1)
	} else {
		fmt.Println(-1)
	}

}
