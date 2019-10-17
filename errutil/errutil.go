package errutil

import "github.com/pkg/errors"

func CheckErrPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func Is(err, wrapErr error) bool {
	return err == errors.Cause(wrapErr)
}
