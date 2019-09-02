package goutil

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestSign(t *testing.T) {
	Convey("Sign", t, func() {
		now := time.Date(2018, 11, 06, 21, 48, 9, 887862500, time.Now().Location())
		spaceID := uint64(1)
		key := "datas"
		expected := "927154ecd794d079dd66ea036e41b0af"

		Convey("T1", func() {
			Println(now.Format(TimeLayout))

			s := Sign(key, spaceID, now)
			Println(s)
			So(s, ShouldEqual, expected)
		})

		Convey("Time Add Nanosecond Same Sign", func() {
			tm := now.Add(time.Nanosecond)
			Println(tm.Format(TimeLayout))

			s := Sign(key, spaceID, tm)
			Println(s)

			So(s, ShouldEqual, expected)
		})

		Convey("Time Add Microsecond Same Sign", func() {
			tm := now.Add(time.Microsecond)
			Println(tm.Format(TimeLayout))

			s := Sign(key, spaceID, tm)
			Println(s)

			So(s, ShouldEqual, expected)
		})

		Convey("Time Add Millisecond Different Sign", func() {
			tm := now.Add(time.Millisecond)
			Println(tm.Format(TimeLayout), " ", now.Format(TimeLayout))

			s := Sign(key, spaceID, tm)
			Println(s)

			So(s, ShouldNotEqual, expected)
		})

		Convey("Time Add Second Different Sign", func() {
			tm := now.Add(time.Second)
			Println(tm.Format(TimeLayout), " ", now.Format(TimeLayout))

			s := Sign(key, spaceID, tm)

			Println(s)
			So(s, ShouldNotEqual, expected)
		})
	})
}

func TestSignNow(t *testing.T) {
	Convey("SignNow", t, func() {
		now := time.Now()
		actual := SignNow("abc", 1)
		expected := Sign("abc", 1, now)
		So(actual, ShouldEqual, expected)

	})
}
