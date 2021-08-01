package main

import (
	"fmt"
	"strings"
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
	//Atoi
	// fmt.Println(functions.Atoi("12345"))
	// fmt.Println(functions.Atoi("0000000012345"))
	// fmt.Println(functions.Atoi("012 345"))
	// fmt.Println(functions.Atoi("Hello World!"))
	// fmt.Println(functions.Atoi("+1234"))
	// fmt.Println(functions.Atoi("-1234"))
	// fmt.Println(functions.Atoi("++1234"))
	// fmt.Println(functions.Atoi("--1234"))
	// fmt.Println(functions.Atoi("-9223372036854775808"))

	//AdvancedSort
	// result1 := []string{"The", "earliest", "computing", "device", "undoubtedly", "consisted", "of", "the", "five", "fingers", "of", "each", "hand"}
	// functions.AdvancedSortWordArr(result1, functions.Compare)
	// result2 := []string{"The", "word", "digital", "comesfrom", "\"digits\"", "or", "fingers"}
	// functions.AdvancedSortWordArr(result2, functions.Compare)
	// result3 := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	// functions.AdvancedSortWordArr(result3, functions.Compare)
	// result4 := []string{"The", "computing", "consisted", "device", "each", "earliest", "fingers", "five", "hand", "of", "of", "the", "undoubtedly"}
	// functions.AdvancedSortWordArr(result4, functions.Compare)
	// result5 := []string{"\"digits\"", "The", "comesfrom", "digital", "fingers", "or", "word"}
	// functions.AdvancedSortWordArr(result5, functions.Compare)
	// result6 := []string{"1", "2", "3", "A", "B", "C", "a", "b", "c"}
	// functions.AdvancedSortWordArr(result6, functions.Compare)

	// fmt.Println(result1)
	// fmt.Println(result2)
	// fmt.Println(result3)
	// fmt.Println(result4)
	// fmt.Println(result5)
	// fmt.Println(result6)

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

	// fmt.Print("Right child of selected item -> ")//Atoi
	// fmt.Println(functions.Atoi("12345"))
	// fmt.Println(functions.Atoi("0000000012345"))
	// fmt.Println(functions.Atoi("012 345"))
	// fmt.Println(functions.Atoi("Hello World!"))
	// fmt.Println(functions.Atoi("+1234"))
	// fmt.Println(functions.Atoi("-1234"))
	// fmt.Println(functions.Atoi("++1234"))
	// fmt.Println(functions.Atoi("--1234"))
	// fmt.Println(functions.Atoi("-9223372036854775808"))

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
	s := "which|=choumi=|itself|=choumi=|used|=choumi=|cards|=choumi=|and|=choumi=|a|=choumi=|central|=choumi=|computing|=choumi=|unit.|=choumi=|When|=choumi=|the|=choumi=|machine|=choumi=|was|=choumi=|finished,"
	fmt.Println(functions.Split(s, "|=choumi=|"))
	a := "!==!which!==!was!==!making!==!all!==!kinds!==!of!==!punched!==!card!==!equipment!==!and!==!was!==!also!==!in!==!the!==!calculator!==!business[10]!==!to!==!develop!==!his!==!giant!==!programmable!==!calculator,"
	fmt.Println(functions.Split(a, "!==!"))
	fmt.Println(strings.Split(s, "|=choumi=|"))
	fmt.Println(functions.Split("HAHAHelloHAHowAreYouHAHAHAHere", "HA"))
	fmt.Println("---------------------------------------------------------------")
	fmt.Println(strings.Split("HAHAHelloHAHowAreYouHAHAHAHere", "HA"))
	fmt.Println(functions.Split("aaaaa", "aa"))
	fmt.Println(strings.Split("aaaaa", "aa"))

	//AtoiBase
	// fmt.Println(functions.AtoiBase("4506C", "0123456789ABCDEF"))
	// fmt.Println(functions.AtoiBase("0001", "01"))
	// fmt.Println(functions.AtoiBase("00", "01"))
	// fmt.Println(functions.AtoiBase("saDt!I!sI", "CHOUMIisDAcat!"))
	// fmt.Println(functions.AtoiBase("AAho?Ao", "WhoAmI?"))
	// fmt.Println(functions.AtoiBase("thisinputshouldnotmatter", "abca"))

	//SplitWhiteSpaces
	// fmt.Println(functions.SplitWhiteSpaces("The earliest foundations of what would become computer science predate the invention of the modern digital computer"))
	// fmt.Println(functions.SplitWhiteSpaces("Machines for calculating fixed numerical tasks such as the abacus have existed since antiquity"))
	// fmt.Println(functions.SplitWhiteSpaces("aiding in computations such as multiplication and division ."))
	// fmt.Println(functions.SplitWhiteSpaces("Algorithms for performing computations have existed since antiquity, even before the development of sophisticated computing equipment."))
	// fmt.Println(functions.SplitWhiteSpaces("Wilhelm Schickard designed and constructed the first working mechanical calculator in 1623.[4]"))
	// fmt.Println(functions.SplitWhiteSpaces("In 1673, Gottfried Leibniz demonstrated a digital mechanical calculator,"))

	//Split
	// s := " !==!which!==!was!==!making!==!all!==!kinds!==!of!==!punched!==!card!==!equipment!==!and!==!was!==!also!==!in!==!the!==!calculator!==!business[10]!==!to!==!develop!==!his!==!giant!==!programmable!==!calculator,"
	// fmt.Println(functions.Split(s, "!==!"))
	// st := " AFJCharlesAFJBabbageAFJstartedAFJtheAFJdesignAFJofAFJtheAFJfirstAFJautomaticAFJmechanicalAFJcalculator,"
	// fmt.Println(functions.Split(st, "AFJ"))
	// sb := " <<==123==>>In<<==123==>>1820,<<==123==>>Thomas<<==123==>>de<<==123==>>Colmar<<==123==>>launched<<==123==>>the<<==123==>>mechanical<<==123==>>calculator<<==123==>>industry[note<<==123==>>1]<<==123==>>when<<==123==>>he<<==123==>>released<<==123==>>his<<==123==>>simplified<<==123==>>arithmometer,"
	// fmt.Println(functions.Split(sb, "<<==123==>>"))

	//ConvertBase
	// result1 := functions.ConvertBase("4506C", "0123456789ABCDEF", "choumi")
	// fmt.Println(result1)
	// result2 := functions.ConvertBase("babcbaaaaabcb", "abc", "0123456789ABCDEF")
	// fmt.Println(result2)
	// result3 := functions.ConvertBase("256850", "0123456789", "01")
	// fmt.Println(result3)
	// result4 := functions.ConvertBase("uuhoumo", "choumi", "Zone01")
	// fmt.Println(result4)
	// result5 := functions.ConvertBase("683241", "0123456789", "0123456789")
	// fmt.Println(result5)

	//PrinrNbrBase

	// functions.PrintNbrBase(-125, "01")
	// z01.PrintRune('\n')
	// functions.PrintNbrBase(125, "0123456789ABCDEF")
	// z01.PrintRune('\n')
	// functions.PrintNbrBase(919617, "01")
	// z01.PrintRune('\n')
	// functions.PrintNbrBase(753639, "CHOUMIisDAcat!")
	// z01.PrintRune('\n')
	// functions.PrintNbrBase(-174336, "CHOUMIisDAcat!")
	// z01.PrintRune('\n')
	// functions.PrintNbrBase(-661165, "1")
	// z01.PrintRune('\n')
	// functions.PrintNbrBase(-661165, "Zone01Zone01")
	// z01.PrintRune('\n')
	// functions.PrintNbrBase(-661165, "")
	// z01.PrintRune('\n')

	//PrintCombn
	// functions.PrintCombN(1)
	// functions.PrintCombN(3)
	// functions.PrintCombN(9)
}
