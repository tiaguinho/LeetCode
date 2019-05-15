/**
 * Link to the problem
 * https://leetcode.com/problems/valid-parentheses
 */
package leetcode

type Node struct {
	Value rune
	Next  *Node
}

type Stack struct {
	Head *Node
}

func (s *Stack) Push(node *Node) {
	node.Next = s.Head
	s.Head = node
}

func (s *Stack) Pop() *Node {
	node := s.Head
	s.Head = node.Next
	return node
}

var chars = map[rune]rune{'(': ')', '[': ']', '{': '}'}

func isValid(s string) bool {
	if s == "" {
		return true
	}

	if len(s)%2 != 0 {
		return false
	}

	stack := Stack{}
	for _, l := range s {
		if chars[l] != 0 {
			stack.Push(&Node{Value: l})
		} else {
			if stack.Head == nil {
				return false
			}

			node := stack.Pop()
			if chars[node.Value] != l {
				return false
			}
		}
	}

	return stack.Head == nil
}
