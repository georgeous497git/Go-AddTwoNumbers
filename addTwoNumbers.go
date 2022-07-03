package addTwoNumbers

import (
	"math"
)

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var complete *ListNode
	var current *ListNode
	var tenths = 0.0

	for {
		if l1 == nil && l2 == nil {

			newValue := int(tenths)
			if newValue > 0 {
				aux := &ListNode{
					newValue, nil,
				}
				addNode(complete, current, aux)
			}
			break
		}

		var num1, num2 int
		if l1 != nil {
			num1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			num2 = l2.Val
			l2 = l2.Next
		}

		add := (num1 + num2) + int(tenths)
		sum := math.Mod(float64(add), 10)
		sum2 := int(sum)
		tenths = math.Mod(float64(add), 100) / 10

		aux := &ListNode{
			sum2, nil,
		}

		complete, current = addNode(complete, current, aux)
	}

	return complete
}

func addNode(complete *ListNode, current *ListNode, aux *ListNode) (*ListNode, *ListNode) {
	if complete == nil {
		complete = aux
		current = complete
	} else {
		current.Next = aux
		current = current.Next
	}

	return complete, current
}
