package userService

import "net/http"

// handleGenericError writes an internal server error response if err is non-nil and returns true; returns false otherwise.
func handleGenericError(w http.ResponseWriter, errorMessage string, err error) bool {
	if err != nil {
		http.Error(w, errorMessage, http.StatusInternalServerError)
		return true
	}
	return false
}
