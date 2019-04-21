package dictionariesandhashmaps

import (
	"fmt"
	"github.com/zrma/1d1c/hacker_rank/common/utils/str"
	"sort"
	"strings"
)

func checkMagazine(magazine []string, note []string) {
	sort.Sort(str.SortAdapter(magazine))
	sort.Sort(str.SortAdapter(note))

	i := len(magazine) - 1
	j := len(note) - 1

	for i >= 0 && j >= 0 {
		diff := strings.Compare(magazine[i], note[j])
		if diff > 0 {
			i--
		} else if diff == 0 {
			i--
			j--
		} else {
			fmt.Println("No")
			return
		}
	}

	fmt.Println("Yes")
}