package od

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"testing"
)

/**
给定一个字符串描述的算术表达式，计算出结果值。

输入字符串长度不超过 100 ，合法的字符包括 ”+, -, *, /, (, )” ， ”0-9” 。
*/

func isDig(b byte) bool {
	if b == '-' || b == '+' || b == '*' || b == '/' || b == '(' || b == ')' {
		return false
	}
	return true
}

func tokenize(s string) []string {
	var tokens []string
	i := 0
	for i < len(s) {
		if s[i] == '-' && i+1 < len(s) && isDig(s[i+1]) {
			start := i
			// 如果-前只要不是左括号，就是说-得被视为减号
			if i-1 >= 0 && s[i-1] != '(' {
				tokens = append(tokens, string('-'))
				start++
			}
			i++
			for i < len(s) && isDig(s[i]) {
				i++
			}
			fmt.Println("B:", s[start:i])
			tokens = append(tokens, s[start:i])
		} else if !isDig(s[i]) {
			tokens = append(tokens, string(s[i]))
			i++
		} else {
			start := i
			i++
			for i < len(s) && isDig(s[i]) {
				i++
			}
			fmt.Println("A:", s[start:i])
			tokens = append(tokens, s[start:i])
		}
	}
	fmt.Println(tokens)
	return tokens
}

func TestGSD(t *testing.T) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()

	// 把输入字符串切割为string数组
	input := tokenize(text)
	fmt.Println(len(input))
	// 开始计算
	postfix := infixToPostfix(input)
	res := evaluatePostfix(postfix)
	fmt.Println(res)
}
func evaluatePostfix(inputs []string) int {
	stack := make([]int, 0)

	// 遍历后缀表达式
	for _, token := range inputs {
		// 数字直接入栈
		if isDigit(token) {
			// 如果是数字，将其转换为整数并压入栈中
			num, err := strconv.Atoi(token)
			if err != nil {
				return 0
			}
			stack = append(stack, num)
		} else {
			// 如果是运算符，从栈中弹出两个数字进行运算，运算结果再入栈
			op2 := stack[len(stack)-1]
			op1 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			switch token {
			case "+":
				stack = append(stack, op1+op2)
			case "-":
				stack = append(stack, op1-op2)
			case "*":
				stack = append(stack, op1*op2)
			case "/":
				stack = append(stack, op1/op2)
			}
		}
	}

	// 最终栈中只剩下一个数字，即为表达式的计算结果
	return stack[0]
}

func isDigit(token string) bool {
	_, err := strconv.Atoi(token)
	return err == nil
}

func infixToPostfix(infix []string) []string {
	var postfix []string
	var stack []string

	precedence := map[string]int{
		"+": 1, "-": 1,
		"*": 2, "/": 2,
	}

	for _, token := range infix {
		switch {
		case token == "(": // 左括号直接出栈
			stack = append(stack, token)
		case token == ")": // 当前是右括号则出栈，直到遇到(
			for len(stack) > 0 && stack[len(stack)-1] != "(" {
				postfix = append(postfix, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			// 右括号直接匹配栈顶的左括号，左括号直接出栈
			if len(stack) > 0 && stack[len(stack)-1] == "(" {
				stack = stack[:len(stack)-1]
			}
		case precedence[token] > 0: // 四则运算
			// 优先级小于等于栈顶元素。出栈元素直到遇到大于的
			for len(stack) > 0 && precedence[token] <= precedence[stack[len(stack)-1]] {
				postfix = append(postfix, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			// 出栈完成后入栈当前元素
			stack = append(stack, token)
		default:
			// 其他的就是数字，直接输出
			postfix = append(postfix, token)
		}
	}
	// 元素都遍历完了，把站内元素输出完
	for len(stack) > 0 {
		postfix = append(postfix, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	return postfix
}
