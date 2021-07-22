package main

import (
	"fmt"
	"student/functions"
	// "functions/functions"
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
	//AdvancedSort
	// result := []string{"The", "computing", "consisted", "device", "each", "earliest", "fingers", "five", "hand", "of", "of", "the", "undoubtedly"}
	// functions.AdvancedSortWordArr(result, functions.Compare)

	// fmt.Println(result)

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

	// link = functions.ListPushBackNodeI(link, 3)
	// link = functions.ListPushBackNodeI(link, 5)
	// link = functions.ListPushBackNodeI(link, 7)

	// link2 = functions.ListPushBackNodeI(link2, -2)
	// link2 = functions.ListPushBackNodeI(link2, 9)

	// functions.PrintList(functions.SortedListMerge(link2, link))

	//ListRemoveIf
	// link := &functions.List{}
	// link2 := &functions.List{}

	// fmt.Println("----normal state----")
	// functions.ListPushBackNodeL(link2, 1)
	// PrintList(link2)
	// functions.ListRemoveIf(link2, 1)
	// fmt.Println("------answer-----")
	// PrintList(link2)
	// fmt.Println()

	// fmt.Println("----normal state----")
	// functions.ListPushBackNodeL(link, 1)
	// functions.ListPushBackNodeL(link, "Hello")
	// functions.ListPushBackNodeL(link, 1)
	// functions.ListPushBackNodeL(link, "There")
	// functions.ListPushBackNodeL(link, 1)
	// functions.ListPushBackNodeL(link, 1)
	// functions.ListPushBackNodeL(link, "How")
	// functions.ListPushBackNodeL(link, 1)
	// NodeIrintList(link)
	// var str = []string{}
	// new := ""
	// check := ""
	// length := len(sep)
	// separ := false
	// count := 0
	// for index, char := range s {
	// 	if char == 'H' {
	// 		for index1, val := range s[index:(length + index)] {
	// 			if index1 == (length - 1) {
	// 				check = check + string(val)
	// 				if check == sep {
	// 					check = ""
	// 					separ = true
	// 					if len(new) > 0 {
	// 						str = append(str, new)
	// 					}
	// 					new = ""
	// 					break
	// 				} else {
	// 					check = ""
	// 					separ = false
	// 					new = new + string(char)
	// 					break
	// 				}
	// 				break
	// 			} else {
	// 				check = check + string(val)
	// 			}
	// 		}
	// 	} else if separ == true {
	// 		count++
	// 		if count == (length - 1) {
	// 			separ = false
	// 			count = 0
	// 		}
	// 	} else if index == (len(s) - 1) {
	// 		new = new + string(char)
	// 		str = append(str, new)
	// 	} else {
	// 		new = new + string(char)

	// 	}
	// }
	// return str"5")
	// node := functions.BTreeSearchItem(root, "1")
	// replacement := &TreeNode{Data: "3"}
	// root = functions.BTreeTransplanttData(root, "1")
	// functions.BTreeInsertData(root, "7")
	// functions.BTreeInsertData(root, "5")
	// fmt.Println(root, fmt.Println)
	// }

	//BTreeInsertData
	// root := &functions.TreeNode{Data: "4"}
	// functions.BTreeInsertData(root, "1")
	// functions.BTreeInsertData(root, "7")
	// functions.BTreeInsertData(root, "5")
	// fmt.Println(root.Left.Data)
	// fmt.Println(root.Data)
	// fmt.Println(root.Right.Left.Data)
	// fmt.Println(root.Right.Data)

	//BTreeApplyInorder
	// root := &functions.TreeNode{Data: "4"}
	// functions.BTreeInsertData(root, "1")
	// functions.BTreeInsertData(root, "7")
	// functions.BTreeInsertData(root, "5")
	// functions.BTreeApplyInorder(root, fmt.Println)

	//BTreeApplyPostorder
	// root := &functions.TreeNode{Data: "4"}
	// functions.BTreeInsertData(root, "1")
	// functions.BTreeInsertData(root, "7")
	// functions.BTreeInsertData(root, "5")
	// functions.BTreeApplyPreorder(root, fmt.Println)

	//BTreeSearchItem
	// root := &functions.TreeNode{Data: "4"}
	// functions.BTreeInsertData(root, "1")
	// functions.BTreeInsertData(root, "7")
	// functions.BTreeInsertData(root, "5")
	// selected := functions.BTreeSearchItem(root, "7")
	// fmt.Print("Item selected -> ")
	// if selected != nil {
	// 	fmt.Println(selected.Data)
	// } else {
	// 	fmt.Println("nil")
	// }

	// fmt.Print("Parent of selected item -> ")
	// if selected.Parent != nil {
	// 	fmt.Println(selectefunctions

	// fmt.Print("Right child of selected item -> ")
	// if selected.Right != nil {
	// 	fmt.Println(selected.Right.Data)
	// } else {
	// 	fmt.Println("nil")
	// }

	//BTreeTransplant
	// root := &functions.TreeNode{Data: "4"}
	// functions.BTreeInsertData(root, "1")
	// functions.BTreeInsertData(root, "7")
	// functions.BTreeInsertData(root, "5")
	// node := functions.BTreeSearchItem(root, "1")
	// replacement := &functions.TreeNode{Data: "3"}
	// root = functions.BTreeTransplant(root, node, replacement)
	// functions.BTreeApplyInorder(root, fmt.Println)

	//BTreeApplyByLevel
	// root := &functions.TreeNode{Data: "4"}
	// functions.BTreeInsertData(root, "1")
	// functions.BTreeInsertData(root, "7")
	// functions.BTreeInsertData(root, "5")
	// functions.BTreeApplyByLevel(root, fmt.Println)

	//BTreeDeleteNode
	// root := &functions.TreeNode{Data: "4"}
	// functions.BTreeInsertData(root, "1")
	// functions.BTreeInsertData(root, "7")
	// functions.BTreeInsertData(root, "5")
	// node := functions.BTrfunctionsDeleteNode(root, node)
	// fmt.Println("After delete:")
	// functions.BTreeApplyInorder(root, fmt.Println)

	//Split
	// 	s := "HAHAHelloHAHAHAHAhowHAareHAyou?HAHA"
	// 	fmt.Println(functions.Split(s, "HA"))
}
