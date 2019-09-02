package goutil

import "fmt"

type HexByte byte

func (b HexByte) String() string {
	return fmt.Sprintf("%02X", byte(b))
}

type HexBytes []byte

func (b HexBytes) String() string {
	var hb = make([]HexByte, len(b))
	for i, v := range b {
		hb[i] = HexByte(v)
	}

	return fmt.Sprint(hb)
}
