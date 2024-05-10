package errors

import "errors"

func ErrorPanic(err error) {
	if err != nil {
		panic(err)
	}
}

var (
	ERR_USER_NOT_FOUND  = errors.New("User Not Found")
	ERR_WRONG_PASSWORD  = errors.New("Password incorrect")
	ERR_BAD_REQUEST     = errors.New("Bad request")
	ERR_REGISTER        = errors.New("Failed to register")
	ERR_BCRYPT_PASSWORD = errors.New("Error Bcrypt password")
	ERR_TOKEN           = errors.New("Error creating new token")
)
