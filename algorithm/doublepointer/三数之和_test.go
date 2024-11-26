package doublepointer

import (
    "fmt"
    "sort"
    "testing"
)

func TestNamesad(t *testing.T) {
    youxu := Hj([]int{-1, 0, 1, 2, -1, -4}, 0)
    fmt.Println(youxu)
}


// 三指针 ,中间值为基准
func Hj(nums []int, target int) [][]int {
    sort.Ints(nums)
    fmt.Println("排序结果：", nums)
    res, length := make([][]int, 0), len(nums)
    // index所指的数作为中间的值
    for index := 1; index < length-1; index++ {
        l, r := 0, length-1
        // 在这里跳过相同的基准数字，避免出现重复的结果
        if nums[index] == nums[index-1] {
            l = index - 1 // 不是直接跳过，直接跳过会丢失结果的
        }
        for l < index && r > index {
            if l > 0 && nums[l] == nums[l-1] { // 避免重复
                l++
                continue
            }
            if r < length-1 && nums[r] == nums[r+1] { // 避免重复
                r--
                continue
            }
            addNum := nums[l] + nums[r] + nums[index]
            if addNum == target {
                res = append(res, []int{nums[l], nums[index], nums[r]}) // 这时候由于是有序的，所以只l++或者只是r--都不能再得到目标值了（不允许结果有重复）
                l++
                r--
            } else if addNum > target {
                r--
            } else {
                l++
            }
        }
    }
    return res
}

// 三指针， 左边值为基准 ，对于输出结果有顺序要求的，就可以使用的这种
func find(numbers []int, target int) [][]int {
    var res [][]int
    sort.Ints(numbers) // 排序后，后面只要移动指针就可以控制三数之和的增减，而且可以避免重复结果
    fmt.Println(numbers)
    for i := 0; i < len(numbers)-1; i++ {
        // 在这里跳过相同的基准数字，避免出现重复的结果
        if i-1 >= 0 && numbers[i-1] == numbers[i] {
            continue // 之前执行过了，跳过
        }

        l, r := i+1, len(numbers)-1
        for l < r && l != r {
            if numbers[l] == numbers[l-1] && i < l-1 { // i<l-1是防止丢失结果
                l++
                continue
            }
            if r+1 < len(numbers) && numbers[r+1] == numbers[r] {
                r--
                continue
            }

            sum := numbers[l] + numbers[r] + numbers[i]
            if sum == target {
                res = append(res, []int{numbers[i], numbers[l], numbers[r]}) // 这时候由于是有序的，所以只l++或者只是r--都不能再得到目标值了（不允许结果有重复）
                l++
                r--
            } else if sum < target {
                l++
            } else {
                r--
            }
        }

    }
    return res
}