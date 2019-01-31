package fileutil

import (
	"bytes"
	"crypto/md5"
	"io"
	"os"
)

func FileMD5(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	d := md5.New()
	_, err = io.Copy(d, f)
	if err != nil {
		return nil, err
	}

	return d.Sum(nil), nil
}

func Equal(f1, f2 string) bool {
	m1, err := FileMD5(f1)
	if err != nil {
		return false
	}
	m2, err := FileMD5(f2)
	if err != nil {
		return false
	}

	return bytes.Equal(m1, m2)
}
