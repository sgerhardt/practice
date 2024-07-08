package InterviewCake

//import "fmt"

/*
Suppose we have some input data describing a graph of relationships between parents and children over multiple generations. The data is formatted as a list of (parent, child) pairs, where each individual is assigned a unique positive integer identifier.

For example, in this diagram, 6 and 8 have common ancestors of 4 and 14.

               15
              / \
         14  13  21
         |   |
1   2    4   12
 \ /   / | \ /
  3   5  8  9
   \ / \     \
    6   7     11

pairs = [
    (1, 3), (2, 3), (3, 6), (5, 6), (5, 7), (4, 5),
    (15, 21), (4, 8), (4, 9), (9, 11), (14, 4), (13, 12),
    (12, 9), (15, 13)
]

Write a function that takes this data and the identifiers of two individuals as inputs and returns true if and only if they share at least one ancestor.

Sample input and output:

hasCommonAncestor(pairs, 3, 8)   => false
hasCommonAncestor(pairs, 5, 8)   => true
hasCommonAncestor(pairs, 6, 8)   => true
hasCommonAncestor(pairs, 6, 9)   => true
hasCommonAncestor(pairs, 1, 3)   => false
hasCommonAncestor(pairs, 3, 1)   => false
hasCommonAncestor(pairs, 7, 11)  => true
hasCommonAncestor(pairs, 6, 5)   => true
hasCommonAncestor(pairs, 5, 6)   => true
hasCommonAncestor(pairs, 3, 6)   => true
hasCommonAncestor(pairs, 21, 11) => true

n: number of pairs in the input
*/

func main() {
	/*  pairs := [][]int {
	    []int{5, 6},
	    []int{1, 3},
	    []int{2, 3},
	    []int{3, 6},
	    []int{15, 12},
	    []int{5, 7},
	    []int{4, 5},
	    []int{4, 9},
	    []int{9, 12},
	    []int{30, 16},
	} */

	pairs := [][]int{
		[]int{1, 3},
		[]int{2, 3},
		[]int{3, 6},
		[]int{5, 6},
		[]int{5, 7},
		[]int{4, 5},
		[]int{15, 21},
		[]int{4, 8},
		[]int{4, 9},
		[]int{9, 11},
		[]int{14, 4},
		[]int{13, 12},
		[]int{12, 9},
		[]int{15, 13},
	}

	zero, exactlyOneParent := findNodesWithZeroAndOneParents(pairs)
	for _, val := range zero {
		println("zero case")
		println(val)
	}
	for _, val := range exactlyOneParent {
		println("exactly one parent case")
		println(val)
	}
}

func hasCommonAncestor([][]int, int, int) bool {

	return false
}

func findNodesWithZeroAndOneParents(pairs [][]int) (zeroParents []int, exactlyOneParent []int) {
	zero := []int{}
	exactlyOne := []int{}

	individualToParentCount := map[int]int{}

	parents := []int{}

	for _, pair := range pairs {
		for idx, individual := range pair {
			if idx == 0 {
				parents = append(parents, individual)
			}
			if idx == 1 {
				_, found := individualToParentCount[individual]
				if found {
					individualToParentCount[individual]++
				} else {
					individualToParentCount[individual] = 1
				}
			}
		}
	}

	for k, v := range individualToParentCount {
		if v == 1 {
			exactlyOne = append(exactlyOne, k)
		}
	}

	for _, i := range parents {
		_, found := individualToParentCount[i]
		if found {
		} else {
			zero = append(zero, i)
		}
	}

	zeroParent := map[int]bool{}
	zeroDeduped := []int{}
	for _, val := range zero {
		zeroParent[val] = true
	}
	for k, _ := range zeroParent {
		zeroDeduped = append(zeroDeduped, k)
	}

	return zeroDeduped, exactlyOne
}
