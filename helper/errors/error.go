package errors

import "errors"

func ErrorPanic(err error) {
	if err != nil {
		panic(err)
	}
}

var (
	ERR_USER_NOT_FOUND = errors.New("User Not Found")
	ERR_BAD_REQUEST = errors.New("Bad request")
)
