package main

import (
	"fmt"
	"slices"
	"sort"
)

// ============================================================================
// BUILT-IN DATA TYPE SORTING
// ============================================================================

// 1. INTEGERS
//    sort.Ints(slice []int) - sorts in ascending order (modifies slice in-place)
//    sort.IntsAreSorted(slice []int) bool - checks if slice is sorted
//    sort.SearchInts(slice []int, x int) int - binary search (slice must be sorted)

// 2. STRINGS
//    sort.Strings(slice []string) - sorts lexicographically (modifies slice in-place)
//    sort.StringsAreSorted(slice []string) bool - checks if slice is sorted
//    sort.SearchStrings(slice []string, x string) int - binary search (slice must be sorted)

// 3. FLOAT64
//    sort.Float64s(slice []float64) - sorts in ascending order (modifies slice in-place)
//    sort.Float64sAreSorted(slice []float64) bool - checks if slice is sorted
//    sort.SearchFloat64s(slice []float64, x float64) int - binary search (slice must be sorted)

// 4. DESCENDING ORDER
//    sort.Sort(sort.Reverse(sort.IntSlice(slice)))
//    sort.Sort(sort.Reverse(sort.StringSlice(slice)))
//    sort.Sort(sort.Reverse(sort.Float64Slice(slice)))

// 5. CUSTOM SORTING
//    sort.Slice(slice interface{}, less func(i, j int) bool)
//    sort.SliceStable(slice interface{}, less func(i, j int) bool) - stable sort

// 6. SEARCH FUNCTIONS
//    sort.Search(n int, f func(int) bool) int - binary search with custom condition

func main() {
	// int sorting
	intslice := []int{3, -2, 4, 5, 6, 2, 3, 10, -1, 1}
	fmt.Println("intslice: ", intslice)

	sort.Ints(intslice)
	fmt.Println("after sorting intslice: ", intslice)

	intslice = []int{3, -2, 4, 5, 6, 2, 3, 10, -1, 1}
	slices.Sort(intslice)
	fmt.Println("after sorting intslice: ", intslice)

	// string sorting
	strSlice := []string{"hello", "dhairya", "how", "are", "you"}
	fmt.Println("strSlice: ", strSlice)

	sort.Strings(strSlice)
	fmt.Println("after sorting strSlice: ", strSlice)

	strSlice = []string{"hello", "dhairya", "how", "are", "you"}
	slices.Sort(strSlice)
	fmt.Println("after sorting strSlice: ", strSlice)

	// int sorting on the operation results of int slice
	sort.Slice(intslice, func(i, j int) bool {
		return abs(intslice[i]) < abs(intslice[j])
	})
	fmt.Println("afte reverse sorting: ", intslice)

	slices.SortFunc(strSlice, func(a, b string) int {
		if len(a) < len(b) {
			return 1
		}
		return -1
	})
	fmt.Println("sorting string slice: ", strSlice)

	type Person struct {
		Age  int
		Name string
	}

	people := []Person{{1, "Dhairya"}, {3, "Aevin"}, {10, "Zar"}, {5, "Zar"}}
	sort.Slice(people, func(i, j int) bool {
		if people[i].Name < people[j].Name {
			return true
		}
		return people[i].Age < people[j].Age
	})
	fmt.Println("After sorting people: ", people)

	type Pair struct {
		Time int
		Dis  int
	}

	pairs := []Pair{
		{1, 6}, {2, 5}, {3, 4}, {4, 3}, {5, 2}, {6, 1},
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Time > pairs[j].Time
	})
	fmt.Println("After sorting time and dis: ", time, dis)

}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
