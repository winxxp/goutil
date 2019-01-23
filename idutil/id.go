package idutil

import (
	"encoding/binary"
	"strconv"
)

type ID interface {
	String() string
	Byte() []byte
}

type UID uint64

func (i UID) String() string {
	return strconv.FormatUint(uint64(i), 10)
}

func (i UID) Byte() []byte {
	return []byte(i.String())
}

func EncodeBinaryUint64(i uint64) []byte {
	var b []byte = make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(i))
	return b
}

func DecodeBinaryUint64(b []byte) uint64 {
	return binary.LittleEndian.Uint64(b)
}

type IID int64

func (i IID) String() string {
	return strconv.FormatInt(int64(i), 10)
}

func (i IID) Byte() []byte {
	return []byte(i.String())
}
