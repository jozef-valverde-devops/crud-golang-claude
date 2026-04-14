package userService

import (
	"crud-golang/database"
	"crud-golang/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// CreateUser inserts a user into the database.
func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if handleGenericError(w, "Failed to read request body!", err) {
		return
	}

	var newUser model.User
	if err = json.Unmarshal(body, &newUser); handleGenericError(w, "Failed to unmarshal request body!", err) {
		return
	}

	db, err := database.DbConnection()
	if handleGenericError(w, "Failed to connect to the database!", err) {
		return
	}
	defer db.Close()

	statement, err := db.Prepare("INSERT INTO users (name, email) VALUES (?, ?)")
	if handleGenericError(w, "Failed to create statement!", err) {
		return
	}
	defer statement.Close()

	result, err := statement.Exec(newUser.Name, newUser.Email)
	if handleGenericError(w, "Failed to execute statement!", err) {
		return
	}

	createdID, err := result.LastInsertId()
	if handleGenericError(w, "Failed to retrieve last insert ID!", err) {
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Successfully created user %d", createdID)
}
