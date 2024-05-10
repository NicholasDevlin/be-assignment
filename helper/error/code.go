package errors

import "net/http"

func GetErrorCode(err error) int {
	switch err {
	case ERR_USER_NOT_FOUND:
		return http.StatusNotFound
	case ERR_BAD_REQUEST:
		return http.StatusBadRequest
	case ERR_BCRYPT_PASSWORD:
		return http.StatusInternalServerError
	case ERR_TOKEN:
		return http.StatusInternalServerError
	case ERR_WRONG_PASSWORD:
		return http.StatusBadRequest
	case ERR_REGISTER:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}
