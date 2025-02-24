package isomorphic_string

import (
	"testing"
)

func TestIsIsomorphic(t *testing.T) {
	res := isIsomorphic("egg", "add")
	t.Log(res)

	res = isIsomorphic("dof", "add")
	t.Log(res)

	res = isIsomorphic("paper", "title")
	t.Log(res)
}
