/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
  var buffListNode = &ListNode{}
	var currListNode = buffListNode
	var carry int = 0
	var p, q *ListNode = l1, l2
	for p != nil || q != nil {
		var x, y, summ int = 0, 0, 0
		if p != nil {
			x = p.Val
		}
		if q != nil {
			y = q.Val
		}
		summ = carry + x + y
		carry = summ / 10
		currListNode.Next = &ListNode{summ % 10, nil}
		currListNode = currListNode.Next
		if p != nil {
			p = p.Next
		}
		if q != nil {
			q = q.Next
		}
	}
	if carry > 0 {
		currListNode.Next = &ListNode{carry, nil}
	}
	return buffListNode.Next
}
