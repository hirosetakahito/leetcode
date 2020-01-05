/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

package main
import "fmt"


type ListNode struct {
    Val int
    Next *ListNode
}


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var result *ListNode = nil
    var previous *ListNode = nil

    ln1 := l1
    ln2 := l2

    v1 := 0
    v2 := 0
    sum := 0
    advance := 0

    for ln1 != nil || ln2 != nil || advance != 0 {
        if ln1 != nil {
            v1 = ln1.Val
        } else {
            v1 = 0
        }
        if ln2 != nil {
            v2 = ln2.Val
        } else {
            v2 = 0
        }
        sum = v1 + v2 + advance
        current := new(ListNode)
        current.Val = sum % 10
        current.Next = nil

        if result == nil {
            previous = current
            result = current
        } else {
            previous.Next = current
            previous = current
        }

        advance = sum / 10
        if ln1 != nil {
            ln1 = ln1.Next
        }
        if ln2 != nil {
            ln2 = ln2.Next
        }
    }

    return result
}


func main() {
    var l1, l2 ListNode
    l1 = ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{}}}}
    l2 = ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: &ListNode{}}}}
    result := addTwoNumbers(&l1, &l2) 
    fmt.Println(result.Val, result.Next.Val, result.Next.Next.Val)
}
