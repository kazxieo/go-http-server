package middleware

import (
	"log"
	"net/http"

	"github.com/kazxieo/go-http-server/database"
)

func Authentication(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		user := r.FormValue("user")
		password := r.FormValue("password")

		log.Println(user, password)

		if pass, ok := database.PasswordDatabase[user]; !ok || pass != password {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
