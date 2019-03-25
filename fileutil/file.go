package fileutil

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

func Name() string {
	return "fileutil"
}

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

//CopyFile Copy regular file
func CopyFile(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()

	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}
