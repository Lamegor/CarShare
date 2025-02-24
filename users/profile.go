package users

import (
	"carrent/db"
	"encoding/json"
	"net/http"
)

func ViewProfile(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(db.User)

	response := make(map[string]interface{})
	response["username"] = user.Login
	response["id"] = user.ID
	response["first_name"] = user.FirstName
	response["second_name"] = user.SecondName
	response["last_name"] = user.LastName
	response["phone"] = user.ContactPhone
	response["e-mail"] = user.Email
	response["gender"] = user.Gender

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func SendLicense(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(db.User)
	_ = user
}
