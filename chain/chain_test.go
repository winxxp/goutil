package chain

import (
	"errors"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestC_Handles(t *testing.T) {
	Convey("Chain", t, func() {
		var (
			val     = 0
			errTest = errors.New("test")
		)

		f1 := func() error { val++; return nil }
		f2 := func(i int) { val += i }
		f3 := func() error { return errTest }
		f4 := func() error { val++; return nil }

		Convey("Error", func() {
			err := New().Handles(f1, func() error { f2(2); return nil }, f3, f4).Run()
			So(err, ShouldBeError, errTest)
			So(val, ShouldEqual, 3)
		})

		Convey("No Error", func() {
			err := New().Handles(f1, f4).Run()
			So(err, ShouldBeNil)
			So(val, ShouldEqual, 2)
		})

		Convey("Run", func() {
			err := Run(f1, f4)
			So(err, ShouldBeNil)
			So(val, ShouldEqual, 2)
		})
	})
}
