package goutil

import (
	"crypto/md5"
	"fmt"
	"time"
)

func GenToken() string {
	now := time.Now()

	m := md5.New()
	_, _ = fmt.Fprintf(m, "%s", now.Format(time.RFC3339Nano))
	return fmt.Sprintf("%x", m.Sum(nil))
}
