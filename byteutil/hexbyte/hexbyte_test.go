package hexbyte

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestHexByte(t *testing.T) {
	Convey("HexByte", t, func() {
		Convey("Byte", func() {
			i := HexByte(128)
			s := i.String()
			So(s, ShouldEqual, "80")
		})

		Convey("Byte Array", func() {
			a := HexBytes{10, 128, 255}
			s := a.String()
			So(s, ShouldEqual, "[0A 80 FF]")
		})
	})

}
