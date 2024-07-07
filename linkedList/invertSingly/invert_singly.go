package invertsingly

import sll "github.com/alecGarBarba/leetGo/linkedlist/singly_linked_list"

func invertSinglyLinkedList(head *sll.SinglyListNode) *sll.SinglyListNode {

	var prev *sll.SinglyListNode = nil
	curr := head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}
