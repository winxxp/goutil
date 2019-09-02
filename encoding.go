package goutil

import (
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
)

func UTF82ANSIDecode(s []byte) ([]byte, error) {
	input := bytes.NewReader(s)
	output := transform.NewReader(input, simplifiedchinese.GB18030.NewEncoder())
	return ioutil.ReadAll(output)
}

func ANSI2UTF8Decode(s []byte) ([]byte, error) {
	input := bytes.NewReader(s)
	output := transform.NewReader(input, simplifiedchinese.GB18030.NewDecoder())
	d, e := ioutil.ReadAll(output)
	if e != nil {
		return nil, e
	}
	return d, nil
}
