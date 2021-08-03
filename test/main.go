package main

import (
	"fmt"
	"strings"
	"student"
)

func main() {
	//Atoi
	// fmt.Println(student.Atoi("12345"))
	// fmt.Println(student.Atoi("0000000012345"))
	// fmt.Println(student.Atoi("012 345"))
	// fmt.Println(student.Atoi("Hello World!"))
	// fmt.Println(student.Atoi("+1234"))
	// fmt.Println(student.Atoi("-1234"))
	// fmt.Println(student.Atoi("++1234"))
	// fmt.Println(student.Atoi("--1234"))
	// fmt.Println(student.Atoi("-9223372036854775808"))

	//RecursivePower
	// fmt.Println(student.RecursivePower(-7, -2))
	// fmt.Println(student.RecursivePower(-8, -7))
	// fmt.Println(student.RecursivePower(4, 8))
	// fmt.Println(student.RecursivePower(1, 3))
	// fmt.Println(student.RecursivePower(-1, 1))
	// fmt.Println(student.RecursivePower(-6, 5))

	//PrintCombn
	// student.PrintCombN(1)
	// student.PrintCombN(2)
	// student.PrintCombN(3)
	// student.PrintCombN(4)
	// student.PrintCombN(5)
	// student.PrintCombN(6)
	// student.PrintCombN(7)
	// student.PrintCombN(8)
	// student.PrintCombN(9)

	//PrinrNbrBase
	// student.PrintNbrBase(-125, "01")
	// z01.PrintRune('\n')
	// student.PrintNbrBase(125, "0123456789ABCDEF")
	// z01.PrintRune('\n')
	// student.PrintNbrBase(919617, "01")
	// z01.PrintRune('\n')
	// student.PrintNbrBase(753639, "CHOUMIisDAcat!")
	// z01.PrintRune('\n')
	// student.PrintNbrBase(-174336, "CHOUMIisDAcat!")
	// z01.PrintRune('\n')
	// student.PrintNbrBase(-661165, "1")
	// z01.PrintRune('\n')
	// student.PrintNbrBase(-661165, "Zone01Zone01")
	// z01.PrintRune('\n')
	// student.PrintNbrBase(-661165, "")
	// z01.PrintRune('\n')

	//AtoiBase
	// fmt.Println(student.AtoiBase("bcbbbbaab", "abc"))
	// fmt.Println(student.AtoiBase("uoi", "choumi"))
	// fmt.Println(student.AtoiBase("bbbbbab", "-ab"))
	// fmt.Println(student.AtoiBase("125", "0123456789"))
	// fmt.Println(student.AtoiBase("AAho?Ao", "WhoAmI?"))
	// fmt.Println(student.AtoiBase("4506C", "0123456789ABCDEF"))
	// fmt.Println(student.AtoiBase("0001", "01"))
	// fmt.Println(student.AtoiBase("00", "01"))
	// fmt.Println(student.AtoiBase("saDt!I!sI", "CHOUMIisDAcat!"))
	// fmt.Println(student.AtoiBase("AAho?Ao", "WhoAmI?"))
	// fmt.Println(student.AtoiBase("thisinputshouldnotmatter", "abca"))

	//SplitWhiteSpaces
	// fmt.Println(student.SplitWhiteSpaces("The earliest foundations of what would become computer science predate the invention of the modern digital computer"))
	// fmt.Println(student.SplitWhiteSpaces("Machines for calculating fixed numerical tasks such as the abacus have existed since antiquity"))
	// fmt.Println(student.SplitWhiteSpaces("aiding in computations such as multiplication and division ."))
	// fmt.Println(student.SplitWhiteSpaces("Algorithms for performing computations have existed since antiquity, even before the development of sophisticated computing equipment."))
	// fmt.Println(student.SplitWhiteSpaces("Wilhelm Schickard designed and constructed the first working mechanical calculator in 1623.[4]"))
	// fmt.Println(student.SplitWhiteSpaces("In 1673, Gottfried Leibniz demonstrated a digital mechanical calculator,"))

	//Split
	// s := "which|=choumi=|itself|=choumi=|used|=choumi=|cards|=choumi=|and|=choumi=|a|=choumi=|central|=choumi=|computing|=choumi=|unit.|=choumi=|When|=choumi=|the|=choumi=|machine|=choumi=|was|=choumi=|finished,"
	// a := "!==!which!==!was!==!making!==!all!==!kinds!==!of!==!punched!==!card!==!equipment!==!and!==!was!==!also!==!in!==!the!==!calculator!==!business[10]!==!to!==!develop!==!his!==!giant!==!programmable!==!calculator,"
	// fmt.Println(student.Split(a, "!==!"))
	// fmt.Println(student.Split(s, "|=choumi=|"))
	// fmt.Println(student.Split("HAHAHelloHAHowAreYouHAHAHAHere", "HA"))
	// fmt.Println(student.Split("aaaaa", "aa"))
	// sa := " !==!which!==!was!==!making!==!all!==!kinds!==!of!==!punched!==!card!==!equipment!==!and!==!was!==!also!==!in!==!the!==!calculator!==!business[10]!==!to!==!develop!==!his!==!giant!==!programmable!==!calculator,"
	// fmt.Println(student.Split(sa, "!==!"))
	// st := " AFJCharlesAFJBabbageAFJstartedAFJtheAFJdesignAFJofAFJtheAFJfirstAFJautomaticAFJmechanicalAFJcalculator,"
	// fmt.Println(student.Split(st, "AFJ"))
	// sb := " <<==123==>>In<<==123==>>1820,<<==123==>>Thomas<<==123==>>de<<==123==>>Colmar<<==123==>>launched<<==123==>>the<<==123==>>mechanical<<==123==>>calculator<<==123==>>industry[note<<==123==>>1]<<==123==>>when<<==123==>>he<<==123==>>released<<==123==>>his<<==123==>>simplified<<==123==>>arithmometer,"
	// fmt.Println(student.Split(sb, "<<==123==>>"))
	fmt.Println(len(strings.Split("aa", "")))
	fmt.Println(len(student.Split("aa", "")))

	//ConvertBase
	// result1 := student.ConvertBase("4506C", "0123456789ABCDEF", "choumi")
	// fmt.Println(result1)
	// result2 := student.ConvertBase("babcbaaaaabcb", "abc", "0123456789ABCDEF")
	// fmt.Println(result2)
	// result3 := student.ConvertBase("256850", "0123456789", "01")
	// fmt.Println(result3)
	// result4 := student.ConvertBase("uuhoumo", "choumi", "Zone01")
	// fmt.Println(result4)
	// result5 := student.ConvertBase("683241", "0123456789", "0123456789")
	// fmt.Println(result5)

	//AdvancedSort
	// result1 := []string{"The", "earliest", "computing", "device", "undoubtedly", "consisted", "of", "the", "five", "fingers", "of", "each", "hand"}
	// student.AdvancedSortWordArr(result1, student.Compare)
	// fmt.Println(result1)
	// result2 := []string{"The", "word", "digital", "comesfrom", "\"digits\"", "or", "fingers"}
	// student.AdvancedSortWordArr(result2, student.Compare)
	// fmt.Println(result2)
	// result3 := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	// student.AdvancedSortWordArr(result3, student.Compare)
	// fmt.Println(result3)
	// result4 := []string{"The", "computing", "consisted", "device", "each", "earliest", "fingers", "five", "hand", "of", "of", "the", "undoubtedly"}
	// student.AdvancedSortWordArr(result4, student.Compare1)
	// fmt.Println(result4)
	// result5 := []string{"\"digits\"", "The", "comesfrom", "digital", "fingers", "or", "word"}
	// student.AdvancedSortWordArr(result5, student.Compare1)
	// fmt.Println(result5)
	// result6 := []string{"1", "2", "3", "A", "B", "C", "a", "b", "c"}
	// student.AdvancedSortWordArr(result6, student.Compare1)
	// fmt.Println(result6)

	//ActiveBits
	// fmt.Println(student.ActiveBits(15))
	// fmt.Println(student.ActiveBits(17))
	// fmt.Println(student.ActiveBits(4))
	// fmt.Println(student.ActiveBits(11))
	// fmt.Println(student.ActiveBits(9))
	// fmt.Println(student.ActiveBits(12))
	// fmt.Println(student.ActiveBits(2))

	//SortList
	// var link *student.NodeI
	// link = student.ListPushBackNodeI(link, 1)
	// link = student.ListPushBackNodeI(link, 2)
	// link = student.ListPushBackNodeI(link, 3)
	// link = student.ListPushBackNodeI(link, 4)
	// link = student.ListPushBackNodeI(link, 5)
	// link = student.ListPushBackNodeI(link, 24)
	// link = student.ListPushBackNodeI(link, 25)
	// link = student.ListPushBackNodeI(link, 54)
	// student.PrintList(link)
	// link = student.SortListInsert(link, 33)
	// student.PrintList(link)

	//SortListMerge
	// var link *student.NodeI
	// var link2 *student.NodeI

	// link = student.ListPushBackNodeI(link, 0)
	// link = student.ListPushBackNodeI(link, 7)
	// link = student.ListPushBackNodeI(link, 39)
	// link = student.ListPushBackNodeI(link, 92)
	// link = student.ListPushBackNodeI(link, 97)
	// link = student.ListPushBackNodeI(link, 93)
	// link = student.ListPushBackNodeI(link, 91)
	// link = student.ListPushBackNodeI(link, 28)
	// link = student.ListPushBackNodeI(link, 64)

	// link2 = student.ListPushBackNodeI(link2, 80)
	// link2 = student.ListPushBackNodeI(link2, 92)
	// link2 = student.ListPushBackNodeI(link2, 27)
	// link2 = student.ListPushBackNodeI(link2, 30)
	// link2 = student.ListPushBackNodeI(link2, 85)
	// link2 = student.ListPushBackNodeI(link2, 81)
	// link2 = student.ListPushBackNodeI(link2, 75)
	// link2 = student.ListPushBackNodeI(link2, 70)

	// student.PrintList(student.SortedListMerge(link2, link))

	//ListRemoveIf
	// link := &student.List{}
	// student.ListPushBackNodeL(link, 98)
	// student.ListPushBackNodeL(link, 98)
	// student.ListPushBackNodeL(link, 33)
	// student.ListPushBackNodeL(link, 34)
	// student.ListPushBackNodeL(link, 33)
	// student.ListPushBackNodeL(link, 34)
	// student.ListPushBackNodeL(link, 33)
	// student.ListPushBackNodeL(link, 89)
	// student.ListPushBackNodeL(link, 33)
	// student.PrintListL(link)
	// student.ListRemoveIf(link, 34)
	// student.PrintListL(link)

	//BTreeTransplant
	// root := &student.TreeNode{Data: "01"}
	// student.BTreeInsertData(root, "07")
	// student.BTreeInsertData(root, "12")
	// student.BTreeInsertData(root, "10")
	// student.BTreeInsertData(root, "05")
	// student.BTreeInsertData(root, "02")
	// student.BTreeInsertData(root, "03")
	// node := student.BTreeSearchItem(root, "12")
	// replacement := &student.TreeNode{Data: "55"}
	// root = student.BTreeTransplant(root, node, replacement)
	// student.BTreeApplyInorder(root, fmt.Println)

	//BTreeApplyByLevel
	// root := &student.TreeNode{Data: "01"}
	// student.BTreeInsertData(root, "07")
	// student.BTreeInsertData(root, "12")
	// student.BTreeInsertData(root, "10")
	// student.BTreeInsertData(root, "05")
	// student.BTreeInsertData(root, "02")
	// student.BTreeInsertData(root, "03")
	// student.BTreeApplyByLevel(root, fmt.Println)

	//BTreeDeleteNode
	// root := &student.TreeNode{Data: "01"}
	// student.BTreeInsertData(root, "07")
	// student.BTreeInsertData(root, "12")
	// student.BTreeInsertData(root, "10")
	// student.BTreeInsertData(root, "05")
	// student.BTreeInsertData(root, "02")
	// student.BTreeInsertData(root, "03")
	// node := student.BTreeSearchItem(root, "02")
	// root = student.BTreeDeleteNode(root, node)
	// fmt.Println("After delete:")
	// student.BTreeApplyInorder(root, fmt.Println)

}
