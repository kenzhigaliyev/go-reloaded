package main

import "student/functions"

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
	var link *functions.NodeI
	var link2 *functions.NodeI

	link = functions.ListPushBack(link, 3)
	link = functions.ListPushBack(link, 5)
	link = functions.ListPushBack(link, 7)

	link2 = functions.ListPushBack(link2, -2)
	link2 = functions.ListPushBack(link2, 9)

	functions.PrintList(functions.SortedListMerge(link2, link))

}
