package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		if n == 0 {
			break
		}
		number := n
		var ans int
		for {
			if number < 3 {
				if (number+1)/3 == 1 {
					ans++
				}
				break
			}
			ans += number / 3 // 换汽水
			// fmt.Println("ans:", ans)

			number = number%3 + number/3 // 新的瓶子树
		}
		fmt.Println(ans)
	}
}
