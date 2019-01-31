package fileutil

import (
	. "github.com/smartystreets/goconvey/convey"
	"io/ioutil"
	"os"
	"testing"
)

func TestEqual(t *testing.T) {
	Convey("File", t, func() {
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

			Convey("Equal", func() {
				Convey("No error", func() {
					So(Equal(filename, filename), ShouldBeTrue)
				})
				Convey("Error", func() {
					So(Equal(filename, filename+"0"), ShouldBeFalse)
				})
			})
		})
	})
}
