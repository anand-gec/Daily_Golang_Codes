/*  Input: arr = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because arr[0] + arr[1] == 9, we return [0, 1].
*/

package main

import "fmt"

func twoSum(arr []int, target int) []int {
	var arr2 = []int{}
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == target {
				arr2 = append(arr2, i, j)
				return arr2
			}
		}
	}
	return arr2
}

func twoSums(arr []int, target int) []int {
	maps := make(map[int]int)
	for index, val := range arr {
		sub := target - val
		if idx, ok := maps[sub]; ok {
			return []int{idx, index}
		}
		maps[val] = index
	}
	return nil
}

func main() {
	var arr = [...]int{2, 5, 5, 11}
	var target = 10
	fmt.Println(twoSum(arr[:], target))

	var targets = 10
	var arr3 = [...]int{2, 5, 5, 11}
	fmt.Println(twoSums(arr3[:], targets))

}
