package od

import (
    "bufio"
    "fmt"
    "os"
    "sort"
    "testing"
)

/**

每人每次面试的时长不等，并已经安排给定，
用(S1,E1)、(S2,E2)、(Sj,Ej)...(Si < Ei，均为非负整数)表示每场面试的开始和结束时间。

即一名面试官同时只能面试一名应试者，一名面试官完成一次面试后可以立即进行下一场面试，

且每个面试官的面试人次不超过m。

请你计算至少需要多少名面试官。
2
4
1 2
3 5
4 7
6 8

结果是2

*/
func TestGakl(t *testing.T) {

    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    scanner.Text()
    var a [][]int
    var m , n int
    fmt.Scan(&m,&n)

    for i := 0; i < n; i++ {
        a[i]=make([]int, 2)
    }
    for i := 0; i < n ; i++ {
        fmt.Scan(&a[i][0])
        fmt.Scan(&a[i][1])
    }

    fmt.Println(a)
    sort.Ints(a[0]) // 按照开始时间排序
    fmt.Println(a)
    //Arrays.sort(arr,Comparator.comparing(c1->c1[0]));
    //Deque<int[]> deque = new ArrayDeque<>();
    //deque.offer(new int[]{arr[0][1],1});
    //// 队列存需要的面试官(结束时间m，面试官编号)
    //for (int i = 1; i < n; i++) {
    //    boolean flag = false;
    //    for (int[] p : deque) {
    //        if (p[0]<= arr[i][0] && p[1]<m){
    //            p[1]++;
    //            p[0] = arr[i][1];
    //            flag = true;
    //            break;
    //        }
    //    }
    //    if (!flag){
    //        deque.offer(new int[]{arr[i][1],1});
    //    }
    //}
    //System.out.println(deque.size());

}

