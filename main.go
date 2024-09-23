package main

import (
	"fmt"

	_ "bulentkeskin.com/dsa/arrays"
	"bulentkeskin.com/dsa/binarytrees"
	_ "bulentkeskin.com/dsa/linkedlists"
	_ "bulentkeskin.com/dsa/recursion"
)

func main() {

	//fmt.Println(arrays.Anagrams("hello", "olelh"))

	//fmt.Println(string(arrays.MostFrequentChar("eleventennine")))

	//fmt.Println(arrays.PairSum([]int{3, 2, 5, 4, 1}, 8))

	//fmt.Println(arrays.Uncompress("3n12e2z"))

	//fmt.Println(arrays.Compress("nnneeeeeeeeeeeez"))

	//fmt.Println(arrays.Intersect([]int{1, 2, 3, 4, 5}, []int{4, 5, 6, 7, 8}))

	//fmt.Println(arrays.FiveSort([]int{5, 2, 5, 6, 5, 1, 10, 2, 5, 5}))

	//fmt.Println(recursion.SumNumbers([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))

	//fmt.Println(recursion.ReverseString("hello"))

	//fmt.Println(recursion.IsPalindrome("abcdefedcba"))

	//fmt.Println(recursion.Fibonacci(7))

	//fmt.Println(linkedlists.Values(linkedlists.IntLinkedList))

	//fmt.Println(linkedlists.SumList(linkedlists.IntLinkedList))

	//fmt.Println(linkedlists.HasValue(linkedlists.IntLinkedList, 4))

	//fmt.Println(linkedlists.Reverse(linkedlists.IntLinkedList))

	//fmt.Println(linkedlists.Zip(linkedlists.IntLinkedList, linkedlists.IntLinkedList2))

	//fmt.Println(linkedlists.MergeSorted(linkedlists.IntLinkedList, linkedlists.IntLinkedList2))

	//fmt.Println(linkedlists.LongestStreak(linkedlists.IntLinkedList))

	//fmt.Println(linkedlists.RemoveNode(linkedlists.IntLinkedList, 40))

	//fmt.Println(linkedlists.InsertNode(linkedlists.IntLinkedList, 25, 2))

	//fmt.Println(linkedlists.Create([]int{1, 2, 3, 4, 5}))

	//fmt.Println(binarytrees.DepthFirst(binarytrees.IntBinaryTree))

	//fmt.Println(binarytrees.BreadthFirst(binarytrees.IntBinaryTree))

	//fmt.Println(binarytrees.TreeSum(binarytrees.IntBinaryTree))

	//fmt.Println(binarytrees.TreeContains(binarytrees.IntBinaryTree, 5))

	//fmt.Println(binarytrees.MaxValue(binarytrees.IntBinaryTree))

	//fmt.Println(binarytrees.MaxRootToLeafPathSum(binarytrees.ArrayToBinaryTree([]int{5, 11, 54, 20, 15, -1, -1, -1, -1, 1, 3})))

	fmt.Println(binarytrees.FindPath(binarytrees.IntBinaryTree, 7))
}
