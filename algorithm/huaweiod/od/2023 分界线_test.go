package od

import (
    "bufio"
    "fmt"
    "os"
    "sort"
    "strings"
)

/**
   电视剧《分界线》里面有一个片段，男主为了向警察透露案件细节，且不暴露自己， 于是将报刊上的字剪切下来，剪拼成匿名信。
现在有一名举报人，希望借鉴这种手段，使用英文报刊完成举报操作。
但为了增加文章的混淆度，只需满足每个单词中字母数量一致即可，不关注每个字母 的顺序。解释：单词’on’允许通过单词’no’进行替代
报纸代表newspaper,匿名信代表anonymousLetter,求报纸内容是否可以拼成匿 名信。
第一行输入newspaper内容，包括1-N个字符串，用空格分开
第二行输入anonymousLetter内容，包括1-N个字符串，用空格分开
1、newspaper和anonymousLetter的字符串由小写英文字母组成且每个字母只能 使用一次
2、newspaper内容中的每个字符串字母顺序可以任意调整，但必须保证字符串的完 整性（每个字符串不能有多余字母）
3. 1<N<100,1<= newspaper.length, anonymousLetter.length <= 104
*/
func main() {
    var newspaper, anonymousLetter []string

    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    text := scanner.Text()
    newspaper = strings.Split(text, " ")

    scanner.Scan()
    s := scanner.Text()
    anonymousLetter = strings.Split(s, " ")

    res := funcName(anonymousLetter, newspaper)
    fmt.Println(res)
}

func funcName(anonymousLetter []string, newspaper []string, ) bool {
    for i := 0; i < len(anonymousLetter); i++ {
        for j := 0; j < len(newspaper); j++ {
            news := newspaper[j]
            if len(anonymousLetter[i]) == len(news) && handle(news, anonymousLetter[i]) {

                fmt.Println(news, "和", anonymousLetter[i], "匹配成功")

                //长度相等才有资格匹配
                newspaper[j] = " " //使用过的字符串之后就不能使用了
                break
            }
            if j == len(newspaper)-1 { //遍历到最后都没有匹配的，直接false
                return false
            }
        }
    }
    return true
}

func handle(news, anony string) bool {
    newsChar := []byte(news)
    anonyChar := []byte(anony)
    // 因为任何顺序都可以
    sort.Slice(newsChar, func(i, j int) bool {
        return newsChar[i] < newsChar[j]
    })
    sort.Slice(anonyChar, func(i, j int) bool {
        return anonyChar[i] < anonyChar[j]
    })
    fmt.Println("排序后")
    fmt.Println(string(newsChar))
    fmt.Println(string(anonyChar))
    // 匹配
    for i := 0; i < len(newsChar); i++ {
        if newsChar[i] != anony[i] {
            return false
        }
    }
    return true

}
