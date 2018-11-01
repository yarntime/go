/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode)  {
    if head == nil {
        return
    }
    firstHead, firstCur := head, head
    secondHead, secondCur := head.Next, head.Next
    for secondCur != nil && secondCur.Next != nil {
        firstCur.Next = firstCur.Next.Next
        firstCur = firstCur.Next
        secondCur.Next = secondCur.Next.Next
        secondCur = secondCur.Next
    }
    firstCur.Next = nil
    
    if secondHead == nil {
        return
    }
    secondCur = secondHead.Next
    secondHead.Next = nil
    for secondCur != nil {
        t := secondCur.Next
        secondCur.Next = secondHead
        secondHead = secondCur
        secondCur = t
    }
    
    firstCur = firstHead.Next
    secondCur = secondHead
    
    newHead := firstHead
    newHead.Next = nil
    newCur := newHead
    
    for {
        if secondCur != nil {
            t := secondCur.Next
            secondCur.Next = newCur.Next
            newCur.Next = secondCur
            newCur = newCur.Next
            secondCur = t
        }
        if firstCur != nil {
            t := firstCur.Next
            firstCur.Next = newCur.Next
            newCur.Next = firstCur
            newCur = newCur.Next
            firstCur = t
        }
        if secondCur == nil && firstCur == nil {
            break
        }
    }
    head = newHead
}