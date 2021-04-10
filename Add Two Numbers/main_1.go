/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var readListNode = func(ln *ListNode) []int {
		var list []int
		var currentNode *ListNode = ln
		for {
			list = append(list, currentNode.Val)
			if currentNode.Next == nil {
				break
			}
			currentNode = currentNode.Next
		}
		return list
	}
	var summElements = func(l1 []int, l2 []int) []int {
		var list []int
		var c, i int = 0, 0
		for {
			if i > len(l1)-1 && i > len(l2)-1 {
				break
			}
			var x, y, summ int = 0, 0, 0
			if i < len(l1) {
				x = l1[i]
			}
			if i < len(l2) {
				y = l2[i]
			}
			summ = c + x + y
			c = summ / 10
			list = append(list, summ%10)
			i++
		}
		if c > 0 {
			list = append(list, c)
		}
		return list
	}
	var arrayInListNode = func(list []int) *ListNode {
		var ListNode = &ListNode{}
		var currListNode = ListNode
		for _, v := range list {
			currListNode.Next = &ListNode{v, nil}
			currListNode = currListNode.Next
		}
		return ListNode.Next
	}
	return arrayInListNode(summElements(readListNode(l1), readListNode(l2)))
}
