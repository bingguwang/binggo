package od

import "fmt"

/*
公司老板做了一笔大生意，想要给每位员工分配一些奖金，想通过游戏的方式来决定每个人分多少钱。按照员工的工号顺序，每个人随机抽取一个数字。按照工号的顺序往后排列，遇到第一个数字比自己数字大的，那么，前面的员工就可以获得“距离*数字差值”的奖金。如果遇不到比自己数字大的，就给自己分配随机数数量的奖金。例如，按照工号顺序的随机数字是：2,10,3。那么第2个员工的数字10比第1个员工的数字2大，所以，第1个员工可以获得1*（10-2）=8。第2个员工后面没有比他数字更大的员工，所以，他获得他分配的随机数数量的奖金，就是10。第3个员工是最后一个员工，后面也没有比他更大数字的员工，所以他得到的奖金是3。

请帮老板计算一下每位员工最终分到的奖金都是多少钱。
 */

func main() {
    // 员工数量
    n := 4

    // 随机数
    randomNumbers := []int{2, 5, 10, 3}
    fmt.Println("员工随机数字:", randomNumbers)

    // 初始化员工奖金
    bonuses := make([]int, n)

    // 计算每个员工的奖金
    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {
            if randomNumbers[j] > randomNumbers[i] {
                bonuses[i] = (j - i) * (randomNumbers[j] - randomNumbers[i])
                break
            }
        }
        // 如果向右遍历，没有j的值大于i，就选随机数
        if bonuses[i] == 0 {
            bonuses[i] = randomNumbers[i]
        }
    }
    fmt.Println("每位员工的奖金:", bonuses)
}
