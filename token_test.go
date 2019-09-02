package goutil

import "testing"

func TestGenToken(t *testing.T) {
	tk := GenToken()
	t.Log(tk)
}
