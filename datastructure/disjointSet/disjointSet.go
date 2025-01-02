// Implement a Disjoint set in Golang with union by rank and size

package disjointset

// import (
// 	"fmt"
// )

type DisjointSet struct {
	parent []int
	size   []int
	rank   []int
}

// Constructor to create a new disjoint set
func newDisjointSet(nodes int) *DisjointSet {
	parent := make([]int, nodes+1)
	size := make([]int, nodes+1)
	rank := make([]int, nodes+1)
	for i := range parent {
		parent[i] = i
		size[i] = 1
		rank[i] = 0
	}
	return &DisjointSet{
		parent: parent,
		size:   size,
		rank:   rank,
	}
}

func (ds *DisjointSet) findUltimateParent(node int) int {
	if node == ds.parent[node] {
		return node
	}
	ds.parent[node] = ds.findUltimateParent(ds.parent[node])
	return ds.parent[node]
}

func (ds *DisjointSet) unionBySize(u, v int) {
	ultimateParentU := ds.findUltimateParent(u)
	ultimateParentV := ds.findUltimateParent(v)

	sizeOfParentU := ds.size[ultimateParentU]
	sizeOfParentV := ds.size[ultimateParentV]

	if sizeOfParentU > sizeOfParentV {
		ds.parent[ultimateParentV] = ultimateParentU
		ds.size[ultimateParentU] = ds.size[ultimateParentU] + ds.size[ultimateParentV]
	} else if sizeOfParentU < sizeOfParentV {
		ds.parent[ultimateParentU] = ultimateParentV
		ds.size[ultimateParentV] = ds.size[ultimateParentV] + ds.size[ultimateParentU]
	} else {
		ds.parent[ultimateParentV] = ultimateParentU
		ds.size[ultimateParentU] = ds.size[ultimateParentU] + ds.size[ultimateParentV]
	}

}

func (ds *DisjointSet) unionByRank(u, v int) {
	ultimateParentU := ds.findUltimateParent(u)
	ultimateParentV := ds.findUltimateParent(v)

	rankOfParentU := ds.rank[ultimateParentU]
	rankOfParentV := ds.rank[ultimateParentV]

	if rankOfParentU > rankOfParentV {
		ds.parent[ultimateParentV] = ultimateParentU
	} else if rankOfParentU < rankOfParentV {
		ds.parent[ultimateParentU] = ultimateParentV
	} else {
		ds.parent[ultimateParentV] = ultimateParentU
		ds.rank[ultimateParentU]++
	}
}

// func main() {
// 	new_set := newDisjointSet(7)
// 	fmt.Println("Parent array", new_set.parent)
// 	// fmt.Println("Size array", new_set.size)
// 	fmt.Println("Rank array", new_set.rank)
// 	new_set.unionByRank(1, 2)
// 	new_set.unionByRank(2, 3)
// 	new_set.unionByRank(4, 5)
// 	new_set.unionByRank(6, 7)
// 	new_set.unionByRank(5, 6)
// 	if new_set.findUltimateParent(3) == new_set.findUltimateParent(7) {
// 		fmt.Println("Same parent")
// 	} else {
// 		fmt.Println("Different parent")
// 	}
// 	new_set.unionByRank(3, 7)
// 	if new_set.findUltimateParent(3) == new_set.findUltimateParent(7) {
// 		fmt.Println("Same parent")
// 	} else {
// 		fmt.Println("Different parent")
// 	}
// 	new_set.findUltimateParent(2)
// 	fmt.Println("Parent array", new_set.parent)
// 	// fmt.Println("Size array", new_set.size)
// 	fmt.Println("Rank array", new_set.rank)

// }
