package mergetwosortedlists

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) String() string {
	if n == nil {
		return ""
	}
	if n.Next == nil {
		return fmt.Sprintf("%d", n.Val)
	}
	return fmt.Sprintf("%d, %s", n.Val, n.Next.String())
}

func NewList(inputList []int) *ListNode {
	var curentNode *ListNode
	for i := len(inputList); i > 0; i-- {
		newNode := &ListNode{
			Val:  inputList[i-1],
			Next: curentNode,
		}
		curentNode = newNode
	}
	return curentNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	var nodeResult, greaterNode = getMinorAndGreaterNode(list1, list2)
	nodeResult.Next = mergeTwoLists(nodeResult.Next, greaterNode)
	return &nodeResult
}

func getMinorAndGreaterNode(list1 *ListNode, list2 *ListNode) (minor ListNode, greater *ListNode) {
	if list1.Val <= list2.Val {
		return *list1, list2
	}
	return *list2, list1
}
