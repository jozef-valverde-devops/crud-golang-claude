package userService

import (
	"crud-golang/database"
	"crud-golang/model"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// UpdateUser updates the data of a user in the database.
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if handleGenericError(w, "Failed to convert parameter to integer", err) {
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if handleGenericError(w, "Failed to read request body!", err) {
		return
	}

	var updatedUser model.User
	if err = json.Unmarshal(body, &updatedUser); handleGenericError(w, "Failed to unmarshal request body!", err) {
		return
	}

	db, err := database.DbConnection()
	if handleGenericError(w, "Failed to connect to the database!", err) {
		return
	}
	defer db.Close()

	statement, err := db.Prepare("UPDATE users SET name = ?, email = ? WHERE id = ?")
	if handleGenericError(w, "Failed to create statement!", err) {
		return
	}
	defer statement.Close()

	if _, err := statement.Exec(updatedUser.Name, updatedUser.Email, ID); handleGenericError(w, "Failed to execute statement!", err) {
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
