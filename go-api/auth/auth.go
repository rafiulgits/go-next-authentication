package auth

import (
	"auth-api/models/domains"
	"auth-api/models/dtos"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

const JWT_SECRET = "abcefghijklmopqrst"
const JWT_TTL_HOUR = 24

const (
	KeyUser         = "user"
	KeyUserID       = "id"
	KeyTokenExpired = "exp"
)

type AuthUser struct {
	ID   int
	Role int
}

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		extractToken := func() string {
			bearerToken := r.Header.Get("Authorization")
			strArr := strings.Split(bearerToken, " ")
			if len(strArr) == 2 {
				return strArr[1]
			}
			return ""
		}
		if r.Header["Authorization"] != nil {
			tokenString := extractToken()
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				//Make sure that the token method conform to "SigningMethodHMAC"
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, errors.New("invalid authorization token")
				}
				return []byte(JWT_SECRET), nil
			})
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(&dtos.ErrorDto{Message: err.Error()})
				return
			}
			if token.Valid {
				claims := token.Claims.(jwt.MapClaims)
				user := &AuthUser{
					ID: int(claims[KeyUserID].(float64)),
				}
				userContext := context.WithValue(r.Context(), KeyUser, user)
				next.ServeHTTP(w, r.WithContext(userContext))
			} else {

				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(&dtos.ErrorDto{Message: errors.New("invalid authorization token").Error()})
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(&dtos.ErrorDto{Message: errors.New("an authorization header is required").Error()})
		}
	})
}

func GenerateToken(user *domains.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims[KeyUserID] = user.ID
	claims[KeyTokenExpired] = time.Now().Add(time.Hour * time.Duration(JWT_TTL_HOUR)).Unix()
	signedToken, err := token.SignedString([]byte(JWT_SECRET))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func GetAuthUser(r *http.Request) *AuthUser {
	user := r.Context().Value(KeyUser)
	if user == nil {
		return nil
	}
	return user.(*AuthUser)
}

func GetEmailVerifierToken(email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims[KeyUserID] = email
	claims[KeyTokenExpired] = time.Now().Add(time.Hour * time.Duration(JWT_TTL_HOUR)).Unix()
	signedToken, err := token.SignedString([]byte(JWT_SECRET))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func GetEmailFromVerifierToken(emailToken string) (string, error) {
	token, err := jwt.Parse(emailToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(JWT_SECRET), nil
	})
	if err != nil {
		return "", err
	}
	if token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		return claims[KeyUserID].(string), nil
	}
	return "", errors.New("invalid token")
}
