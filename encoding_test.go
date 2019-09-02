package goutil

import (
	. "github.com/smartystreets/goconvey/convey"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestCopyFile(t *testing.T) {
	Convey("Copy File", t, func() {
		src := "encoding_test.go"
		dst := filepath.Join(os.TempDir(), src)

		os.Remove(dst)

		Convey("Copy", func() {
			_, err := CopyFile(src, dst)
			So(err, ShouldBeNil)

			data, err := ioutil.ReadFile(dst)
			So(err, ShouldBeNil)
			Println(string(data))
		})

		Convey("Dir", func() {
			_, err := CopyFile(".", dst)
			So(err, ShouldNotBeNil)
			Println(err.Error())
		})
	})
}
