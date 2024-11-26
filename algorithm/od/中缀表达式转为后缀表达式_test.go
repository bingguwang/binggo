package od

import "testing"
/**
    中缀表达式 -》 后缀表达式
 */
func TestGalw(t *testing.T) {

}
func getBack(str string) string {
    var res []byte
    var stack []byte
    for i := 0; i < len(str); {
        s := str[i]
        //为数字时直接压栈并将索引i向后移动
        if s <= 'z' && s >= 'a' || s <= '9' && s >= '0' {
            res = append(res, s)
            i++
        } else if s == '(' {
            //为左括号时直接压栈并将索引i向后移动
            stack = append(stack, s)
            i++
        } else if s == ')' {
            //为右括号时依次弹栈并存入列表s2，直到栈顶为左括号
            for stack[len(stack)-1] != '(' {
                res = append(res, stack[len(stack)-1])
                stack = stack[:len(stack)-1]
            }
            //将栈顶的左括号弹出
            stack = stack[:len(stack)-1]
            //索引向后移动
            i++
        } else {
            //当栈为空或者栈顶为 “(” 或运算符优先级大于栈顶运算符的优先级时直接压栈，并将索引向后移动。(此处priorty()优先级判断的方法很简单自行编写)
            if len(stack) == 0 || stack[len(stack)-1] == '(' || priority(s) > priority(stack[len(stack)-1]) {
                stack = append(stack, s)
                i++
            } else {
                //当小于或等于栈顶运算符优先级时，将栈顶运算符弹出并存入列表s2中,此时索引不移动,以便让该运算符继续判断。
                res = append(res, stack[len(stack)-1])
                stack = stack[:len(stack)-1]
            }
        }
    }

    for i := len(stack) - 1; i >= 0; i-- {
        res = append(res, stack[i])
    }
    return string(res)
}

func priority(a byte) int {
    switch a {
    case '+':
        return 1
    case '-':
        return 1
    case '*':
        return 2
    case '/':
        return 2
    case '(':
        return 3
    }
    return 0
}
