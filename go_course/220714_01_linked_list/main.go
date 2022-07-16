package main

import (
	"fmt"
)

func main() {

	n0 := CreateNode(0)
	n1 := CreateNode(1)
	n2 := CreateNode(2)
	n3 := CreateNode(3)
	n4 := CreateNode(4)

	// n0 is a pointer to a Node
	// In C we would use n0->next
	// but in Go there is not differentiation
	n0.next = n1
	n1.next = n2
	n2.next = n3
	n3.next = n4

	printLinkedList(n0)
	// 0 - 1 - 2 - 3 - 4 -
}

type Node struct {
	data int
	next *Node
}

// This function is returning an address a believe that this is a case
// of dynamic allocation in Go
// this is another comment
// yet another comment
func CreateNode(data int) *Node {
	aNode := Node{
		data: data,
		next: nil,
	}

	return &aNode
}

// about folding comments in vim:
// https://vi.stackexchange.com/questions/3512/how-to-fold-comments
// set foldmethod=expr foldexpr=getline(v:lnum)=~'^\\s*'.&commentstring[0]
// utocmd FileType go      setlocal foldmethod=expr foldexpr=getline(v:lnum)=~'^\\s*//'

// about folding
// https://vim.fandom.com/wiki/Folding
// setlocal foldmethod=syntax
// this one is good for folding code
// but it is not good for folding comments

// Vim fold on tab
// https://vi.stackexchange.com/questions/2176/how-to-write-a-fold-expr
// set foldexpr=getline(v:lnum)[0]==\"\\t\"
// set foldexpr=getline(v:lnum-1)=~'^\\s*$'&&getline(v:lnum)=~'\\S'?'>1':1

func printLinkedList(head *Node) {
	ptrNode := head
	fmt.Printf("%v - ", ptrNode.data)
	for ptrNode.next != nil {
		ptrNode = ptrNode.next
		fmt.Printf("%v - ", ptrNode.data)
	}
}

// https://learnvim.irian.to/basics/fold
// set foldmethod=indent
// this one is not good

// I should take a look at the vim-go plugin for vim
// https://github.com/fatih/vim-go
// there I can find go-code folding
// and better syntax highlighting
//
