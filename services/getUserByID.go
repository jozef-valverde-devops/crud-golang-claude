package userService

import (
	"crud-golang/database"
	"crud-golang/model"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetUserByID retrieves a user from the database by ID.
func GetUserByID(w http.ResponseWriter, r *http.Request) {
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

	row, err := db.Query("SELECT * FROM users WHERE id = ?", ID)
	if handleGenericError(w, "Failed to retrieve user "+strconv.FormatUint(ID, 10), err) {
		return
	}
	defer row.Close()

	var user model.User
	if row.Next() {
		if err := row.Scan(&user.ID, &user.Name, &user.Email); handleGenericError(w, "Failed to scan users!", err) {
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); handleGenericError(w, "Failed to convert user to JSON!", err) {
		return
	}
}
