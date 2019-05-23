package goutil

import (
	"github.com/winxxp/goutil/byteutil"
	"github.com/winxxp/goutil/chain"
	"github.com/winxxp/goutil/errutil"
	"github.com/winxxp/goutil/fileutil"
	"github.com/winxxp/goutil/ginutil"
	"github.com/winxxp/goutil/idutil"
	"github.com/winxxp/goutil/signutil"
	"github.com/winxxp/goutil/testutil/convey"
	"github.com/winxxp/goutil/testutil/matcher"
	"github.com/winxxp/goutil/textutil"
)

func Utils() []string {
	return []string{
		byteutil.Name(),
		chain.Name(),
		errutil.Name(),
		fileutil.Name(),
		ginutil.Name(),
		idutil.Name(),
		signutil.Name(),
		convey.Name(),
		textutil.Name(),
		matcher.Name(),
	}
}
