package userService

import (
	"crud-golang/database"
	"crud-golang/model"
	"encoding/json"
	"net/http"
)

// GetUsers retrieves all users from the database.
func GetUsers(w http.ResponseWriter, r *http.Request) {
	db, err := database.DbConnection()
	if handleGenericError(w, "Failed to connect to the database!", err) {
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users")
	if handleGenericError(w, "Failed to retrieve users!", err) {
		return
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); handleGenericError(w, "Failed to scan users!", err) {
			return
		}
		users = append(users, user)
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(users); handleGenericError(w, "Failed to convert users to JSON!", err) {
		return
	}
}
