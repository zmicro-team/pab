package utils

import (
	"sort"
	"strings"
)

func ConcatSortMap(mp map[string]string, sep1, sep2 string) string {
	ks := make([]string, 0, len(mp))
	n := 0
	for k, v := range mp {
		if v != "" {
			ks = append(ks, k)
			n += len(k) + len(v)
		}
	}
	n += len(ks)*2 - 1 // 加上sep1和sep2的总个数

	sort.Strings(ks)

	bs := strings.Builder{}
	bs.Grow(n)
	for _, k := range ks {
		if bs.Len() > 0 {
			bs.WriteString(sep2)
		}
		bs.WriteString(k)
		bs.WriteString(sep1)
		bs.WriteString(mp[k])
	}
	return bs.String()
}
