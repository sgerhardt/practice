package linked_lists

import (
	"math"
	"strconv"
)

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	l1Val, l2Val := 0, 0
	multiplier := 0
	for n := l1; n != nil; n = n.Next {
		l1Val += int(math.Pow(float64(10), float64(multiplier))) * n.Val
		multiplier += 1
	}
	multiplier = 0
	for n := l2; n != nil; n = n.Next {
		l2Val += int(math.Pow(float64(10), float64(multiplier))) * n.Val
		multiplier += 1
	}

	sum := l1Val + l2Val
	sVal := strconv.Itoa(sum)

	resultNodes := make([]*ListNode, len(sVal))

	for n := 0; n < len(resultNodes); n++ {
		resultNodes[n] = &ListNode{}
	}
	nodeIndex := 0
	for i := len(sVal); i > 0; i-- {
		val, err := strconv.Atoi(string(sVal[i-1]))
		if err != nil {
			panic(err)
		}
		resultNodes[nodeIndex].Val = val
		nodeIndex++
		if nodeIndex < len(resultNodes) {
			resultNodes[nodeIndex-1].Next = resultNodes[nodeIndex]
		}
	}
	return resultNodes[0]
}

func addTwoNumbersOptimal(l1 *ListNode, l2 *ListNode) *ListNode {
	c1 := l1
	c2 := l2

	sentinel := &ListNode{}
	d := sentinel

	sum := 0

	for c1 != nil || c2 != nil {
		sum /= 10 // note, divide by 10 s
		if c1 != nil {
			sum += c1.Val
			c1 = c1.Next
		}
		if c2 != nil {
			sum += c2.Val
			c2 = c2.Next
		}
		d.Next = &ListNode{
			Val: sum % 10,
			// mod 10 in case the value is greater than or equal to 10,
			// but don't modify the sum variable as we need that on the iteration (think of a carry value)
		}
		d = d.Next
	}

	if sum/10 == 1 {
		// add another node for the carry value...
		d.Next = &ListNode{
			Val: 1,
		}
	}
	return sentinel.Next
}
