package errors

import "net/http"

func GetErrorCode(err error) int {
	switch err {
	case ERR_USER_NOT_FOUND:
		return http.StatusNotFound
	case ERR_BAD_REQUEST:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
