package main

import (
	"fmt"
	"student/functions"
)

func PrintList(l *functions.List) {
	it := l.Head
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}

	fmt.Print(nil, "\n")
}

func main() {
	// nbits := functions.ActiveBits(7)
	// fmt.Println(nbits)

	//SortList
	// var link *functions.NodeI

	// link = functions.ListPushBack(link, 1)
	// link = functions.ListPushBack(link, 4)
	// link = functions.ListPushBack(link, 9)

	// functions.PrintList(link)

	// link = functions.SortListInsert(link, -2)
	// link = functions.SortListInsert(link, 2)
	// functions.PrintList(link)

	//Merge
	// var link *functions.NodeI
	// var link2 *functions.NodeI

	// link = functions.ListPushBack(link, 3)
	// link = functions.ListPushBack(link, 5)
	// link = functions.ListPushBack(link, 7)

	// link2 = functions.ListPushBack(link2, -2)
	// link2 = functions.ListPushBack(link2, 9)

	// functions.PrintList(functions.SortedListMerge(link2, link))

	//ListRemoveIf
	link := &functions.List{}
	link2 := &functions.List{}

	fmt.Println("----normal state----")
	functions.ListPushBackNodeL(link2, 1)
	PrintList(link2)
	functions.ListRemoveIf(link2, 1)
	fmt.Println("------answer-----")
	PrintList(link2)
	fmt.Println()

	fmt.Println("----normal state----")
	functions.ListPushBackNodeL(link, 1)
	functions.ListPushBackNodeL(link, "Hello")
	functions.ListPushBackNodeL(link, 1)
	functions.ListPushBackNodeL(link, "There")
	functions.ListPushBackNodeL(link, 1)
	functions.ListPushBackNodeL(link, 1)
	functions.ListPushBackNodeL(link, "How")
	functions.ListPushBackNodeL(link, 1)
	functions.ListPushBackNodeL(link, "are")
	functions.ListPushBackNodeL(link, "you")
	functions.ListPushBackNodeL(link, 1)
	PrintList(link)

	functions.ListRemoveIf(link, 1)
	fmt.Println("------answer-----")
	PrintList(link)

}
