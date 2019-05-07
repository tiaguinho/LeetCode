/**
 * Link to the problem
 * https://leetcode.com/problems/evaluate-reverse-polish-notation/
 */
package leetcode

import "strconv"

// Node
type Node struct {
	Val  int
	Prev *Node
}

// Stack struct
type Stack struct {
	Tail *Node
}

func (s *Stack) Push(n *Node) {
	n.Prev = s.Tail
	s.Tail = n
}

func (s *Stack) Pop() *Node {
	tmp := s.Tail
	s.Tail = tmp.Prev
	return tmp
}

func evalRPN(tokens []string) int {
	stack := Stack{}

	var result int
	for _, token := range tokens {
		if r, err := strconv.ParseInt(token, 10, 64); err == nil {
			result = int(r)
		} else {
			switch token {
			case "+":
				val2 := stack.Pop()
				val1 := stack.Pop()
				result = val1.Val + val2.Val
				break

			case "-":
				val2 := stack.Pop()
				val1 := stack.Pop()
				result = val1.Val - val2.Val
				break

			case "*":
				val2 := stack.Pop()
				val1 := stack.Pop()
				result = val1.Val * val2.Val
				break

			case "/":
				val2 := stack.Pop()
				val1 := stack.Pop()
				result = val1.Val / val2.Val
				break
			}
		}

		stack.Push(&Node{Val: result})
	}

	return stack.Pop().Val
}
