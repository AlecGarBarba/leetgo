package invertsingly

import (
	"testing"

	sll "github.com/alecGarBarba/leetGo/linkedlist/singly_linked_list"
	"github.com/stretchr/testify/assert"
)

func TestNilNode(t *testing.T) {
	var node *sll.SinglyListNode

	invertedNode := invertSinglyLinkedList(node)

	assert.Nil(t, invertedNode)
}

func TestSingleNode(t *testing.T) {
	node := &sll.SinglyListNode{Val: 1}

	invertedNode := invertSinglyLinkedList(node)

	assert.Equal(t, 1, invertedNode.Val, "The value of the node should be 1")
}

func TestTwoNodes(t *testing.T) {
	node := &sll.SinglyListNode{Val: 1, Next: &sll.SinglyListNode{Val: 2}}

	invertedNode := invertSinglyLinkedList(node)

	assert.Equal(t, 2, invertedNode.Val, "The value of the node should be 2")
	assert.Equal(t, 1, invertedNode.Next.Val, "The value of the next node should be 1")
}

func TestFourNodes(t *testing.T) {
	/* LI structure:
	1 -> 2 -> 3 -> 4
	*/
	node := &sll.SinglyListNode{Val: 1, Next: &sll.SinglyListNode{Val: 2, Next: &sll.SinglyListNode{Val: 3, Next: &sll.SinglyListNode{Val: 4}}}}
	invertedNode := invertSinglyLinkedList(node)

	/* LI structure after inversion:
	4 -> 3 -> 2 -> 1
	*/
	assert.Equal(t, 4, invertedNode.Val, "The value of the node should be 4")
	assert.Equal(t, 3, invertedNode.Next.Val, "The value of the next node should be 3")
	assert.Equal(t, 2, invertedNode.Next.Next.Val, "The value of the next node should be 2")
	assert.Equal(t, 1, invertedNode.Next.Next.Next.Val, "The value of the next node should be 1")
}
