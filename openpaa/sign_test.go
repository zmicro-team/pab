package openpaa

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type B struct {
	C []string
	D bool
}

type A struct {
	B B
	C uint
}

func TestName(t *testing.T) {
	mp, err := struct2Map(A{
		B: B{
			C: []string{"1", "2", "3"},
			D: false,
		},
		C: 100,
	})
	require.NoError(t, err)
	t.Logf("%+v", mp)

	t.Logf("%+v", needSignMap(mp))
}
