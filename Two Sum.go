package main

import "fmt"

func twoSum(nums []int, target int) []int {

	//sort.Ints(nums[:])

	hash_map := make(map[int]int)

	for k := 0; k < len(nums); k++ {
		desire_num := target - nums[k]
		requiredIdx, isPresent := hash_map[desire_num]
		if isPresent {
			return []int{requiredIdx, k}
		}

		hash_map[nums[k]] = k

	}
	return []int{}
}

func main() {
	input := [4]int{-3, 4, 3, 90}
	target := 0
	fmt.Println(twoSum(input[:], target))
}
