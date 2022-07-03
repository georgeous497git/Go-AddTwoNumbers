package addTwoNumbers

import (
	"strconv"
	"testing"
)

func TestAddTwoNumbers2(t *testing.T) {
	l1 := []int{9, 9, 9, 9, 9, 9, 9}
	l2 := []int{9, 9, 9, 9}
	l3 := []int{8, 9, 9, 9, 0, 0, 0, 1}

	node1 := buildNode(l1)
	node2 := buildNode(l2)
	wantList := buildNode(l3)
	newList := AddTwoNumbers(node1, node2)

	wantNumbers := getNumbers(wantList)
	newNumbers := getNumbers(newList)

	if wantNumbers != newNumbers {
		t.Errorf("The wanted list %s is different of %s", wantNumbers, newNumbers)
	}
}

func TestAddTwoNumbers(t *testing.T) {
	l1 := []int{2, 4, 3}
	l2 := []int{5, 6, 4}
	l3 := []int{7, 0, 8}

	node1 := buildNode(l1)
	node2 := buildNode(l2)
	wantList := buildNode(l3)
	newList := AddTwoNumbers(node1, node2)

	wantNumbers := getNumbers(wantList)
	newNumbers := getNumbers(newList)

	if wantNumbers != newNumbers {
		t.Errorf("The list are different %s %s", wantNumbers, newNumbers)
	}
}

func getNumbers(node *ListNode) string {
	var numbers string
	for {
		if node == nil {
			break
		}
		numbers += strconv.Itoa(node.Val)
		node = node.Next
	}
	return numbers
}

func buildNode(list []int) *ListNode {

	var listNode *ListNode
	var temp *ListNode

	size := len(list)
	for i := size - 1; i >= 0; i-- {

		value := list[i]

		if listNode == nil {
			listNode = &ListNode{
				value, nil,
			}
		} else {
			temp = &ListNode{
				value, listNode,
			}
			listNode = temp
		}
	}

	return listNode
}
