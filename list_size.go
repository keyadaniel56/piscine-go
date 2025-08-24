package main

import "fmt"

type ListNode struct {
	Data interface{}
	Next *ListNode
}

func ListSize(head *ListNode) int {
	count := 0
	current := head
	for current != nil {
		count++
		current = current.Next
	}
	return count
}
