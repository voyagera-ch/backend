package loginHandler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/voyagera/backend/auth"
	voyageraUser "github.com/voyagera/backend/internal/models/user"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var u voyageraUser.User
	json.NewDecoder(r.Body).Decode(&u)
	fmt.Printf("The user request value %v", u)

	if u.Username == "Chek" && u.Password == "123456" {
		tokenString, err := auth.CreateToken(u.Username)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, tokenString)
		return
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Invalid credentials")
	}
}

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	username, err := auth.ValidateRequest(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	fmt.Fprint(w, fmt.Sprintf("Welcome to the the protected area %s", username))

}
