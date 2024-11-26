package first

import (
    "fmt"
    "testing"
)

type Node struct { // 每个节点有一个优先级
    value    interface{}
    priority int // 优先级
    next     *Node
}

type PriorityQueue struct {
    head *Node
    size int // 优先队列长度
}

func (pq *PriorityQueue) Insert(value interface{}, priority int) {
    newNode := &Node{value, priority, nil}

    if pq.head == nil {
        pq.head = newNode
    } else if pq.head.priority < priority { // 优先级比头节点高，插入头部
        newNode.next = pq.head
        pq.head = newNode
    } else { //  遍历整个队列，找到第一个优先级大于等于新节点的位置，并将新节点插入到此位置之前
        current := pq.head

        for current.next != nil && current.next.priority <= priority {
            current = current.next
        }

        newNode.next = current.next
        current.next = newNode
    }

    pq.size++
}

func (pq *PriorityQueue) Pop() interface{} { // 弹出优先级最高的队列
    if pq.head == nil {
        return nil
    }

    poppedNode := pq.head
    pq.head = poppedNode.next
    poppedNode.next = nil
    pq.size--

    return poppedNode.value
}

func (pq *PriorityQueue) Size() int {
    return pq.size
}

func TestRgsa(t *testing.T) {
    pq := PriorityQueue{}

    pq.Insert("task 1", 3)
    pq.Insert("task 2", 1)
    pq.Insert("task 3", 2)

    for pq.Size() > 0 {
        fmt.Println(pq.Pop())
    }
}
