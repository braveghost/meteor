package errutil

func CheckErrPanic(err error) {
	if err != nil {
		panic(err)
	}
}
