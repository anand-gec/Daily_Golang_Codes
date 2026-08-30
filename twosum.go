/*  Input: arr = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because arr[0] + arr[1] == 9, we return [0, 1].
*/

package main

import "fmt"

func twoSum(arr []int, target int) []int {
	var arr2 = []int{}
	for i := 0; i < len(arr); i++ {
		for j := 1; j < len(arr); j++ {
			if arr[i]+arr[j] == target {
				arr2 = append(arr2, i, j)
				return arr2
			}
		}
	}
	return arr2
}

func main() {
	var arr = [...]int{2, 7, 11, 15}
	var target = 9
	fmt.Println(twoSum(arr[:], target))

}
