package response

import "net/http"

type Error struct {
	Error string `json:"error"`
	Code  string `json:"error_code"`
}

func BuildError(error string, statusCode int) Error {
	return Error{
		Error: error,
		Code:  http.StatusText(statusCode),
	}
}

func BuildInternalError() Error {
	return BuildError("Something went wrong", http.StatusInternalServerError)
}
