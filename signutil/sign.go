package signutil

import (
	"crypto/md5"
	"fmt"
	"io"
	"time"
)

var (
	TimeLayout = "20060102150405.000Z07:00"
)

//
// md5(signKey+spaceId+d)
// signKey:签名密钥,由data中心发放，默认密钥：std-jim-cfs
// spaceId:集群ID
// date:当前请求时间
func Sign(key string, spaceID uint64, date time.Time) string {
	h := md5.New()

	io.WriteString(h, key)
	io.WriteString(h, fmt.Sprint(spaceID))
	io.WriteString(h, date.Format(TimeLayout))

	return fmt.Sprintf("%x", h.Sum(nil))
}

func SignNow(key string, spaceID uint64) string {
	return Sign(key, spaceID, time.Now())
}
