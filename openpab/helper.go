package openpab

import (
	"strings"
)

// JoinVc |::| 拼接
func JoinVc(s ...string) string {
	return strings.Join(s, "|::|")
}
