package errutil

func Name() string {
	return "errutil"
}

//ErrResult return err state
// defStr[0], defStr[1] = fail, succeed comment
func ErrResult(err error, defStr ...string) string {
	var (
		defFail    = "fail"
		defSucceed = "succeed"
	)

	if len(defStr) > 0 {
		defFail = defStr[0]
	}
	if len(defStr) > 1 {
		defSucceed = defStr[1]
	}

	if err != nil {
		return defFail + ":" + err.Error()
	}

	return defSucceed
}
