package goutil

import (
	. "github.com/smartystreets/goconvey/convey"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestEqual(t *testing.T) {
	Convey("File", t, func(c C) {
		var (
			err      error
			filename string
			_        = err
		)

		Convey("Init", func() {
			f, err := ioutil.TempFile("", "fileutil.*")
			So(err, ShouldBeNil)
			defer f.Close()
			filename = f.Name()
			defer os.Remove(filename)

			Println("filename:", filename)

			_, err = f.WriteString("12345678")
			So(err, ShouldBeNil)

			Convey("FileEqual", func() {
				Convey("No error", func() {
					So(FileEqual(filename, filename), ShouldBeTrue)
				})
				Convey("Error", func() {
					So(FileEqual(filename, filename+"0"), ShouldBeFalse)
				})
			})
		})

		Convey("CopyFile", func() {
			dir, err := ioutil.TempDir("", "")
			So(err, ShouldBeNil)

			src := filepath.Join(dir, "src")
			dst := filepath.Join(dir, "dst")

			err = ioutil.WriteFile(src, []byte("hello"), os.ModePerm)
			So(err, ShouldBeNil)

			_, err = CopyFile(src, dst)
			So(err, ShouldBeNil)
		})

		Convey("Copy2UniqueFile", func() {
			dir, err := ioutil.TempDir("", "")
			So(err, ShouldBeNil)

			src := filepath.Join(dir, "src")
			err = ioutil.WriteFile(src, []byte("hello"), os.ModePerm)
			So(err, ShouldBeNil)

			_, filename, err := Copy2UniqueFile(src, dir, "test")
			So(err, ShouldBeNil)
			Println("fileName: ", filename)
		})
	})
}
