package userService

import (
	"crud-golang/database"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// DeleteUser deletes a user from the database.
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if handleGenericError(w, "Failed to convert parameter to integer", err) {
		return
	}

	db, err := database.DbConnection()
	if handleGenericError(w, "Failed to connect to the database!", err) {
		return
	}
	defer db.Close()

	statement, err := db.Prepare("DELETE FROM users WHERE id = ?")
	if handleGenericError(w, "Failed to create statement!", err) {
		return
	}
	defer statement.Close()

	if _, err := statement.Exec(ID); handleGenericError(w, "Failed to execute statement!", err) {
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
