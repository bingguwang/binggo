package od

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "testing"
)

func TestQuest(t *testing.T) {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    text := strings.ToLower(scanner.Text())
    scanner.Scan()
    char := strings.ToLower(scanner.Text())

    var count int
    for i := 0; i < len(text); i++ {
        if text[i] == char[0] {
            count++
        }
    }
    fmt.Println(count)
}
