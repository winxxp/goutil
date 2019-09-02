package goutil

import "strconv"

func Uint64ToByte(i uint64) []byte {
	return []byte(strconv.FormatUint(i, 10))
}
func Int64ToByte(i int64) []byte {
	return []byte(strconv.FormatInt(i, 10))
}
