package main

import (
	"auth-api/auth"
	"auth-api/db"
	"auth-api/models/dtos"
	"auth-api/services"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	db.ConnectDB()
	db.GetInstance().Migration()

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	router.Post("/signup/credential", func(w http.ResponseWriter, r *http.Request) {
		var signupData dtos.CredentialSignupDto
		parsingErr := json.NewDecoder(r.Body).Decode(&signupData)
		if parsingErr != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&dtos.ErrorDto{Message: parsingErr.Error()})
			return
		}

		result, err := services.CredentialSignup(&signupData)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	})

	router.Post("/signup/oauth", func(w http.ResponseWriter, r *http.Request) {
		var signupData dtos.OAuthDto
		parsingErr := json.NewDecoder(r.Body).Decode(&signupData)
		if parsingErr != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&dtos.ErrorDto{Message: parsingErr.Error()})
			return
		}

		result, err := services.OAuthSignup(&signupData)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	})

	router.Post("/login/credential", func(w http.ResponseWriter, r *http.Request) {
		var signinData dtos.CredentialSigninDto
		parsingErr := json.NewDecoder(r.Body).Decode(&signinData)
		if parsingErr != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&dtos.ErrorDto{Message: parsingErr.Error()})
			return
		}

		token, err := services.UserCredentialLogin(&signinData)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(token)
	})

	router.Post("/login/oauth", func(w http.ResponseWriter, r *http.Request) {
		var signinData dtos.OAuthDto
		parsingErr := json.NewDecoder(r.Body).Decode(&signinData)
		if parsingErr != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&dtos.ErrorDto{Message: parsingErr.Error()})
			return
		}

		token, err := services.UserOAuthLogin(&signinData)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(token)
	})

	router.With(auth.Authenticate).Get("/profile", func(w http.ResponseWriter, r *http.Request) {
		currentUser := auth.GetAuthUser(r)
		profile, errDto := services.GetUserByID(currentUser.ID)
		if errDto != nil {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(errDto)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(profile)
	})

	router.Post("/email-verify", func(w http.ResponseWriter, r *http.Request) {
		var verificationData dtos.EmailVerifyDto
		parsingErr := json.NewDecoder(r.Body).Decode(&verificationData)
		if parsingErr != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&dtos.ErrorDto{Message: parsingErr.Error()})
			return
		}

		user, err := services.VerifyUserEmail(&verificationData)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)
	})

	http.ListenAndServe(":8000", router)

}
