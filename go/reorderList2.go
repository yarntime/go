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
    firstCur := head
    secondCur := head.Next
    for secondCur != nil && secondCur.Next != nil {
        firstCur = firstCur.Next
        secondCur = secondCur.Next.Next
    }
   
    secondHead := firstCur.Next
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
    
    for firstCur, secondCur = head, secondHead; firstCur != nil; {
        t := firstCur.Next
        firstCur.Next = secondCur
        firstCur = secondCur
        secondCur = t
    }
}