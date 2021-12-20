package reflectx

import "testing"

func TestIsNull(t *testing.T) {

	a := 1
	b := 0

	type tt struct {
		A int
	}

	t.Log(Empty(a))
	t.Log(Empty(b))
	t.Log(Empty(tt{}))
	t.Log(Empty(tt{A: 1}))
	var (
		ttt1 *tt
		ttt2 = new(tt)
	)
	t.Log(Empty(ttt1))
	t.Log(Empty(ttt2))
	t.Log(Empty(""))
	t.Log(Empty(".Elem()"))

}
