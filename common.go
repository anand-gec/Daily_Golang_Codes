// Common string
package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	common := strs[0]
	for _, value := range strs[1:] {
		if len(common) > len(value) {
			common = common[:len(value)]
		}
		for len(common) > 0 && !strings.HasPrefix(value, common) {
			common = common[:len(common)-1]
		}
		if len(common) == 0 {
			break
		}
	}
	return common
}

func main() {
	str := [...]string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(str[:]))

}
