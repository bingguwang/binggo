package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	numscnt, colorcnt = make(map[string]int, 0), make(map[string]int, 0)
	var nums []string
	var colors []string

	for i := 0; i < 5; i++ {
		s.Scan()
		s2 := strings.Fields(s.Text())
		nums = append(nums, s2[0])
		colors = append(colors, s2[1])
		numscnt[s2[0]]++
		colorcnt[s2[1]]++
	}
	fmt.Println(nums)
	fmt.Println(numscnt)
	fmt.Println(colorcnt)

	if tonghua() && shunzi() {
		fmt.Print(1)
	} else if sitiao() {
		fmt.Print(2)
	} else if hulu() {
		fmt.Print(3)
	} else if tonghua() {
		fmt.Print(4)
	} else if shunzi() {
		fmt.Print(5)
	} else if santiao() {
		fmt.Print(6)
	}
}

var numberMap = map[string]int{
	"2": 2, "3": 3,
	"4": 4, "5": 5,
	"6": 6, "7": 7,
	"8": 8, "9": 9,
	"10": 10, "J": 11,
	"Q": 12, "K": 13,
	"A": 14,
}

var (
	colorcnt map[string]int
	numscnt  map[string]int
)

func tonghua() bool {
	return len(colorcnt) == 1
}

func hulu() bool {
	var istwo, isthree bool
	for _, count := range numscnt {
		if count == 2 {
			istwo = true
		}
		if count == 3 {
			istwo = true
		}
	}
	return isthree && istwo
}

func shunzi() bool {
	var nums []int
	for k, _ := range numscnt {
		nums = append(nums, numberMap[k])
	}
	sort.Ints(nums)

	return len(nums) == 5 &&
		(nums[len(nums)-1]-nums[0] == 4 || fmt.Sprint(nums) == "[2,3,4,5,14]")

}
func santiao() bool {
	threecount, singlecount := 0, 0
	for _, count := range numscnt {
		if count == 3 {
			threecount++
		}
		if count == 1 {
			singlecount++
		}
	}
	return threecount == 3 && singlecount == 2
}

func sitiao() bool {
	for _, count := range numscnt {
		if count == 4 {
			return true
		}
	}
	return false
}
